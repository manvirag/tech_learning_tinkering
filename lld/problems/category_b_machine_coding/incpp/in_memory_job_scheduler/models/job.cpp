#pragma once
#include<string>
#include<chrono>
#include"../util/utils.cpp"
using namespace std; 
enum class JobType {
    RECURRING, NON_RECURRING
};

enum class JobStatus {
    CREATED, SCHEDULED, SUCCESS, FAILED
};

class JobI {
    public: 
        virtual int getJobId() = 0; 
        virtual void setJobId(int id) = 0; 
        virtual string getData() = 0;
};
class ImmediateJob: public JobI {
    public: 
        int id; 
        string jobData; 
        JobType type; 
        chrono::seconds ts; 
    ImmediateJob(){}
    ImmediateJob(string jd, chrono::seconds t): jobData(jd), type(JobType::NON_RECURRING), ts(t) {}
    int getJobId()  {
        return this -> id; 
    }
    void setJobId(int id) {
        this -> id = id;
    }
    string getData(){
        return jobData;
    }
        
};


class ScheduleJob: public JobI {
    public: 
        int id; 
        string jobData; 
        JobType type; 
        int period; // in seconds;
        chrono::steady_clock::time_point nextTs; 

        ScheduleJob(){}
        ScheduleJob(string jd, int p): id(Counter::getCount()), jobData(jd), type(JobType::RECURRING), nextTs(chrono::steady_clock::now()+chrono::seconds(p)) {}
        int getJobId()  {
            return this -> id; 
        }
        void setJobId(int id) {
            this -> id = id;
        }
        string getData(){
            return jobData;
        }
        
};

class ExecutionJob {
    public: 
        int id; 
        JobI* job; 
        chrono::steady_clock::time_point ts; 
        JobStatus status; 
    ExecutionJob(){}
    ExecutionJob(JobI* jb, chrono::steady_clock::time_point t): id(Counter::getCount()), job(jb), ts(t), status(JobStatus::CREATED) {}
    int getJobId()  {
        return this -> id; 
    }
    void setJobId(int id) {
        this -> id = id;
    }
    string getData() {
        return this -> job -> getData();
    }
    void updateStatue(JobStatus status) {
        this -> status = status;
    }

        
};