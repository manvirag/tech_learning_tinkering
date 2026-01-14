/*

Implement the following functions:

1.addCustomer(double revenue) → Adds a new customer with a unique ID and assigns the specified revenue, representing the amount the customer brings to the company.

1. addCustomer(double revenue, string referredBy) → Adds a new customer with the given revenue and records the referring customer. The revenue of the referring customer is also increased by the revenue of the newly added customer.
    
    Example:
    
    1 (100) → 2 (50) → 3 (30)
    
    Customer 2 is referred by customer 1.
    
    Customer 3 is referred by customer 2.
    
    Final revenue distribution:
    
    Customer 1: 150 (100 + 50)
    
    Customer 2: 80 (50 + 30)
    
    Customer 3: 30
    
2. getKCustomersWithRevenueGreaterThanValue(int k, double value) → Returns the IDs of the first k customers whose revenue is greater than the given value.
    
    Example:
    
    Customer ID (Revenue)
    
    1 (100)
    
    2 (120)
    
    3 (150)
    
    4 (200)
    
    5 (300)
    
    If k = 2 and value = 90, the function should return customers 1 and 2.
    

`Follow up - For getTopKCustomerWithGreaterValue, the expected solution should have a time complexity of less than O(N log K). Since the dataset can be large or in stream and this query is frequent, constructing a priority queue for each query is not feasible.`

*/