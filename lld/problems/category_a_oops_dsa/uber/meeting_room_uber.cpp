/*
->  Design meeting Scheduler
 Build an event assignment module. Imagine an office where you want to book rooms. 
 We basically need to assign a room to an input calendar event.

 Meeting room booking: 

There are a fixed number of meeting rooms with various sizes. Given a start time 
and end time, I needed to book an available meeting room. I also had to maintain past 
booking history of meeting rooms. The interviewer expected me to abstract different types 
of implementations for meeting room scheduling. Fully working code was expected within 1 hour.

1. There are N rooms.
2. We are given a stream of meeting requests (start time, end time, capacity).
3. We have to assign a room to the meeting if available, considering:
    - The room must be free during the requested time.
    - The room must have at least the required capacity.
4. We must minimize spillage of free time (i.e., use the room that has the least free time that can accommodate the meeting).
5. We have to store audit logs for each room (when a meeting is scheduled, etc.) and delete audit logs after X days.

**Additional considerations:**

- Concurrency: multiple requests may come in simultaneously.



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



...
are these ordred ? interval ? list {a1,b1} <= {a2,b2}   
*/

#include<iostream>
#include <mutex>
#include<set>
using namespace std;


struct Meeting {
    int start, end;
    int capacity;
    Meeting(int s,int e,int c): start(s), end(e), capacity(c) {}
};

struct AuditLog {
    int timestamp;
    string message;
    AuditLog(int ts, const string& msg): timestamp(ts), message(msg) {}
};

class Room {
public:
    int id;
    int capacity;
    set<pair<int,int>> schedule; // sorted intervals
    vector<AuditLog> auditLogs;
    mutex mtx;

    Room(int id,int cap): id(id), capacity(cap) {}

    // Check availability in O(log M)
    bool isAvailable(int start, int end) {
        lock_guard<mutex> lock(mtx);
        auto it = schedule.lower_bound({start, 0});
        if (it != schedule.end() && it->first < end) return false; // overlap with next
        if (it != schedule.begin()) {
            --it;
            if (it->second > start) return false; // overlap with previous
        }
        return true;
    }

    void addMeeting(int start,int end) {
        lock_guard<mutex> lock(mtx);
        schedule.insert({start,end});
    }

    void addAuditLog(int timestamp, const string& msg) {
        lock_guard<mutex> lock(mtx);
        auditLogs.push_back({timestamp, msg});
    }

    void cleanupAuditLogs(int currentTime, int ttlDays) {
        lock_guard<mutex> lock(mtx);
        auditLogs.erase(remove_if(auditLogs.begin(), auditLogs.end(),
                                  [&](const AuditLog& log){
                                      return currentTime - log.timestamp > ttlDays*86400;
                                  }), auditLogs.end());
    }

    // Calculate free time waste if this meeting is scheduled
    int freeTimeSpillage(int start,int end) {
        lock_guard<mutex> lock(mtx);
        int waste = INT_MAX;
        if (schedule.empty()) return 0;
        auto it = schedule.lower_bound({start,0});
        int prevEnd = 0;
        if (it != schedule.begin()) {
            auto temp = it; 
            temp--;
            prevEnd = temp->second;
        }
        int nextStart = INT_MAX;
        if (it != schedule.end()) {
            nextStart = it->first;
        }
        if (start >= prevEnd && end <= nextStart) {
            waste = nextStart - end + start - prevEnd;
        }
        return waste;
    }
};

class MeetingScheduler {
private:
    vector<Room*> rooms;
    mutex mtx;

public:
    MeetingScheduler(const vector<pair<int,int>>& roomInfo) {
        for (auto& p: roomInfo) {
            rooms.push_back(new Room(p.first,p.second));
        }
    }

    int scheduleMeeting(int start,int end,int capacity) {
        lock_guard<mutex> lock(mtx);
        Room* bestRoom = nullptr;
        int minWaste = INT_MAX;

        for (auto room: rooms) {
            if (room->capacity < capacity) continue;
            if (!room->isAvailable(start,end)) continue;
            int waste = room->freeTimeSpillage(start,end);
            if (waste < minWaste) {
                minWaste = waste;
                bestRoom = room;
            }
        }

        if (!bestRoom) return -1;

        bestRoom->addMeeting(start,end);
        bestRoom->addAuditLog(start,"Scheduled "+to_string(start)+"-"+to_string(end));
        return bestRoom->id;
    }

    void cleanupLogs(int currentTime, int ttlDays) {
        for (auto& room: rooms) room->cleanupAuditLogs(currentTime,ttlDays);
    }
};

/* ================= Demo ================== */
int main() {
    vector<pair<int,int>> roomInfo = {{1,10},{2,10},{3,10}}; // room capacity
    MeetingScheduler scheduler(roomInfo);

    cout << scheduler.scheduleMeeting(1647718624,1647718731,5) << endl; // 1
    cout << scheduler.scheduleMeeting(1647718624,1647718931,5) << endl; // 2
    cout << scheduler.scheduleMeeting(1647718624,1647718431,5) << endl; // 3
    cout << scheduler.scheduleMeeting(1647718624,1647768731,5) << endl; // -1
    cout << scheduler.scheduleMeeting(1647718732,1647728789,5) << endl; // 1

    return 0;
}
