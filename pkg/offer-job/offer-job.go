package offer_job

import (
	"errors"
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

func init() {
	cronJobs.Start()
}

func StartJob(offerID int, name string, startDate, endDate time.Duration) {
	offerJob := &OfferJob{
		ch:        make(chan int),
		offerID:   offerID,
		name:      name,
		startDate: startDate,
		endDate:   endDate,
	}

	go cronJobWorker(offerJob)

	<-offerJob.ch
	fmt.Println("Started jobs", len(cronJobs.Entries()))
	_ = removeJobByID(offerID)
}

func DeleteJobByID(offerID int) error {
	err := removeJobByID(offerID)
	if err != nil {
		return err
	}
	fmt.Println("Started jobs", len(cronJobs.Entries()))
	return nil
}

func GetRunningJobs() int {
	return len(cronJobs.Entries())
}

func cronJobWorker(offerJob *OfferJob) {
	for {
		next := time.Now().Add(time.Second * 10)
		if next.Before(time.Now()) {
			next = next.Add(0)
		}
		first := time.After(next.Sub(time.Now()))

		go func() {
			cronID, err := cronJobs.AddFunc("@every 0h0m1s", func() {
				fmt.Println("Offer job id", offerJob.offerID)
			})
			if err != nil {
				fmt.Printf("cronJobWorker error: %v", err)
				close(offerJob.ch)
				return
			}

			jobIDs[offerJob.offerID] = cronID

			fmt.Println("Started jobs", len(cronJobs.Entries()))
		}()
		select {
		case <-first:
			close(offerJob.ch)
			return
		}
	}
}

func removeJobByID(offerID int) error {
	if cronID, ok := jobIDs[offerID]; ok {
		cronJobs.Remove(cronID)
		delete(jobIDs, offerID)
		return nil
	}
	return errors.New("Not Found")
}
