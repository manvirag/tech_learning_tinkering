Requirements
1. The online stock brokerage system should allow users to create and manage their trading accounts.
2. Users should be able to buy and sell stocks, as well as view their portfolio and transaction history.
3. The system should provide real-time stock quotes and market data to users.
4. The system should handle order placement, execution, and settlement processes.
5. The system should enforce various business rules and validations, such as checking account balances and stock availability.
6. The system should handle concurrent user requests and ensure data consistency and integrity.
7. The system should be scalable -> 
8. and able to handle a large number of users and transactions.



Rough Work:

models:
- User
- Stock
- Order
- StockExchange
- Portfolio
- Account

repositories
- User
- Stock

usecases
- user crud  -> user repository -> account
- stock listing -> stock repository -> 
- user order history -> transaction history -> list orders
- order execute -> order model -> state design pattern
  
  

Doubts:
- What we want to show in real-time stock listing ?
  - Just show stock information , counts , price
- How to do stock availability , who knows this ?
  - maintain counts
- What does trading account mean here ?
  - acount contain user and its balances portfolio
- How to execute when someone wants to buy and some wants to sell.
  - 
- Can you define placement, execution and settlement process means over here ?
  - Make this as stratedy , there can be different way
    - simple infinite stock
    - buy and sell matching
    - concurrent buy and sell maching.