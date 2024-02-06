package scheduler

import (
	"time"
)

type Job interface {
	Execute()
}

type JobScheduler struct {
	JobQueue chan Job
	Interval time.Duration
}

func NewJobScheduler(interval time.Duration) *JobScheduler {
	return &JobScheduler{
		JobQueue: make(chan Job),
		Interval: interval,
	}
}

func (s *JobScheduler) Start() {
	go func() {
		ticker := time.NewTicker(s.Interval)

		for {
			select {
			case job := <-s.JobQueue:
				job.Execute()
			case <-ticker.C:
				for job := range s.JobQueue {
					job.Execute()
				}
			}
		}
	}()
}

func (s *JobScheduler) ScheduleOnce(duration time.Duration, job Job) {
	go func() {
		time.Sleep(duration)
		s.JobQueue <- job
	}()
}
