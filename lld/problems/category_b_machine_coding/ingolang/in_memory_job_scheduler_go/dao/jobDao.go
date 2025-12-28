package dao

import (
	"in_memory_job_scheduler_go/models"
	"sync"
	"time"
)

type JobDao struct {
	mu                sync.RWMutex
	immediateJobStore map[int]*models.ImmediateJob
	scheduleJobStore  map[int]*models.ScheduleJob
	executionJob      map[int]*models.ExecutionJob
}

func NewJobDao() *JobDao {
	return &JobDao{
		immediateJobStore: make(map[int]*models.ImmediateJob),
		scheduleJobStore:  make(map[int]*models.ScheduleJob),
		executionJob:      make(map[int]*models.ExecutionJob),
	}
}

func (d *JobDao) CreateImmediateJob(job *models.ImmediateJob) *models.ImmediateJob {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.immediateJobStore[job.Id] = job
	return job
}

func (d *JobDao) CreateExecutionJob(job *models.ExecutionJob) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.executionJob[job.Id] = job
}

func (d *JobDao) CreateSchedule(job *models.ScheduleJob) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.scheduleJobStore[job.Id] = job
}

func (d *JobDao) GetJob(id int, jobType models.JobType) models.JobI {
	d.mu.RLock()
	defer d.mu.RUnlock()
	switch jobType {
	case models.RECURRING:
		return d.scheduleJobStore[id]
	case models.NON_RECURRING:
		return d.immediateJobStore[id]
	}
	return nil
}

func (d *JobDao) GetExJobRange(st, et time.Time) []*models.ExecutionJob {
	d.mu.RLock()
	defer d.mu.RUnlock()
	var jobs []*models.ExecutionJob
	for _, ej := range d.executionJob {
		if ej.Status == models.CREATED && !ej.Ts.Before(st) && !ej.Ts.After(et) {
			jobs = append(jobs, ej)
		}
	}
	return jobs
}

func (d *JobDao) GetScJobRange(st, et time.Time) []*models.ScheduleJob {
	d.mu.RLock()
	defer d.mu.RUnlock()
	var jobs []*models.ScheduleJob
	for _, sj := range d.scheduleJobStore {
		if !sj.NextTs.Before(st) && !sj.NextTs.After(et) {
			jobs = append(jobs, sj)
		}
	}
	return jobs
}

func (d *JobDao) UpdateExJob(ej *models.ExecutionJob) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.executionJob[ej.Id] = ej
}

func (d *JobDao) UpdateScJob(sj *models.ScheduleJob) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.scheduleJobStore[sj.Id] = sj
}
