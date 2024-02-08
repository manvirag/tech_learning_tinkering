Design the flipkart/amazon

### Functional Requirements

1. Sellers should be able to add, delete and modify products they want to sell.
2. The website should include a catalog of products.
3. Buyers can search products by name, keyword or category.
4. Buyers can add, delete or update items in a cart.
5. Buyers can purchase items in the cart and make payments.
6. Buyers can view their previous orders.
7. Buyers can review and rate purchased products.

### Non-Functional Requirements

1. High availability
2. High consistency
3. Low latency

Maintaining all three non-functional requirements at all times is difficult, especially with the high traffic that the website will need to handle. However, different components of the system architecture can have different priorities.

For example, payment service and inventory management will need to be highly consistent. Availability is of low priority in this case. Search service, on the other hand, should be highly available and have low latency, even at the cost of consistency. In other words, it’s acceptable if the search service is eventually consistent.

### Capacity Estimation
### High level design

e-commerce platforms have two sides for the system. If you are a customer of Amazon, you can either be selling products through your own isolated store on Amazon, or you can be buying products available from different sellers

### Deep dive in design





Reference:
1. https://readmedium.com/en/https:/medium.com/double-pointer/system-design-interview-amazon-flipkart-ebay-or-similar-e-commerce-applications-35a0bc764421
2. https://www.codekarle.com/system-design/Amazon-system-design.html

