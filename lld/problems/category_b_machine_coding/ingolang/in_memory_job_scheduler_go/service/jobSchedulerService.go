package service

import (
	"context"
	"fmt"
	"sync"
	"time"
	"in_memory_job_scheduler_go/dao"
	"in_memory_job_scheduler_go/models"
)

type JobSchedulerServiceI interface {
	ScheduleImmediate(jobData string, ts time.Duration)
	Schedule(jobData string, period int)
	Run(ctx context.Context)
}

type JobSchedulerService struct {
	jobDao        *dao.JobDao
	workerThreads int
	workerQueue   chan *models.ExecutionJob
	period        time.Duration
}

func NewJobSchedulerService(jd *dao.JobDao, workerThreads int) *JobSchedulerService {
	return &JobSchedulerService{
		jobDao:        jd,
		workerThreads: workerThreads,
		workerQueue:   make(chan *models.ExecutionJob, 100),
		period:        20 * time.Second,
	}
}

func (s *JobSchedulerService) organiser(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("organiser started")
			st := time.Now()
			et := st.Add(s.period)
			jobs := s.jobDao.GetExJobRange(st, et)
			
			for _, job := range jobs {
				job.UpdateStatus(models.SCHEDULED)
				s.jobDao.UpdateExJob(job)
				select {
				case s.workerQueue <- job:
				case <-ctx.Done():
					return
				}
			}
			fmt.Println("organiser sleeping")
		}
	}
}

func (s *JobSchedulerService) scheduler(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("scheduler started")
			st := time.Now()
			et := st.Add(s.period)
			jobs := s.jobDao.GetScJobRange(st, et)
			
			for _, job := range jobs {
				if jobPtr := s.jobDao.GetJob(job.Id, models.RECURRING); jobPtr != nil {
					ej := models.NewExecutionJob(jobPtr, job.NextTs)
					job.NextTs = job.NextTs.Add(s.period)
					s.jobDao.UpdateScJob(job)
					s.jobDao.UpdateExJob(ej)
				}
			}
			fmt.Println("scheduler sleeping")
		}
	}
}

func (s *JobSchedulerService) worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case job := <-s.workerQueue:
			fmt.Printf("working execution job id=%d data=%s\n", job.Id, job.GetData())
			job.UpdateStatus(models.SUCCESS)
			s.jobDao.UpdateExJob(job)
		}
	}
}

func (s *JobSchedulerService) ScheduleImmediate(jobData string, ts time.Duration) {
	cj := s.jobDao.CreateImmediateJob(models.NewImmediateJob(jobData, ts))
	ej := models.NewExecutionJob(cj, time.Now().Add(ts))
	s.jobDao.CreateExecutionJob(ej)
}

func (s *JobSchedulerService) Schedule(jobData string, period int) {
	s.jobDao.CreateSchedule(models.NewScheduleJob(jobData, period))
}

func (s *JobSchedulerService) Run(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(3)
	
	go func() {
		defer wg.Done()
		s.organiser(ctx)
	}()
	
	go func() {
		defer wg.Done()
		s.scheduler(ctx)
	}()
	
	go func() {
		defer wg.Done()
		s.worker(ctx)
	}()
	
	wg.Wait()
}

