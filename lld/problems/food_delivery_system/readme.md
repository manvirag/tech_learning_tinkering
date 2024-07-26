Requirements:

1. The food delivery service should allow customers to browse restaurants, view menus, and place orders.
2. Restaurants should be able to manage their menus, prices, and availability.
3. Delivery agents should be able to accept and fulfill orders.
4. The system should handle order tracking and status updates.
5. The system should support multiple payment methods.
6. The system should handle concurrent orders and ensure data consistency.
7. The system should be scalable and handle a high volume of orders.
8. The system should provide real-time notifications to customers, restaurants, and delivery agents




Rough

- models and repo
	- Restaurent, menu, order, notification -> notfiy -> every person / who wnats to get notficiations implement this interface.
		- list of menu
		- list of orders -> orderstatus
	- restaurent repos -> id: restaurent
		- restaurent : amount
- usecases
	- swiggy
		- list of delivery persons.
		- crud on restaurent repo
		- crud on order status 
		- payment -> cut fees -> send restaurent amount.

- main
	
