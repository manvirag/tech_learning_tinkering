Requirements
1. The system should allow users to view the list of movies playing in different theaters. --> Done
  - so there are list of theatres
  - inside theatre there are list of movies ?
  - Assume only only move in theature at a time 
3. Users should be able to select a movie, theater, and show timing to book tickets.  -> Done
  - there are different movie at different timing of theatre.
4. The system should display the seating arrangement of the selected show and allow users to choose seats. -> Done
  - insider theatre there are seats, vacant or fill
5. Users should be able to make payments and confirm their booking.
  - filled seat by payment.
6. The system should handle concurrent bookings and ensure seat availability is updated in real-time. 
  - locking on seat
7. The system should support different types of seats (e.g., normal, premium) and pricing.  -> Done
  - there are different categories for seat
8. The system should allow theater administrators to add, update, and remove movies, shows, and seating  -> Done
arrangements.
  - crud api for theatre, movie, seats
9. The system should be scalable to handle a large number of concurrent users and bookings.



- models
  - theatre  // assuming one threatre can play only one movie.
    - id
    - list of slot
    - list of seats
  - slot
    - starttime
    - endtime
    - movie
  - movie
    - id
    - name
- repo
  - inventory
    - list of theators
    - crud in theatre.
- usecase
  - repo
    - get movies in theatre.
    - book the ticket in theatre, seat no, movie, slot. concurrent.
- controller
