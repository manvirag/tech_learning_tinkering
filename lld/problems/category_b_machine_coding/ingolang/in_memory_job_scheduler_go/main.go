package main

import (
	"context"
	"fmt"
	"in_memory_job_scheduler_go/dao"
	"in_memory_job_scheduler_go/service"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
	}()

	jd := dao.NewJobDao()
	jss := service.NewJobSchedulerService(jd, 5)

	go jss.Run(ctx)

	// Give scheduler time to start
	time.Sleep(100 * time.Millisecond)

	jss.ScheduleImmediate("anubhav", 30*time.Second)
	jss.ScheduleImmediate("df", 30*time.Second)
	jss.ScheduleImmediate("af", 30*time.Second)
	jss.ScheduleImmediate("fsda", 60*time.Second)
	jss.Schedule("schedule", 20)

	// Keep main alive to see output
	time.Sleep(2 * time.Minute)
	fmt.Println("hello world")
}
