#include<iostream> 
#include<chrono>
#include<map>
#include<queue>
#include<thread>
using namespace std; 
/*

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


*/

using TimePoint = chrono::steady_clock::time_point;
using Duration = chrono::seconds ;
using Clock = chrono::steady_clock;

enum class JobType {
    RECURRING, NON_RECURRING
};
enum class JobStatus {
    CREATED, SCHEDULED, FAILED, SUCCESS
};

class Counter{
    public:
     static int count;
     static int getId(){
        count++;return count;
     }
};
int Counter::count = 0; 
class JobI {
    public: 
        virtual int getId()  = 0;
};
class NonRecurringJob {
    public: 
      int id; 
      string jobData; 
      JobType type; 
      TimePoint ts; 
      NonRecurringJob(){}
      NonRecurringJob(string jd, TimePoint t): id(Counter::getId()),jobData(jd), type(JobType::NON_RECURRING), ts(t){}


};
class RecurringJob {
    public: 
        int id; 
        string jobData; 
        JobType type; 
        int period;  // in seconds. 
        TimePoint nextTs;
        RecurringJob(){}
        RecurringJob(string jd, int p): id(Counter::getId()),jobData(jd), type(JobType::RECURRING), period(p), nextTs(Clock::now() + Duration(period)){}
};

class ExecutionJob {
    public: 
        int id; 
        int jobId; 
        TimePoint sts; 
        JobType type;
        JobStatus status;
        ExecutionJob(){}
        ExecutionJob(int jobId, TimePoint ts, JobType ty): id(Counter::getId()), jobId(jobId), sts(ts), status(JobStatus::CREATED), type(ty){}
};
class JSSystem {
    public: 
        map<int, NonRecurringJob > nrJob;
        map<int, RecurringJob > rJob;
        map<int, ExecutionJob> eJob;
        int workerThreads;
        const int period = 5;
        queue<ExecutionJob> workerJobQueue;
    JSSystem(){}
    JSSystem(int wt): workerThreads(wt){}

    void organiser() { // execution to queue
        while(true) {
            TimePoint et = Clock::now();
            TimePoint st = Clock::now() - Duration(this->period);
            for(auto e: this->eJob) {
                if(e.second.sts >= st && e.second.sts <= et && e.second.status == JobStatus::CREATED){
                    this -> workerJobQueue.push(e.second);
                    this -> eJob[e.second.id].status = JobStatus::SCHEDULED; 
                }
            }
            this_thread::sleep_for(Duration(this->period));
        }
    }

    void scheduler() { // execution to queue
        while(true) {
            TimePoint et = Clock::now();
            TimePoint st = Clock::now() - Duration(this->period);
            for(auto e: this->rJob) {
                if(e.second.nextTs >= st && e.second.nextTs <= et){
                    ExecutionJob ej = ExecutionJob(e.second.id, e.second.nextTs, JobType::RECURRING);
                    ej.status = JobStatus::SCHEDULED;
                    this -> eJob[ej.id]=ej; 
                    this -> workerJobQueue.push(ej);
                    this -> rJob[e.second.id].nextTs = this -> rJob[e.second.id].nextTs + Duration(e.second.period);
                }
            }
            this_thread::sleep_for(Duration(this->period));
        }
    }

    void worker() {
        while(true) {

            if(this->workerJobQueue.size() !=0 ){
                ExecutionJob ej = this-> workerJobQueue.front();
                this->workerJobQueue.pop();

                cout<<ej.jobId<<" ";
                if(ej.type == JobType::RECURRING) {
                    cout<<this->rJob[ej.jobId].jobData<<endl;
                } else {
                    cout<<this->nrJob[ej.jobId].jobData<<endl;
                }
            } 
            this_thread::sleep_for(Duration(this->period));
        }
    }


    // apis 
    void nonRecurringSchedule(NonRecurringJob nr) {
        this -> nrJob[nr.id]  = nr; 
        ExecutionJob ej = ExecutionJob(nr.id, nr.ts, JobType::NON_RECURRING);
        this -> eJob[ej.id] = ej;
    }
    void recurringSchedule(RecurringJob rj) {
        this -> rJob[rj.id] = rj;
    }

    void run() {
        thread t2(&JSSystem::organiser, this);
        thread t3(&JSSystem::scheduler, this);
        thread t4(&JSSystem::worker, this);
        t2.join();
        t3.join();
        t4.join();
    }

};

int main() {
    JSSystem js = JSSystem();
    NonRecurringJob nr = NonRecurringJob("nr", Clock::now() + Duration(30));
    NonRecurringJob nr2 = NonRecurringJob("nr2", Clock::now() + Duration(20));
    NonRecurringJob nr3 = NonRecurringJob("nr3", Clock::now() + Duration(10));
    RecurringJob r = RecurringJob("r", 10);
    js.nonRecurringSchedule(nr);
    js.nonRecurringSchedule(nr2);
    js.nonRecurringSchedule(nr3);
    js.recurringSchedule(r);
    thread t1(&JSSystem::run, &js);
    t1.join();
    cout<<"hello js system "<<endl;
    return 0;
}

