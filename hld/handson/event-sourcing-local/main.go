package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/exp/mmap"
	_ "modernc.org/sqlite"
)

type EventType string

const (
	Deposit  EventType = "deposit"
	Withdraw EventType = "withdraw"
)

type Event struct {
	Type    EventType `json:"type"`
	Account string    `json:"account"`
	Amount  int       `json:"amount"`
}

func appendEvent(filename string, evt Event) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	data = append(data, '\n')
	_, err = f.Write(data)
	return err
}

func loadEvents(filename string) ([]Event, error) {
	var events []Event
	r, err := mmap.Open(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	data := make([]byte, r.Len())
	_, err = r.ReadAt(data, 0)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		var evt Event
		if err := json.Unmarshal(scanner.Bytes(), &evt); err != nil {
			return nil, err
		}
		events = append(events, evt)
	}

	return events, nil
}

func initDB(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS account_balance (
		account TEXT PRIMARY KEY,
		balance INTEGER
	)
	`)
	return err
}

func applyEventToDB(db *sql.DB, evt Event) error {
	switch evt.Type {
	case Deposit:
		_, err := db.Exec(`
			INSERT INTO account_balance (account, balance)
			VALUES (?, ?)
			ON CONFLICT(account) DO UPDATE SET balance = balance + ?`,
			evt.Account, evt.Amount, evt.Amount)
		return err
	case Withdraw:
		_, err := db.Exec(`
			INSERT INTO account_balance (account, balance)
			VALUES (?, ?)
			ON CONFLICT(account) DO UPDATE SET balance = balance - ?`,
			evt.Account, -evt.Amount, evt.Amount)
		return err
	default:
		return nil
	}
}

func main() {
	const logFile = "ledger_events.log"
	const dbFile = "ledger.db"

	// Append some events
	if err := appendEvent(logFile, Event{Type: Deposit, Account: "alice", Amount: 100}); err != nil {
		panic(err)
	}
	if err := appendEvent(logFile, Event{Type: Withdraw, Account: "alice", Amount: 30}); err != nil {
		panic(err)
	}
	if err := appendEvent(logFile, Event{Type: Deposit, Account: "bob", Amount: 50}); err != nil {
		panic(err)
	}

	// Load events with mmap
	events, err := loadEvents(logFile)
	if err != nil {
		panic(err)
	}

	// Open SQLite DB
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize DB schema
	if err := initDB(db); err != nil {
		panic(err)
	}

	// Replay events into DB
	for _, evt := range events {
		if err := applyEventToDB(db, evt); err != nil {
			panic(err)
		}
	}

	// Query balances
	rows, err := db.Query(`SELECT account, balance FROM account_balance`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("Account balances:")
	for rows.Next() {
		var account string
		var balance int
		if err := rows.Scan(&account, &balance); err != nil {
			panic(err)
		}
		fmt.Printf("%s: %d\n", account, balance)
	}
}
