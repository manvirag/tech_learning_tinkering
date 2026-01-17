/*

- Design a **Train-Platform Management System** with functionalities:
    - Assign trains to platforms based on input.
    - Query which train is at a given platform at a specific time.
    - Query which platform a train is at, at a specific time.
- **Expectations:**
    - Write runnable code (in my preferred language – chose **Java**).
    - Explain design patterns used.
    - Walkthrough every step while coding.
- **My Approach:**
    - Designed classes: Train, Platform, Scheduler, Schedule Manager.
    - coded using minHeap Startegy and then explained random Assigning Startegy , wrote extensible code
    - Focused on **time-based queries** and **mapping train schedules**.
    - Did not added all class variables , just explained and just added variables requried for getting answer ( for example , for trian I jus added trian id , not capacity or type )
    - The interviewer continuously cross-questioned choices and pushed for **clean design + working code in 60–75 minutes**.


p1 -> t1 t2 t3 ( sorted time also assume train assign)

p1 
-> tr1 (st1, end2)
-> tr2 ) end2+1,end3) 



p1 -> end2+1 -> tr2 
t1 -> end2+1 => p1 

unique platform and unique train. 

models 
train
platform
trainState -> train platformid, state, endtime. 
pllatformState -> platform trainId, state, endtime. 




Trainplatformmanagerment
set<trainState>
set<pllatformState> - to fin sate


*/
#include<iostream> 
#include<set> 
#include<map> 
#include<vector> 
#include<utility> 

using namespace std; 

class Train {
    public: 
        int id ;
    Train() {}
    Train(int id): id(id) {}
};

class Platform {
    public: 
        int id ;
    Platform() {}
    Platform(int id): id(id) {}
};
class TrainPlatformState {
    public: 
        Train train; 
        set<pair<pair<int,int>, int>> trainPlatformStates; // start time, endtime, platformid // can be change to different type
    TrainPlatformState() {}
    TrainPlatformState(Train train): train(train) {}
    
    void addPlatform(int platformId, int startTime ,int endTime) {
        trainPlatformStates.insert({{startTime, endTime}, platformId});
        cout<<endl;
        for(auto x: trainPlatformStates){
            cout<<x.first.first<<" "<<x.first.second<<" "<<x.second<<endl;
        }
        cout<<endl;
    } 
    
    
};
class PlatformTrainState {
    public: 
        Platform platform; 
        set<pair<pair<int,int>, int>> platformTrainStates; // start time, endtime, platformid
    PlatformTrainState() {}
    PlatformTrainState(Platform platform): platform(platform) {}
    void addTrain(int trainId, int startTime ,int endTime) {
        platformTrainStates.insert({{startTime, endTime}, trainId});
        cout<<endl;
        for(auto x: platformTrainStates){
            cout<<x.first.first<<" "<<x.first.second<<" "<<x.second<<endl;
        }
        cout<<endl;
    }
};
class TrainManagenentSystem {
  public: 
     map<int, TrainPlatformState> trainVsPlatfoms;
     map<int, PlatformTrainState> platformVsTrains;

   void AssignTrain(int platformId, int trainId, int startTime, int endTime){ // disjion interval on a platform and for a train. 
        if(trainVsPlatfoms.find(trainId) == trainVsPlatfoms.end()) {
                Train train = Train(trainId);
                TrainPlatformState trainPlatformState = TrainPlatformState(train);     
                trainVsPlatfoms[trainId] = trainPlatformState;
        }
        if(platformVsTrains.find(platformId) == platformVsTrains.end()) {
                Platform platform = Platform(platformId);
                PlatformTrainState platformTrainStates = PlatformTrainState(platform);     
                platformVsTrains[platformId] = platformTrainStates;
        }
        trainVsPlatfoms[trainId].addPlatform(platformId,  startTime,  endTime);
        platformVsTrains[platformId].addTrain(trainId,  startTime,  endTime);
   }
   int getTrain(int platformId, int ts){
        if(platformVsTrains.find(platformId)!= platformVsTrains.end()) {
             auto ub = platformVsTrains[platformId].platformTrainStates.lower_bound({{ts, -1}, -1});
             
             if(ub != platformVsTrains[platformId].platformTrainStates.end()){
                   
                if(ub ->first.first == ts) {
                    return ub->second; 
                }
                
                auto firstP =  platformVsTrains[platformId].platformTrainStates.begin();
                if(ub!=firstP) {
                    ub--;
                    
                    if(ts >= ub->first.first && ts <= ub->first.second) {
                        return ub->second;
                    }
                }
                
             } else if(ub != platformVsTrains[platformId].platformTrainStates.begin()){
                    ub--;
                    if(ts >= ub->first.first && ts <= ub->first.second) {
                        return ub->second;
                    }
             }
        }
        return -1;
   }
   int getPlatform(int trainId, int ts){
         if(trainVsPlatfoms.find(trainId)!= trainVsPlatfoms.end()) {
             auto ub = trainVsPlatfoms[trainId].trainPlatformStates.lower_bound({{ts, -1}, -1});
             
             if(ub != trainVsPlatfoms[trainId].trainPlatformStates.end()){
                   
                if(ub ->first.first == ts) {
                    return ub->second; 
                }
                
                auto firstP =  trainVsPlatfoms[trainId].trainPlatformStates.begin();
                if(ub!=firstP) {
                    ub--;
                    
                    if(ts >= ub->first.first && ts <= ub->first.second) {
                        return ub->second;
                    }
                }
                
             } else if(ub != trainVsPlatfoms[trainId].trainPlatformStates.begin()){
                    ub--;
                    if(ts >= ub->first.first && ts <= ub->first.second) {
                        return ub->second;
                    }
             }
        }
        return -1;
   }
    
};

int main() {
    TrainManagenentSystem *tms = new TrainManagenentSystem();
    tms->AssignTrain(1, 11, 100, 200);
    tms->AssignTrain(1, 11, 301, 400);
    tms->AssignTrain(2, 11, 201, 300);
    cout<<tms->getTrain(1, 105)<<endl;
    cout<<tms->getTrain(2, 105)<<endl;
    cout<<tms->getTrain(2, 200)<<endl;
    cout<<tms->getTrain(2, 201)<<endl;
    cout<<tms->getTrain(2, 300)<<endl;
    cout<<tms->getTrain(2, 250)<<endl;
    cout<<tms->getTrain(2, 320)<<endl;
    cout<<tms->getPlatform(11, 5)<<endl;
    cout<<tms->getPlatform(11, 250)<<endl;
    
    return 0; 
}