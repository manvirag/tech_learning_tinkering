package models

import (
	"time"
	"in_memory_job_scheduler_go/util"
)

type JobType int

const (
	RECURRING JobType = iota
	NON_RECURRING
)

type JobStatus int

const (
	CREATED JobStatus = iota
	SCHEDULED
	SUCCESS
	FAILED
)

type JobI interface {
	GetJobId() int
	SetJobId(int)
	GetData() string
}

type ImmediateJob struct {
	Id      int
	JobData string
	Type    JobType
	Ts      time.Duration
}

func NewImmediateJob(jobData string, ts time.Duration) *ImmediateJob {
	return &ImmediateJob{
		Id:      util.GetCount(),
		JobData: jobData,
		Type:    NON_RECURRING,
		Ts:      ts,
	}
}

func (j *ImmediateJob) GetJobId() int { return j.Id }
func (j *ImmediateJob) SetJobId(id int) { j.Id = id }
func (j *ImmediateJob) GetData() string { return j.JobData }

type ScheduleJob struct {
	Id      int
	JobData string
	Type    JobType
	Period  int // in seconds
	NextTs  time.Time
}

func NewScheduleJob(jobData string, period int) *ScheduleJob {
	return &ScheduleJob{
		Id:      util.GetCount(),
		JobData: jobData,
		Type:    RECURRING,
		Period:  period,
		NextTs:  time.Now().Add(time.Duration(period) * time.Second),
	}
}

func (j *ScheduleJob) GetJobId() int { return j.Id }
func (j *ScheduleJob) SetJobId(id int) { j.Id = id }
func (j *ScheduleJob) GetData() string { return j.JobData }

type ExecutionJob struct {
	Id     int
	Job    JobI
	Ts     time.Time
	Status JobStatus
}

func NewExecutionJob(job JobI, ts time.Time) *ExecutionJob {
	return &ExecutionJob{
		Id:     util.GetCount(),
		Job:    job,
		Ts:     ts,
		Status: CREATED,
	}
}

func (j *ExecutionJob) GetJobId() int { return j.Id }
func (j *ExecutionJob) SetJobId(id int) { j.Id = id }
func (j *ExecutionJob) GetData() string { return j.Job.GetData() }
func (j *ExecutionJob) UpdateStatus(status JobStatus) { j.Status = status }

