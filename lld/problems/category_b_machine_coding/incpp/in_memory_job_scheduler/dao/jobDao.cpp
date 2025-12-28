#include<map> 
#include<iostream> 
#include"../models/job.cpp"
#include"../util/utils.cpp"
using namespace std; 

class JobDao {
    public:  // taken decision after writing few method , sicne its timetaking. we can do this in MC.
    map<int, ImmediateJob > immediateJobStore; 
    map<int, ScheduleJob > scheduleJobStore; 
    map<int, ExecutionJob  > executionJob; 
    
    JobDao(){};

    ImmediateJob createImmediateJob(ImmediateJob job, JobType type) {
        job.setJobId(Counter::getCount());
        this-> immediateJobStore[job.getJobId()] = job;
        return job;
    }
    void createExecutionJob(ExecutionJob job) {
        this-> executionJob[job.getJobId()] = job;
    }
    void createSchedule(ScheduleJob job) {
        this-> scheduleJobStore[job.getJobId()] = job;
    }
    JobI* getJob(int id, JobType type) {
        switch (type)
        {
        case JobType::RECURRING:
            return &this-> scheduleJobStore[id];
            break;
        case JobType::NON_RECURRING:
            return &this-> immediateJobStore[id];
            break;
        default:
            break;
        }
        
    }
    vector<ExecutionJob> getExJobRange(chrono::steady_clock::time_point st , chrono::steady_clock::time_point et) {
        vector<ExecutionJob> jobs; 
        for(auto ej : this -> executionJob) {
            if(ej.second.status == JobStatus::CREATED && ej.second.ts >= st && ej.second.ts <= et) {
                    jobs.push_back(ej.second);
            }
        }
        return jobs;
    }
    vector<ScheduleJob> getScJobRange(chrono::steady_clock::time_point st , chrono::steady_clock::time_point et) {
        vector<ScheduleJob> jobs; 
        for(auto ej : this -> scheduleJobStore) {
            if( ej.second.nextTs >= st && ej.second.nextTs <= et) {
                    jobs.push_back(ej.second);
            }
        }
        cout<<jobs.size()<<endl;
        return jobs;
    }
    void updateExJob(ExecutionJob ej) {
        this -> executionJob[ej.id] = ej;
    }
    void updateScJob(ScheduleJob sj) {
        this -> scheduleJobStore[sj.id] = sj;
    }
};

