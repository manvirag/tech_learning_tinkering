#include<iostream> 
#include<chrono>
#include<map>
#include<queue>
#include<thread>
using namespace std; 
/*


Multithreaded job scheduler



two type of job
recurring -> periodic
and non recurring.  -> at particular time only


Immediate Job
    - jobid
    - jobdata
    - time 
    - type 
Schedume Job -> curr + period -> first time .
    - jobid 
    - jobdata
    - period 
    - type 
    - nextschedule time 

Execution Job -> can be of immeidate and schedule. 
    - id 
    - jobid 
    - executiontime 
    - status


jssystem -> take job -> add in database 
    - immediate -> execution. 
organiser -> push execution -> queie
scheduler -> create execution  periodically -> of schedule job and update nexschedule 
-> 10 second
worker -> consuemr from queue , scheduler push to this queue.

fail -> log skip. 
schedule -> independent of previous 
min schedule perioud -> 20 

in case of miss -> reconsilation job -> ignore as of now. 




https://leetcode.com/discuss/post/2463807/uber-sde2-l4-bengaluru-aug-2022-accepted-9j90/

Design a multithreaded job scheduler. A job can have multiple tasks,
 and each task can have a pre-requisite task as well. Jobs should run 
 in as much parallelization as possible.


https://algomaster.io/learn/lld

https://www.youtube.com/watch?v=FcbsppIX0bg&list=PLQEaRBV9gAFvzp6XhcNFpk1WdOcyVo9qT&index=26 [IMP check c++]

https://blog.gagan93.me/low-level-design-interviews

-> Code an in-memory Job Scheduler which handles 3 types of jobs - 
Recurring and Non recurring jobs. There can be multiple jobs to be run concurrently 
 

**Problem**

Implement an InMemory Task scheduler Library that supports these functionalities:
Submit a task and a time at which the task should be executed. --> schedule(task, time)

Schedule a task at a fixed interval --> scheduleAtFixedInterval(task, interval) 
- interval is in seconds

The first instance will trigger it immediately and the next execution would
 start after interval seconds of completion of the preceding execution.

If a task has an interval of 10 seconds and submitted at 2:00 pm then

It will be executed at 2:00 pm

Once the execution is completed + 10 seconds(interval) it will trigger the next execution and so on.

Expectations

The number of worker threads should be configurable and manage them effectively.

Code/Design should be modular and follow design patterns.

Don’t use any external/internal libs that provide the same functionality 
and core APIs should be used.

Expectation was to share the approach, implement the logic and walkthrough 
demoable code by adding TCs.

- You can use IDE of your choice.

https://medium.com/@prashant558908/uber-low-level-design-interview-questions-from-recent-interviews-7035fadfcb3d
*/


#include<iostream> 
#include<queue>
#include<vector> 
#include<mutex> 
using namespace std; 


enum class JobType {
    FIXED, RECURRING
};
class Job {
  public: 
    int id ; 
    JobType type; 
    int period ; // recurring. 
    int executionTime;
    Job(){} 
    Job(int id, JobType type, int executionTime, int period = 0): id(id), type(type), executionTime(executionTime), period(period) {}
};
class Compare {
    public: 
        bool operator()(Job a, Job b) {
            return a.executionTime > b.executionTime;
        }
};
class JobScheduler {
    public: 
        priority_queue<Job, vector<Job> , Compare> jobScheduler;
        queue<Job> workerQueue;
        int workerThreads;
        bool stopWorker; 
        vector<thread> workers; 
        thread scheduler;
        mutex mu; 
        condition_variable cvW; 
        condition_variable cvS; 

        ~JobScheduler() {
            {
                lock_guard<mutex> lock(mu); // after review
                stopWorker = true; 
            }
            
            cvS.notify_all();
            cvW.notify_all();
            for(auto &t: workers){
                if(t.joinable()) t.join();
            }
            if(scheduler.joinable()) scheduler.join();
            
        }
        
        JobScheduler(int workerThreads) {
            this->workerThreads = workerThreads;
            stopWorker = false; 
        }

        void initialise() {
            for(int i=0;i<workerThreads;i++) {
                workers.emplace_back([this](){
                    Worker();
                });
            }
            scheduler = thread([this](){ Scheduler(); });
         
        }

        void Worker() {
            while(true) { // change after review since not lock were here.
                Job jb;
                {
                    unique_lock<mutex> lock(mu);
                    cvW.wait(lock, [&](){
                        return !workerQueue.empty() || stopWorker;  
                    }) ;
                    if(workerQueue.empty() && stopWorker) return;
                    jb = workerQueue.front();
                    workerQueue.pop();
                }
                cout<<"Executed Job "<<jb.id << " at time: "<<jb.executionTime<<endl;
                
            }
        }

        void Scheduler() {
            
            while(true){ // after review, !stopWorker || !jobScheduler.empty()             
                unique_lock<mutex> lock(mu);
                cvS.wait(lock, [&](){
                return !jobScheduler.empty() || stopWorker;  
                }) ;
                if(jobScheduler.empty() && stopWorker) break;
                Job jb = jobScheduler.top();
                cout<<"Schedling job "<<jb.id<<" for time "<<jb.executionTime<<endl;
                jobScheduler.pop();  
                workerQueue.push(jb);
                cvW.notify_one();
                
                

                if(jb.type == JobType::RECURRING) {
                    Job rJob = Job(jb.id, JobType::RECURRING, jb.executionTime+ jb.period, jb.period);
                    jobScheduler.push(rJob);
                }
            }
        }

        void ScheduleFixedJob(int jobId, int exeuctionTime) {
            Job jb = Job(jobId, JobType::FIXED,exeuctionTime);
            lock_guard<mutex> lock(mu);
            jobScheduler.push(jb);
            cvS.notify_one();
        }

        void ScheduleRecurringJob(int jobId, int exeuctionTime, int period) {
            Job jb = Job(jobId, JobType::RECURRING,exeuctionTime , period);
            lock_guard<mutex> lock(mu);
            jobScheduler.push(jb);
            cvS.notify_one();
        }


    
};
int main() {

    JobScheduler *js = new JobScheduler(3);
    js->initialise();
    js->ScheduleFixedJob(1, 10000);
    js->ScheduleFixedJob(2, 3000);
    js->ScheduleFixedJob(3, 14000);
    js->ScheduleFixedJob(4, 56000);
    js->ScheduleRecurringJob(10, 56000, 30);
    js->ScheduleFixedJob(5, 20000);
    js->ScheduleFixedJob(6, 10000);
    js->ScheduleFixedJob(7, 50000);
    js->ScheduleFixedJob(8, 60000);
    this_thread::sleep_for(chrono::seconds(10));
    cout<<"hello world"<<endl;
    return 0;
}

/*
review: 
- use locking properly -> no need to lock every function

*/