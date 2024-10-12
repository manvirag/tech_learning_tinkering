1. The ride sharing service should allow passengers to request rides and drivers to accept and fulfill those ride requests. --> done
  - rider and driver
  - rider request -> find driver -> reponse to rider and update driver state
2. Passengers should be able to specify their pickup location, destination, and  --> done
  - okk at the time of request it will share these
4. Drivers should be able to see available ride requests and choose to accept or decline them. 
  -  it seems like notfication to all driver -> they response
5. The system should match ride requests with available drivers based on proximity and other factors. --> done
  -  para meters of a ride
6. The system should calculate the fare for each ride based on distance, time, and ride type. --> done
  - parameter of rides
7. The system should handle payments and process transactions between passengers and drivers.
  - locking code 
8. The system should provide real-time tracking of ongoing rides and notify passengers and drivers about ride status updates.
  - again notfication type
9. The system should handle concurrent requests and ensure data consistency.
- cool


- Rough
  - scope out as of now different type of rides.desired ride type (e.g., regular, premium).


- lld

- models
  - user -> rider/driver and userrepo
    - name
    - userid
    - type -> Rider / Driver // temp
  - account
    - accountid, amount
  - ride
    - startLocation
    - endLocation
    - distance
    - totaltime
    - totalamount
    - driverid
    - userid
  - payment gateway
    - usera, userb
    - transfer x amount a -> b
      
- usecases
  - UberSystem
      - user repo
      - payment gateway
      - amount repo
      - list of driver as subscriber.
      - list of rider to get notify about ride status
    - once ride start -> rider get information about driver rider updates.
    - rider can create ride.
        - - rider to get notify about the rides and accept/reject it.
        - - rider to driver matching algorithm.
        - - calculate fare with meta inf. like distance etc.  
      - calculate price
      - send notfication of driver.
      - once driver accept
        - send to rider
        - payment gateway to pay money from rider to driver
      - else loading to rider
      


