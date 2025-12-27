#include<iostream> 
using namespace std; 

/*

Interview Problem: In-Memory Banking System

You are asked to design an in-memory banking system that supports account management, transactions, and historical queries.

The problem is divided into four parts.
Each part builds on top of the previous one and may require refactoring your existing solution.

You should assume all data fits in memory.

Part 1 – Basic Account Operations

Implement a banking system that supports the following operations:

deposit(accountId, amount, timestamp)
Deposits the given amount into the specified account.

withdraw(accountId, amount, timestamp)
Withdraws the given amount from the specified account.
The operation should fail if the account does not have sufficient balance.

getBalance(accountId)
Returns the current balance of the account.

Notes

All amounts are positive integers.

Withdrawals that fail should not modify account state.

All operations are timestamped.

Part 2 – Transaction Activity and Analytics

Extend the system to track transaction activity.

Transaction activity is defined as the number of successful financial operations performed by an account.

New requirements

getTopKActiveAccounts(k)
Returns the IDs of the top k accounts with the highest transaction activity.
In case of ties, accounts should be ordered lexicographically.

Ensure that only successful operations contribute to transaction activity.

(You may assume additional helper operations are required internally to support this functionality.)

Part 3 – Two-Phase Money Transfers

Extend the system to support money transfers between accounts.

New operations

beginTransfer(fromAccount, toAccount, amount, timestamp)
Initiates a transfer:

Deducts the amount from fromAccount

Creates a pending transfer

Does not credit toAccount yet

Does not increase transaction activity

acceptTransfer(transferId, timestamp)
Completes a previously initiated transfer:

Credits the amount to toAccount

Marks the transfer as completed

Updates transaction activity for both accounts

Notes

Transfers must be accepted exactly once.

Invalid or duplicate transfer acceptances should fail.

Transfers should fail if any account is invalid or if funds are insufficient at initiation.

Part 4 – Account Merging and Historical Queries

Extend the system to support account merging and historical balance queries.

New requirements

mergeAccounts(sourceAccount, targetAccount)

Merges sourceAccount into targetAccount

The source account becomes inactive

Transaction history for the source account must remain accessible

getBalanceAt(accountId, timestamp)
Returns the balance of the specified account at the given timestamp.

Notes

Historical balance queries must work even after accounts are merged.

Merged accounts should no longer accept new operations, but their history must still be queryable.

General Constraints and Expectations

All timestamps are integers

Operations are processed in the order they are received

The system does not need to be thread-safe

Focus on correctness and clean state transitions rather than performance optimizations

*/



