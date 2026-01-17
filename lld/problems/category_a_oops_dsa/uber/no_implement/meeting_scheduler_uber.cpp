/*

https://leetcode.com/discuss/post/5026730/uber-senior-software-engineer-phone-scre-2ki0/

Build a reservation system for a predefined set of conference rooms given as a list of room Ids [‘roomA’, roomB’...].

It should have a method like `scheduleMeeting(startTime, endTime)` should return a reservation identifier (including `roomId`) and reserve it or an error if no rooms are available.

- Assume there are 100 rooms, later can talk about scaling to N
- Assume any number of meetings can be scheduled on any conference room

YOE: 12 yrs

My approach:

```cpp
classRoomManger
{
using TimeSlot= std::pair<int,int>;// start, end
using Container= std::unordered_map<std::string, std::set<TimeSlot>>;// room-id to timeslots

    size_t occupied_rooms;
    std::vector<std::string> ROOM_LIST;// Some constants
    Container rooms_holder;

RoomManger(size_t N):ROOM_LIST(N),occupied_rooms(0), rooms_holder{}
{
// Fill room list, may be from configuration file ?
for(int i=0; i< N;++i)
{
            ROOM_LIST[i]="room"+ std::to_string(i+1);
}
}

    std::stringgetNextRoom()
{
if(occupied_rooms!= ROOM_LIST.size())
{
return ROOM_LIST[occupied_rooms++];
}
return{};
}

public:
    std::stringscheduleMeeting(int startTime,int endTime)
{
        TimeSlot time_slot={startTime, endTime};
        std::string room_id=getNextRoom();

if(!room_id.empty())
{
// room_id is available
            rooms_holder[room_id].insert(time_slot);
return room_id;
}

for(auto&[room_id, time_slots]: rooms_holder)
{
if(time_slots.empty())
{
                time_slots.insert(time_slot);
return room_id;
}

auto it= time_slots.lower_bound(time_slot);

if(it== time_slots.end())
{
                time_slots.insert(time_slot);
return room_id;
}

if(it->first>= time_slot.second)
{
                time_slots.insert(time_slot);
return room_id;
}
}

throw std::runtime_error("No rooms available");
}
};
```

The expectation was to do a simulation of the code.

I totally messed it up.

I couldn't complete it on time, and probably not the best data structures I have choosen.

The interviewer was pretty chill out, though.

Nothing interesting was hitting to me at that moment, I wasted lot of time explaining the my idea of figuring which rooms to pick and how to avoid overlaps.

At end when I thought I am done, I realized I missed to write code of overlapping timeslots, I tried to add some, qickly, but felt I could not, given the time constraints, so just ended explaining again.

At end we discussed about Time and Space complexity.

Probably it would have been a piece of cake if I'd have touch with DSA and codings.

*/