#include"../dao/jobDao.cpp" 
#include<chrono> 
#include<iostream>
#include<queue> 
#include<mutex>
#include<thread>
using namespace std; 



class JobSchdulerServiceI {
    public: 
        virtual void scheduleImmediate(string jobData, chrono::seconds ts ) = 0;  // later on error handling
        virtual void schedule(string jobData, int period ) = 0; 
        virtual void run() = 0;
};


class JobSchdulerService: public JobSchdulerServiceI {
    JobDao* jobDao; 
    int workerThreads; 
    mutex mu; 
    condition_variable cv;
    queue<ExecutionJob> workerQueue; 
    const int period = 20;
    void organiser() {
        while(true) {
            cout<<"........................"<<endl;
            cout<<"organiser started"<<endl;
            chrono::steady_clock::time_point st = chrono::steady_clock::now();
            chrono::steady_clock::time_point et = chrono::steady_clock::now() + chrono::seconds(period);
            vector<ExecutionJob> jobs =  this -> jobDao -> getExJobRange(st, et);
            for(ExecutionJob job: jobs) {
                job.updateStatue(JobStatus::SCHEDULED);   
                
                {
                    lock_guard<mutex> lock(mu);
                    this -> workerQueue.push(job);

                }
                this -> jobDao -> updateExJob(job);
                cv.notify_all();
            }
            cout<<"organiser sleeping"<<endl;
            cout<<"........................"<<endl;
            this_thread::sleep_for(chrono::seconds(10));
        }
        

    }
    void scheduler() {
        while(true){
            cout<<"........................"<<endl;
            cout<<"scheduler started"<<endl;
            chrono::steady_clock::time_point st = chrono::steady_clock::now();
            chrono::steady_clock::time_point et = chrono::steady_clock::now() + chrono::seconds(period);
            vector<ScheduleJob> jobs =  this -> jobDao -> getScJobRange(st, et);
            for(ScheduleJob job: jobs) {
                JobI* jobPtr = this->jobDao->getJob(job.getJobId(), JobType::RECURRING);
                ExecutionJob ej = ExecutionJob(jobPtr, job.nextTs);
                job.nextTs = job.nextTs + chrono::seconds(this -> period);
                this -> jobDao -> updateScJob(job);
                this -> jobDao -> updateExJob(ej);
            }
            cout<<"scheduler sleeping"<<endl;
            cout<<"........................"<<endl;
            this_thread::sleep_for(chrono::seconds(10));
        }
        
    }
    void worker() {
        while(true){
            unique_lock<mutex> lock(mu);
            cout<<"........................"<<endl;
            cout<<"worker waiting"<<endl;
            cv.wait(lock, [this](){ return this -> workerQueue.size() > 0 ;});
            cout<<"worker starting"<<endl;
            if (!this->workerQueue.size()) continue;
            ExecutionJob j = this -> workerQueue.front();
            this -> workerQueue.pop();
            cout<<"working execution job with executionId "<<j.getJobId()<<" "<<j.getData()<<endl; 
            j.status  = JobStatus::SUCCESS;
            this -> jobDao->updateExJob(j);
            cout<<"........................"<<endl;
        }
        
    }
    public: 
        JobSchdulerService(){}
        JobSchdulerService(JobDao* jd, int wt): jobDao(jd), workerThreads(wt) {}
        void scheduleImmediate(string jobData, chrono::seconds ts ) override {
            ImmediateJob cj = this -> jobDao->createImmediateJob(ImmediateJob(jobData, ts), JobType::NON_RECURRING);
            chrono::steady_clock::time_point ct = chrono::steady_clock::now() + chrono::seconds(ts);
            // auto seconds_since_epoch = chrono::duration_cast<chrono::seconds>(ct.time_since_epoch()).count();
            // cout<<seconds_since_epoch<<endl;
            JobI* jobPtr = this->jobDao->getJob(cj.getJobId(), JobType::NON_RECURRING);
            ExecutionJob ej = ExecutionJob(jobPtr, ct);
            this -> jobDao -> createExecutionJob(ej);
            return ;
        }   
        void schedule(string jobData,  int period  ) override{
            ScheduleJob sj = ScheduleJob(jobData, period);
            this -> jobDao -> createSchedule(sj);
            return ; 
        }

        void run() override{
            thread t1(&JobSchdulerService::organiser, this);
            thread t2(&JobSchdulerService::scheduler, this);
            thread t3(&JobSchdulerService::worker, this);
            t1.join();
            t2.join();
            t3.join();
        }
};

