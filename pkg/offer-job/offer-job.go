package offer_job

import (
	"errors"
	"fmt"

	. "github.com/cronjob-service/pkg/client"
	"github.com/robfig/cron/v3"
)

var (
	jobIDs   = make(map[int]cron.EntryID)
	cronJobs = *cron.New()
)

type OfferJob struct {
	ch         chan int
	offerID    int
	repeatNumb uint8
	repeatTime string
}

func init() {
	cronJobs.Start()
	cronJobs.Location()
}

func NewOfferJob(offerID int, repeatNumb uint8, repeatTime string) *OfferJob {
	return &OfferJob{
		ch:         make(chan int),
		offerID:    offerID,
		repeatNumb: repeatNumb,
		repeatTime: repeatTime,
	}
}

func StartJob(offerJob *OfferJob, offers []byte) {
	go cronJobWorker(offerJob, offers)
	<-offerJob.ch
	_ = removeJobByID(offerJob.offerID)
}

func DeleteJobByID(offerID int) error {
	err := removeJobByID(offerID)
	if err != nil {
		return err
	}
	return nil
}

func GetRunningJobs() int {
	return len(cronJobs.Entries())
}

func cronJobWorker(offerJob *OfferJob, offers []byte) {
	for {
		var repeated uint8 = 0
		limitCh := make(chan bool)

		spec := fmt.Sprintf("* %s * * *", offerJob.repeatTime)
		cronID, err := cronJobs.AddFunc(spec, func() {
			if repeated == offerJob.repeatNumb {
				limitCh <- true
				return
			}
			SendOffer(offers)
			repeated++
		})

		if err != nil {
			fmt.Printf("Error cronJobWorker: %v", err)
			close(offerJob.ch)
			return
		}
		jobIDs[offerJob.offerID] = cronID

		select {
		case <-limitCh:
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
