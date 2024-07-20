Requirements:

1. The digital wallet should allow users to create an account and manage their personal information.
2. Users should be able to add and remove payment methods, such as credit cards or bank accounts. 
3. The digital wallet should support fund transfers between users and to external accounts.
4. The system should handle transaction history and provide a statement of transactions.
5. The digital wallet should support multiple currencies and perform currency conversions.
  - Scoping Out.
7. The system should ensure the security of user information and transactions.
8. The digital wallet should handle concurrent transactions and ensure data consistency.
9. The system should be scalable to handle a large number of users and transactions.


Rough:
  - transfer user account to different account.
  - different ways of account .
  - transaction history.

- models
  - User
  - Account
    - Bank
    - Credit
  - Transactions
  
- usecases
  - DigitalWallet.
  - account repo
    - AddAccount ()
    - Transfer(accountId, ToaccountId)
