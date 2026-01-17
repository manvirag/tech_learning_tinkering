/*


->  Design meeting Scheduler

 Build an event assignment module. Imagine an office where you want to book rooms. We basically need to assign a room to an input calendar event.



 Meeting room booking: 

There are a fixed number of meeting rooms with various sizes. Given a start time and end time, I needed to book an available meeting room. I also had to maintain past booking history of meeting rooms. The interviewer expected me to abstract different types of implementations for meeting room scheduling. Fully working code was expected within 1 hour.

1. There are N rooms.
2. We are given a stream of meeting requests (start time, end time, capacity).
3. We have to assign a room to the meeting if available, considering:
    - The room must be free during the requested time.
    - The room must have at least the required capacity.
4. We must minimize spillage of free time (i.e., use the room that has the least free time that can accommodate the meeting).
5. We have to store audit logs for each room (when a meeting is scheduled, etc.) and delete audit logs after X days.

**Additional considerations:**

- Concurrency: multiple requests may come in simultaneously.

Design Meeting scheduler. There are n meeting rooms. We will keep getting requests for bookings - start time and end time. Allocate any room that is available.

(Designing classes, interactions, design patterns)

```csharp
For eg, Lets say Uber building has3 conference roomsand we start getting the requestsasfollows

scheduleMeeting(1647718624,1647718731)// assign let’s say room1.
scheduleMeeting(1647718624,1647718931)// same start time , assign room2 .
scheduleMeeting(1647718624,1647718431)// assign room3 as same start time.
scheduleMeeting(1647718624,1647768731)// error as no room is available for this time.
scheduleMeeting(1647718732,1647728789)// assign room1 as it became free

** We areusing unixtimestamps here,for simplicity  we can take integer valuesaswell**
```

`Expected Time Complexity: log(meetings) + O(rooms)` (since number of rooms are finite)
*/