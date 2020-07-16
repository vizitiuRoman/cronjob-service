package cron_jobs

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

type Job struct {
	sendOffer func(string)
}

var (
	cronJob = gocron.NewScheduler()
	jobs    = make(map[uint32]Job)
)

func sendOffer(name string) {
	fmt.Println("Offer Name", name)
}

func offerJob(t time.Duration, name string, ch chan bool) {
	for {
		next := time.Now().Add(t)
		if next.Before(time.Now()) {
			next = next.Add(0)
		}
		first := time.After(next.Sub(time.Now()))

		go func() {
			cronJob.Every(1).Second().Do(sendOffer, name)

			fmt.Println(cronJob.Jobs())
			fmt.Println(cronJob.Len())
		}()
		select {
		case <-first:
			close(ch)
			return
		case <-cronJob.Start():
		}
	}
}

func RunJob(id uint32, name string) {
	ch := make(chan bool)

	jobs[id] = struct{ sendOffer func(string) }{sendOffer: sendOffer}

	go offerJob(time.Second*100, name, ch)

	<-ch
	job, _ := jobs[id]
	delete(jobs, id)
	cronJob.Remove(job.sendOffer)

	fmt.Println(cronJob.Jobs())
	fmt.Println(cronJob.Len())
}

func StopJob() {
	job, _ := jobs[10]
	delete(jobs, 10)
	cronJob.Remove(job.sendOffer)
}
