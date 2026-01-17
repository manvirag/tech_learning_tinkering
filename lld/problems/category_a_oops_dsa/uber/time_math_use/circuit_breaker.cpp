/*

https://leetcode.com/discuss/post/5049157/uber-machine-coding-by-anonymous_user-o9i2/

// Implement simple circuit breaker given a rpc method.
// please consider the generic circuit breaker options like the following:
//
// * timeWindowSec - (10s) time sliding window size.
// * failureRatioThreshhold - (50%) failure ratio threshold in which the circuit open
// * circuitCloseTimeSec - (5s) required time to circuit re-close
// * min requests - 10

// Main class should be named 'Solution' and should not be public.

*/



/*

cb 
-> state -> close -> open -> half open -> close 

no half open here -> open -> clost -> open 
request send -> first remove < current - 10 time 
if < 10 -> hit 
if >=10 
    -> check for open ?? 
    -> check >=hafl request failed. 
    -> open 
open 
    -> if open time >= curr - 5 econd -> close
    -> requests ?? -> current - 10 - current - 5 -> let it be as it us
50% -> open -> 

for simplicity -> taking time as int ->  simple int. 


state -> close/open 

request -> 
    -> open 
        -> >5 wait 
            -> state -> close 
            -> reset past requests.
            -> allow 
        -> reject 
    -> close 
        -> remove beofre 10 seconds sliding window 
        -> if < 10 - > allow 
        -> after allow if >=10 
        -> 50% fail 
            -> open 
            -> failuretime 

*/

#include<iostream>
#include<set> 
using namespace std; 
int t = 1;
bool callRequest() {
    return false;
}
enum class CBState {
    OPEN,
    CLOSE,
};

int getCurrentTime() {
    return t++;
}
class CircuitBreakder {
    public: 
        int timeWindowSeconds;  // duration type in prod
        double failureRatioThreshhold; 
        int circuitCloseTimeSec;      // time type in prod
        int minRequests;

        int lastOpenedTime;     // -1 default
        CBState state;
        multiset<pair<int,int> > pastRequests;
        int failedRCount;


    CircuitBreakder(int timeWinSec, int failureRatio, int circuitCloseT, int minR): timeWindowSeconds(timeWinSec), failureRatioThreshhold(failureRatio), circuitCloseTimeSec(circuitCloseT),minRequests(minR), lastOpenedTime(-1), state(CBState::CLOSE) {}

    void callFunction() {
        if(state == CBState::CLOSE) {
            int currentTime = getCurrentTime();
            int ct = currentTime-timeWindowSeconds;
            while(!pastRequests.empty() && pastRequests.begin()->first < ct) {
                if(!pastRequests.begin()->second) {
                    failedRCount--;
                }
                pastRequests.erase(pastRequests.begin());
                
            }
            if(pastRequests.size() < minRequests) {
                bool result = callRequest();
                pastRequests.insert({currentTime,result});
                if(result == false) failedRCount++;
                cout<<"insert request "<<endl;
            }

            if(pastRequests.size() >= minRequests) {
                double per = (double(failedRCount)/minRequests)*100;
                if(per >= failureRatioThreshhold) {
                    cout<<"making it state to fail"<<endl;
                    setFailState(currentTime);
                }
            }

        } else {
            int diff = getCurrentTime()-lastOpenedTime;
            if(diff >= circuitCloseTimeSec) {
                setCloseState();
                cout<<"making it state to close"<<endl;
                int result = callRequest();
                if(!result) failedRCount++;
                pastRequests.insert({getCurrentTime(), result});

            } else {
                cout<<"reject request"<<endl;
            }
        }

    }
    void setFailState(int now) {
        lastOpenedTime = now;
        state = CBState::OPEN;
        pastRequests.clear();
        failedRCount = 0;
    }
    void setCloseState(){
        lastOpenedTime = -1;
        state = CBState::CLOSE;
    }
};
int main() {
    CircuitBreakder cb(10,50,5,5);
    
    cb.callFunction();
    cb.callFunction();
    cb.callFunction();
    cb.callFunction();
    cb.callFunction();
    cb.callFunction();
    cb.callFunction();
    cb.callFunction();
    cb.callFunction();
    cb.callFunction();
    cb.callFunction();
    // cout<<"hello word"<<endl;
    return 0;
}