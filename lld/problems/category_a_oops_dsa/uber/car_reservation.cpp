/*

→ Machine coding round of a car reservation service like Avis, Hertz

 Design a car reservation system to efficiently allocate a list of n cars to maximize the number of customers served. The solution should focus on optimizing resource utilization and ensuring that as many customer requests as possible are fulfilled.



 via chat gpt: 


 Problem Statement

Uber operates a fleet of n cars in a city. Each car can serve one customer at a time.
You are given a list of customer reservation requests, where each request specifies:

-   A start time when the customer wants to pick up the car
-   An end time when the customer will return the car

Your task is to design a system that allocates cars to customers such that:
- The total number of customers served is maximized.

If no car is available for a request, the request must be rejected.



*/

// assumption request time are in sorted order (start,end)

// #include<iostream> 
// #include<set> 
// using namespace std; 

// class CarReservationSystem {
//     public: 
//         int car;
//         int totalServedReq;
//         int availableCarCount;
//         set<pair<int,int>> pastRequestsEndTime;
//         int totalRequestCount;
//     CarReservationSystem(int carCount): car(carCount), totalServedReq(0), availableCarCount(carCount),totalRequestCount(0) {}

//     void serveRequest(int start, int end) {

//         while(!this->pastRequestsEndTime.empty() && this->pastRequestsEndTime.begin()->first <= start) {
//             this->pastRequestsEndTime.erase(this->pastRequestsEndTime.begin());
//             this->availableCarCount++;
//         }
//         if(this->availableCarCount == 0) {
//             cout<<"no car is available "<<start<<" "<<end<<endl;
//             return ; 
//         }
//         this->availableCarCount--;
//         this->pastRequestsEndTime.insert({end, ++(this->totalRequestCount)});
//         this->totalServedReq++;

//     }
//     int getTotalServedRequest() {
//         return  this -> totalServedReq; 
//     }
// };
// int main() {
//     CarReservationSystem cs = CarReservationSystem(2);
//     cs.serveRequest(1,2);
//     cs.serveRequest(2,2);
//     cs.serveRequest(1,5);
//     cout<<cs.getTotalServedRequest()<<endl;
//     return 0;
// }


#include<iostream> 
#include<set> 
#include<shared_mutex>
#include<thread> 
using namespace std; 


class CarReservationSystem {
    public: 
        int car;
        shared_mutex mu; 
        int totalServedReq;
        int availableCarCount;
        set<pair<int,int>> pastRequestsEndTime;
        int totalRequestCount;
    CarReservationSystem(int carCount): car(carCount), totalServedReq(0), availableCarCount(carCount),totalRequestCount(0) {}

    void serveRequest(int start, int end) {
        unique_lock<shared_mutex> lock(mu);
        while(!pastRequestsEndTime.empty() && pastRequestsEndTime.begin()->first <= start) {
            pastRequestsEndTime.erase(pastRequestsEndTime.begin());
            availableCarCount++;
        }
        if(availableCarCount == 0) {
            cout<<"no car is available "<<start<<" "<<end<<endl;
            return ; 
        }
        availableCarCount--;
        pastRequestsEndTime.insert({end, ++(totalRequestCount)});
        totalServedReq++;

    }
    int getTotalServedRequest() {
        shared_lock<shared_mutex> lock(mu);
        return  this -> totalServedReq; 
    }
};
int main() {
    CarReservationSystem cs = CarReservationSystem(2);
    vector<thread> tts;
    tts.emplace_back([&](){
        cs.serveRequest(1,2);
    });

    tts.emplace_back([&](){
        cs.serveRequest(2,2);
    });

    tts.emplace_back([&](){
        cs.serveRequest(1,5);
    });
    
    
    tts.emplace_back([&](){
        cout<<cs.getTotalServedRequest()<<endl;
    });

    for(auto &t: tts) {
        if(t.joinable())
         t.join();
    }

    
    
    return 0;
}