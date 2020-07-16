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

func offerJob(t time.Duration, ch chan bool, name string, offerID uint32) {
	for {
		next := time.Now().Add(t)
		if next.Before(time.Now()) {
			next = next.Add(0)
		}
		first := time.After(next.Sub(time.Now()))

		go func() {
			cronJob.Every(1).Second().Do(jobs[offerID].sendOffer, name)

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

func RunJob(offerID uint32, name string) {
	ch := make(chan bool)

	jobs[offerID] = struct{ sendOffer func(string) }{sendOffer: sendOffer}

	go offerJob(time.Hour*2, ch, name, offerID)

	<-ch
	job, _ := jobs[offerID]
	delete(jobs, offerID)
	cronJob.Remove(job.sendOffer)

	fmt.Println(cronJob.Jobs())
	fmt.Println(cronJob.Len())
}

func StopJob(offerID uint32) {
	fmt.Println(jobs)
	fmt.Println(cronJob.Jobs())
	fmt.Println(cronJob.Len())

	if job, ok := jobs[offerID]; ok {
		cronJob.Remove(job.sendOffer)
		delete(jobs, offerID)
	}

	fmt.Println("------------")

	fmt.Println(jobs)
	fmt.Println(cronJob.Jobs())
	fmt.Println(cronJob.Len())
}
