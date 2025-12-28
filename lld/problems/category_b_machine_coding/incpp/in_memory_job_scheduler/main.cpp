#include<iostream> 
using namespace std; 

/*

Problem
Implement an InMemory Task scheduler Library that supports these functionalities:
Submit a task and a time at which the task should be executed. --> schedule(task, time)
Schedule a task at a fixed interval --> scheduleAtFixedInterval(task, interval) - interval is in seconds
The first instance will trigger it immediately and the next execution would start after interval seconds of completion of the preceding execution.
If a task has an interval of 10 seconds and submitted at 2:00 pm then
It will be executed at 2:00 pm
Once the execution is completed + 10 seconds(interval) it will trigger the next execution and so on.
Expectations
The number of worker threads should be configurable and manage them effectively.
Code/Design should be modular and follow design patterns.
Don’t use any external/internal libs that provide the same functionality and core APIs should be used.

Expectation was to share the approach, implement the logic and walkthrough demoable code by adding TCs.

You can use IDE of your choice.



................................
0. we want to focus on writing code
1. deletion of schedule job ?? 
    -> yes 
2. delay in recurring and non recurring: acceptable ? 
    -> recurring -> 10s 
    -> non redcurring -> 10s 

3. non recurring min period ? -> 20s 

4. error -> 
    - non recurring -> error -> log skip. 
    - recurring -> error -> log skip. ? 
    - usually we put on dlq 


flow -> api -> schedule(jobData, time) 
    -> api -> scheduleRecurring(jobData, period )

5. support concurrency ? yes 

non-recurr

1. api -> save -> run -> update status. 

recurr
1. api -> save -> background cronjob -> run -> update status. 


20 s schedule 

0 20 40 60 

what if 0 got fail ? block 20 -> not 


entities: 
1. Job 
    -> id 
    -> data
    -> type 
    -> period
    -> time  // empty in case of schedule. 

2. ExecutionJob
    - id 
    - scheduledTime
    - JobId
    - status // CREATED, SCHEDULED, SUCCESS/FAILED

3. ScheduleJob
    - id
    - jobId
    - period
    - nextSchedule. 


schuduejob -> create exeution job 

20  -> 20 
update nextschedule to 40 






service 
2. JSSystem -> apis 
    -> schedule  -> add job, add exeuction job -> send for run 
    -> schedulerRecurr -> add job, -> add execution( current + period), update next schedule ( current + 2*perioud)
    -> start -> thread.  -> perioud 10s -> pic schedule hjbo with nex scheduler and push in exuection job, pick execution job and run it.
    -> executionworker -> config parallel 


dao 
1. jobdao ( above entities. )



logic 
-> immediate -> workerqueue -> workers 
-> schedule -> add in db 

-> organisaer -> period -> 10s -> backgroud 
    -> pick the job from execution  -> workder queue. 
-> scheduler -> period -> 10s -> background
    -> pick from schedulejog 
        -> add in exeuction
        -> update nextts 

-> run -> start above 2 threads. 

*/

#include"./service/jobSchedulerService.cpp"

int main() {
    JobDao* jd = new JobDao();
    JobSchdulerServiceI* jss = new JobSchdulerService(jd, 5);
    thread t4(&JobSchdulerServiceI::run, jss);
    jss -> scheduleImmediate("anubhav",chrono::seconds(30));
    jss -> scheduleImmediate("df",chrono::seconds(30));
    jss -> scheduleImmediate("af",chrono::seconds(30));
    jss -> scheduleImmediate("fsda",chrono::seconds(60));
    jss -> schedule("schedule",20);
    t4.join();
    cout<<"hello word"<<endl;
    return 0; 
}