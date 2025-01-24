# Distributed transaction

[https://www.youtube.com/watch?v=eltn4x788UM](https://www.youtube.com/watch?v=eltn4x788UM)

[https://www.youtube.com/watch?v=7FgU1D4EnpQ](https://www.youtube.com/watch?v=7FgU1D4EnpQ)

Problem statement:

we have multiple databases and we want to do transactions including both. Since both have their different transaction, how do we combine them?

Same problem to understand this.:

here store and delivery are the two different service and having their own database.

![Untitled](Distributed%20transaction/Untitled.png)

![Untitled](Distributed%20transaction/Untitled%201.png)

Some Solution:

1. Two-Phase Commit.

[https://drive.google.com/file/d/1FJa0DQPOxVJP2kXdyZDSttdMDI4Q2fUx/view](https://drive.google.com/file/d/1FJa0DQPOxVJP2kXdyZDSttdMDI4Q2fUx/view)

![Untitled](Distributed%20transaction/Untitled%202.png)

![Untitled](Distributed%20transaction/Untitled%203.png)

![Untitled](Distributed%20transaction/Untitled%204.png)

https://www.linkedin.com/pulse/implementing-distributed-transactions-golang-arpit-bhayani/

https://medium.com/@abhishekranjandev/implementing-distributed-transactions-with-golang-and-gin-c6f00297fc21

how to implement pre phase ? 

- There is already predefined PREPARE key word in mysql like abort and commit.

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

What if first commit → success and second → failed ? 

What about the first query then ? how it is atomix > 

Here are a few cases where a failure might happen after the prepare phase

1. **Coordinator or system crash**: If the transaction coordinator crashes between the prepare and commit phases, it may not send the commit command to some of the participants.
2. **Network failures**: If the coordinator cannot communicate with a participant after the prepare phase, the participant may remain in the prepared state indefinitely, waiting for a commit or rollback decision.
3. **Database crash**: Even if a participant has prepared successfully, a database crash before the commit can cause the participant to lose its state.

2PC does not inherently guarantee that a prepared transaction will always commit successfully. Instead, it relies on **recovery mechanisms** to handle failures.

To address the risk of failure between the **prepare** and **commit** phases, most systems that implement 2PC use additional mechanisms like **persistent logs** to ensure reliability

like WAL

### Handling Failures During Commit

1. **Coordinator Failure**:
    - If the coordinator fails after sending the prepare command, the participants wait. Once the coordinator is back up, it reads its logs to determine whether to commit or abort the transaction and notifies the participants accordingly.
2. **Participant Failure**:
    - If a participant crashes after preparing but before committing, it can recover its state by reading its logs after restarting. The participant can then ask the coordinator for the final decision and proceed with either a commit or rollback.

still not clear much

Once a participant has committed, it cannot roll back, which leads to inconsistency in the system (some participants have committed while others have not). This is one of the known limitations of 2PC,  if a participant crashes after committing, it can recover and notify the coordinator that it has successfully committed. However, this logging doesn’t prevent the partial commit scenario but helps in recovering from crashes by knowing what was the last state of the transaction. Still, once a participant commits, there's no way to undo it. ( **ChatGPT** ) → Pager and fix it manually or other way.

There is better solution

3PC → it has pre commit phase → so that to make more sure , that transaction doesn’t fail in commit phase. 

Martin-Fowler: [https://martinfowler.com/articles/patterns-of-distributed-systems/two-phase-commit.html](https://martinfowler.com/articles/patterns-of-distributed-systems/two-phase-commit.html)

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
	- Its very less probable that after prepare, node failed to commit, but possible ( just clarifying phase nothing but begin and exec command without command as mentioned in above golang code). 
	- Failure cases:
		- Fail in middle of prepare -> abort all . --> consistent. ( via node )
		- Fail in middle of commit ( via node ) -> will require to maintain the status of all commit and rollback them and make it consistent state. ( that's why it is important , that our system is fault tolerance to this failure, shouldn't be disacter in consistency , it should work well -> like in case of digital wallet , we remove money first from account A and commit , event after that it fail, that is very disacter at as of now , once we get to know about failure -> we will validate and increase the amount of A)
		- Failure via coordinator crash -> in middle of prepare -> very risky -> all node will be stuck until the coordinator recover and locking those row for other -> disaster. => how to solve this ??
			- There are some solution -> mentioned in the above notes as fault tolerant two phase commit -> high level all nodes including coordinate will be in consensus algorithm and share their heartbeat to other node, and if any node crash , we abort the transactions.
			- Some other solution -> TC/C, Saga, they have their own pros and cons
		- Failure via coordinater -> in middle of commit -> same , after recover with help of status rollback things.

2. TC/C
3. Saga










