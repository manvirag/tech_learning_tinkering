Requirements:

1. User: Each user should have a userId, name, email, mobile number. -> done
2. Expense: Could either be EQUAL, EXACT or PERCENT  -> Done
3. Users can add any amount, select any type of expense and split with any of the available users(atleast one).  -> Done
4. The percent and amount provided could have decimals upto two decimal places. -> validate
5. In case of percent, you need to verify if the total sum of percentage shares is 100 or not. -> Done
6. In case of exact, you need to verify if the total sum of shares is equal to the total amount or not. -> Done
7. When asked to show balances, the application should show balances of a user with all the users where there is a non-zero balance.
8. The amount should be rounded off to two decimal places. Say if User1 paid 100 and amount is split equally among 3 people. Assign 33.34 to first person and 33.33 to others. -> Done


Rough
- models
  - User
    - Id
    - Name
  - UserRepo
  - Expence
    - EqualExpense
      - amount
      - paidby
      - groupid
      - metadata
      - map(userid)flot
      - getamount()
    - Exact
    - Percentage.
  - Group
    - id
    - list of users
    - list of expenses
- usecases
  - splitwise
    - list of group
    - ShowViewByUser(Id)
- Algorithm:
  - https://leetcode.com/problems/optimal-account-balancing/