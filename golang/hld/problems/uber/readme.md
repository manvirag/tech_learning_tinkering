[ Kind of copy paste ]

### Functional Requirements:

We will design our system for two types of users: Customers and Drivers.

***Customers***

1. Customers should be able to see all the cabs in the vicinity with an ETA and pricing information. 
2. Customers should be able to book a cab to a destination. 
3. Customers should be able to see the location of the driver.

***Drivers***

1. Drivers should be able to accept or deny the customer requested ride.
2. Once a driver accepts the ride, they should see the pickup location of the customer.
3. Drivers should be able to mark the trip as complete on reaching the destination.

#### Extended requirements
- Customers can rate the trip after it's completed.
- Payment processing.
- Metrics and analytics.

### Non-Functional Requirements:

1. High reliability.
2. High availability with minimal latency.
3. The system should be scalable and efficient.

### Capacity Estimation

Let's start with the estimation and constraints.

Note: Make sure to check any scale or traffic-related assumptions with your interviewer.

***Traffic***

Let us assume we have 100 million daily active users (DAU) with 1 million drivers and on average our platform enables 10 million rides daily.

If on average each user performs 10 actions (such as request a check available rides, fares, book rides, etc.) we will have to handle 1 billion requests daily.


***What would be Requests Per Second (RPS) for our system?***

1 billion requests per day translate into 12K requests per second.



***Storage***

If we assume each message on average is 400 bytes, we will require about 400 GB of database storage every day.


And for 10 years, we will require about 1.4 PB of storage.


***Bandwidth***

As our system is handling 400 GB of ingress every day, we will require a minimum bandwidth of around 5 MB per second.

***Here is our high-level estimate***:

| Type                   | Estimate      |
|------------------------|---------------|
| Daily active users (DAU) | 100 million   |
| Requests per second (RPS) | 12K/s       |
| Storage (per day)      | ~400 GB       |
| Storage (10 years)     | ~1.4 PB       |
| Bandwidth              | ~5 MB/s       |

### High level design

![alt_text](./images/img.png)

#### Data models

user (can be driver or customer):
```json
{
  "userId": "",
  "type": "DRIVER|CUSTOMER"
}
```

Ride :
```json
{
  "customerId": "",
  "driverId": "",
  "cabId": "",
  "source": "",
  "destination": "",
  "price": "",
  "time": ""
}
```

```json
{
  "cabId": "",
  "ownerId": ""
}
```


#### Api design (copy pasted boring)

Let us do a basic API design for our services:

#### Request a Ride

Through this API, customers will be able to request a ride.

```tsx
requestRide(customerID: UUID, source: Tuple<float>, destination: Tuple<float>, cabType: Enum<string>, paymentMethod: Enum<string>): Ride
```

**Parameters**

Customer ID (`UUID`): ID of the customer.

Source (`Tuple<float>`): Tuple containing the latitude and longitude of the trip's starting location.

Destination (`Tuple<float>`): Tuple containing the latitude and longitude of the trip's destination.

**Returns**

Result (`Ride`): Associated ride information of the trip.

#### Cancel the Ride

This API will allow customers to cancel the ride.

```tsx
cancelRide(customerID: UUID, reason?: string): boolean
```

**Parameters**

Customer ID (`UUID`): ID of the customer.

Reason (`UUID`): Reason for canceling the ride _(optional)_.

**Returns**

Result (`boolean`): Represents whether the operation was successful or not.

#### Accept or Deny the Ride

This API will allow the driver to accept or deny the trip.

```tsx
acceptRide(driverID: UUID, rideID: UUID): boolean
denyRide(driverID: UUID, rideID: UUID): boolean
```

**Parameters**

Driver ID (`UUID`): ID of the driver.

Ride ID (`UUID`): ID of the customer requested ride.

**Returns**

Result (`boolean`): Represents whether the operation was successful or not.

#### Start or End the Trip

Using this API, a driver will be able to start and end the trip.

```tsx
startTrip(driverID: UUID, tripID: UUID): boolean
endTrip(driverID: UUID, tripID: UUID): boolean
```

**Parameters**

Driver ID (`UUID`): ID of the driver.

Trip ID (`UUID`): ID of the requested trip.

**Returns**

Result (`boolean`): Represents whether the operation was successful or not.

#### Database
Usecase:

1. get user and drive information.
2. save the ride.
3. get geo hash.

Since our data is structured, we can use sql.
Also nosql cassandra can also be used.

#### Cache Schema

***Location truth cache***
![alt_text](./images/img_3.png)
This cache is the source of truth when it comes to user location. Phone apps send regular updates to maintain accuracy. If a user gets disconnected, its record expires in 30 seconds.

***Driver proximity cache***
![alt_text](./images/img_4.png)

The proximity cache is crucial for nearby driver search. Given a location, we can use GeoHash to compute its location key and retrieve all drivers in the grid.

#### Architecture

With a clearer understanding of what data to store, now it's time for service-oriented designs!

***Notification Service:*** whenever the backend needs to send information to the clients, the notification service is used to deliver the messages.
***Trip Management Service:*** When a trip is initiated, this service is needed to monitor the locations of all parties as well as plan routes
***Ride Matching Service:*** This service handles ride requests. It finds nearby drivers and matchmakes based on driver responses (either accept or decline)
***Location Service:*** All users must regularly update their locations via this service.


#### Deep dive


1. ***How to efficiently look up nearby driver (proximity service. Also learn differently)***
   Instead of computing the distance between the rider and every driver in the database, we can use a technique called GeoHash to convert the user's location to a unique key that corresponds to one cell in figure 7 and confine the search to a few adjacent cells only.

![alt_text](./images/img_1.png)

Therefore, we can use the following heuristic to quickly lookup drivers:

1. Given a user location, compute its GeoHash from longitude and latitude.
2. With the GeoHash key available, it's easy to compute the keys for the 8 nearby cells. (see this post)
3. Query the Location service with all 9 keys; retrieve drivers in those cells.

The accuracy of GeoHash is determined by the key length. The longer the key is, the smaller each cell would be.
![alt_text](./images/img_2.png)


2. ***Location Update***

There are two tables in the cache — the location truth table and the driver proximity table.

The usage of the location truth table is simple. The user's mobile app sends out its location as well as the user ID to the Location service every 5 seconds. Whenever a user's accurate location is needed, the system can query this table by user ID.

![alt_text](./images/img_5.png)

It is hard to remove drivers from their old cells. The implication is obvious; the same driver can appear in multiple cells because of the stale data.

To deal with this problem, we could introduce a timestamp to each record. With timestamps, it is easy to filter out stale location data [ Has to check more ]

#### Reference:
1. https://github.com/karanpratapsingh/system-design?tab=readme-ov-file#uber
2. https://towardsdatascience.com/ace-the-system-design-interview-uber-lyft-7e4c212734b3