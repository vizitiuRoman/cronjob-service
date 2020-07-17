package offer_job

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

var (
	jobIDs   = make(map[int]cron.EntryID)
	cronJobs = *cron.New()
)

type OfferJob struct {
	ch        chan int
	offerID   int
	name      string
	startDate time.Duration
	endDate   time.Duration
}

func RunJob(offerID int, name string, startDate, endDate time.Duration) {
	offerJob := &OfferJob{
		ch:        make(chan int),
		offerID:   offerID,
		name:      name,
		startDate: startDate,
		endDate:   endDate,
	}
	go cronJobWorker(offerJob)
	cronJobs.Start()

	<-offerJob.ch
	removeJobByID(offerID)
}

func StopJob(offerID int) {
	removeJobByID(offerID)
}

func cronJobWorker(offerJob *OfferJob) {
	for {
		next := time.Now().Add(time.Hour * 3)
		if next.Before(time.Now()) {
			next = next.Add(0)
		}
		first := time.After(next.Sub(time.Now()))

		go func() {
			cronID, _ := cronJobs.AddFunc("@every 0h0m1s", func() {
			})
			jobIDs[offerJob.offerID] = cronID

			fmt.Println("Jobs worker", len(cronJobs.Entries()))
		}()
		select {
		case <-first:
			close(offerJob.ch)
			return
		}
	}
}

func removeJobByID(offerID int) {
	if cronID, ok := jobIDs[offerID]; ok {
		cronJobs.Remove(cronID)
		delete(jobIDs, offerID)
	}

	fmt.Println(len(cronJobs.Entries()))
}
