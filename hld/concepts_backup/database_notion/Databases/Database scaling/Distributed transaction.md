# Distributed transaction


Problem statement:

we have multiple databases and we want to do transactions including both. Since both have their different transaction, how do we combine them?

Sample problem to understand this.:
here store and delivery are the two different service and having their own database.

Some Solution:

1. Two-Phase Commit.

how to implement pre phase ? 

Code:

```
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to multiple MySQL databases (replicas/participants)
	db1, err := sql.Open("mysql", "user:password@tcp(db1_host)/dbname")
	if err != nil {
		log.Fatalf("Failed to connect to db1: %v", err)
	}
	defer db1.Close()

	db2, err := sql.Open("mysql", "user:password@tcp(db2_host)/dbname")
	if err != nil {
		log.Fatalf("Failed to connect to db2: %v", err)
	}
	defer db2.Close()

	// Begin the two-phase commit
	err = twoPhaseCommit([]*sql.DB{db1, db2})
	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	} else {
		log.Println("Transaction successful")
	}
}

func twoPhaseCommit(dbs []*sql.DB) error {
	var transactions []*sql.Tx
	var err error

	// 1. Start the prepare phase
	for _, db := range dbs {
		tx, err := db.Begin()
		if err != nil {
			log.Printf("Failed to begin transaction: %v", err)
			rollbackAll(transactions) // Rollback all successful transactions
			return fmt.Errorf("failed to start transaction: %w", err)
		}
		transactions = append(transactions, tx)
		
		// Simulate some database operations
		_, err = tx.Exec("INSERT INTO some_table (column) VALUES (?)", "some_value")
		if err != nil {
			log.Printf("Failed to execute statement: %v", err)
			rollbackAll(transactions)
			return fmt.Errorf("failed to execute statement: %w", err)
		}
	}

	// 2. Commit phase: If all transactions are successful, commit
	for _, tx := range transactions {
		if err := tx.Commit(); err != nil {
			log.Printf("Failed to commit transaction: %v", err)
			rollbackAll(transactions)
			return fmt.Errorf("commit failed: %w", err)
		}
	}

	return nil
}

// Rollback all transactions if anything fails
func rollbackAll(transactions []*sql.Tx) {
	for _, tx := range transactions {
		if tx != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("Failed to rollback transaction: %v", err)
			}
		}
	}
}

```
................................................................


Re-visit:

We know here, what actually is distributed transaction is right ? 

Now we will discuss the few methods( that i read ), which help to implement the distributed transaction in real life.

1. Two Phase Commit. 

- Resource: -> Martin Ji DDIA wale.
- [Yt Link](https://www.youtube.com/watch?v=-_rdWB9hN1c&list=PLeKd45zvjcDFUEv_ohr_HdUFe97RItdiB&index=19&ab_channel=MartinKleppmann) , notes in book section

- So in this as mentioned in preious visit, all information are correct about implementation.
- Will Add few points, related to cons or general.
	- It has two phase Prepare and commit.
	- Prepare phase -> all db send yes , they make the local transaction and check validate and send yes , and they also ready to receive the commit from coordinator, here its serious that they are ready for commit like promising some human being and suppose if db not get the abort/commit back they will stuck infinitely , until coordination recover. This is problematic 1.
	- Its very less probable that after prepare, node failed to commit, but possible ( just clarifying phase nothing but begin and exec command without commit as mentioned in above golang code). 
	- Failure cases:
		- Fail in middle of prepare -> abort all . --> consistent. ( via node )
		- Fail in middle of commit ( via node ) -> will require to maintain the status of all commit and rollback them and make it consistent state. ( that's why it is important , that our system is fault tolerance to this failure, shouldn't be disacter in consistency , it should work well -> like in case of digital wallet , we remove money first from account A and commit , after that it fail, that is very disacter at as of now , once we get to know about failure -> we will validate and increase the amount of A), that's why sometime it called blocking protol.
		- Failure via coordinator crash -> in middle of prepare -> very risky -> all node will be stuck until the coordinator recover and locking those row for other -> disaster. => how to solve this ??
			- There are some solution -> mentioned in the above notes as fault tolerant two phase commit -> high level all nodes including coordinate will be in consensus algorithm and share their heartbeat to other node, and if any node crash , we abort the transactions.
			- Some other solution -> TC/C, Saga, they have their own pros and cons
		- Failure via coordinater -> in middle of commit -> same , after recover with help of status rollback things.
        - Note cooridination nothing but the server implementing distributed transaction, node are nothing but the database which are part of distributed transaction.

2. TC/C ( Try Confirm/Cancel)
	- Its a compensating transaction as mentioned by ALEX xu, what that mean ?  -> tx which can do undo of failed transaction i.e. rollback
	- On high level it has two phase try -> confirm/cancel.
	- both phases have their commit, not like 2 phase.
	- in first phase we send the tnx for commit to one node and other as NOF ( no operation ) once both commit -> then either confirm second commit or cancle the first commited.( rollback ). [ Just a very high level . Need Revisit and proper understand as of now , let it me like this. ] 
	- Its kind of same as SAGA, but can do parallel execution ( alex xu )

3. 3 Phase commit -> Non-Blocking
	- High level
	- commit phase is divided in two part
	- similar to fault tolerant, all node will send information which is available across all node include coordination.
	- and commit is break into 2 precommit and commit, precommit nothing but the indication to all node that coordination received the prepare phase yes and going to commit.
	- See concept and coding once or other doc [ Assume not 100% sure]
4. Saga 
	- Its nothing but the linear transaction, first do transaction in one db and then in other db with maintaining the state of transaction in durable store.
	- if first fail then fine, if fail in middle rollback one by one all previous commited transactions.
	- This is also a general pattern in distributed systems.
	- This can be implemented either by async way -> called choreography . like queue wise by subscribing other events. [ complexity increases]
	- or by coordinator -> called orchestration. 
	- Not deep diving. Revisit if needed.
	- via chat gpt 
	- https://threedots.tech/post/distributed-transactions-in-go/
	
```
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Saga Coordinator
type SagaStep struct {
	Action     func() error
	Compensate func() error
}

type Saga struct {
	steps []SagaStep
}

func NewSaga() *Saga {
	return &Saga{}
}

func (s *Saga) AddStep(action, compensate func() error) {
	s.steps = append(s.steps, SagaStep{Action: action, Compensate: compensate})
}

func (s *Saga) Execute() error {
	for i, step := range s.steps {
		if err := step.Action(); err != nil {
			log.Printf("Error in step %d: %v. Rolling back...", i, err)
			for j := i - 1; j >= 0; j-- {
				if err := s.steps[j].Compensate(); err != nil {
					log.Printf("Failed to compensate step %d: %v", j, err)
				}
			}
			return err
		}
	}
	return nil
}

// Database Connection
func Connect(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}
	return db, nil
}

// Order Repository
type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(orderID, userID string, amount float64) error {
	_, err := r.db.Exec("INSERT INTO orders (id, user_id, amount, status) VALUES ($1, $2, $3, 'PENDING')", orderID, userID, amount)
	return err
}

func (r *OrderRepository) RollbackOrder(orderID string) error {
	_, err := r.db.Exec("DELETE FROM orders WHERE id = $1", orderID)
	return err
}

// Payment Repository
type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) ProcessPayment(paymentID, userID string, amount float64) error {
	_, err := r.db.Exec("INSERT INTO payments (id, user_id, amount, status) VALUES ($1, $2, $3, 'COMPLETED')", paymentID, userID, amount)
	return err
}

func (r *PaymentRepository) RollbackPayment(paymentID string) error {
	_, err := r.db.Exec("DELETE FROM payments WHERE id = $1", paymentID)
	return err
}

// Order Service
type OrderService struct {
	repo *OrderRepository
}

func NewOrderService(repo *OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(orderID, userID string, amount float64) error {
	return s.repo.CreateOrder(orderID, userID, amount)
}

func (s *OrderService) RollbackOrder(orderID string) error {
	return s.repo.RollbackOrder(orderID)
}

// Payment Service
type PaymentService struct {
	repo *PaymentRepository
}

func NewPaymentService(repo *PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) ProcessPayment(paymentID, userID string, amount float64) error {
	return s.repo.ProcessPayment(paymentID, userID, amount)
}

func (s *PaymentService) RollbackPayment(paymentID string) error {
	return s.repo.RollbackPayment(paymentID)
}

// Main Function
func main() {
	orderDB, err := Connect("postgres://user:password@localhost:5432/orderdb?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to order DB: %v", err)
	}
	defer orderDB.Close()

	paymentDB, err := Connect("postgres://user:password@localhost:5432/paymentdb?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to payment DB: %v", err)
	}
	defer paymentDB.Close()

	orderRepo := NewOrderRepository(orderDB)
	orderSvc := NewOrderService(orderRepo)

	paymentRepo := NewPaymentRepository(paymentDB)
	paymentSvc := NewPaymentService(paymentRepo)

	// Saga Orchestration
	orderID := "order123"
	userID := "user123"
	amount := 100.0

	s := NewSaga()
	s.AddStep(
		func() error { return orderSvc.CreateOrder(orderID, userID, amount) },
		func() error { return orderSvc.RollbackOrder(orderID) },
	)
	s.AddStep(
		func() error { return paymentSvc.ProcessPayment(orderID, userID, amount) },
		func() error { return paymentSvc.RollbackPayment(orderID) },
	)

	if err := s.Execute(); err != nil {
		log.Fatalf("Saga failed: %v", err)
	} else {
		log.Println("Transaction completed successfully")
	} }
```
 











