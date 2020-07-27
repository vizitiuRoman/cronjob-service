package models

import (
	"errors"
	"fmt"
	"time"

	. "github.com/cronjob-service/pkg/client"
	"github.com/robfig/cron/v3"
)

type OfferJobModel interface {
	StartJob()
	cronJobWorker()
	removeJobByID() error
}

type OfferJob struct {
	ch           chan bool
	OfferID      int
	repeatNumb   uint8
	repeatTime   string
	repeatedNumb uint8
	offers       []byte
}

type Entry struct {
	ID   uint64    `json:"offerId"`
	Next time.Time `json:"nextDate"`
	Prev time.Time `json:"prevDate"`
}

var (
	jobIDs  = make(map[int]cron.EntryID)
	cronJob = *cron.New()
)

func init() {
	cronJob.Start()
	cronJob.Location()
}

func NewOfferJob(offerID int, repeatNumb uint8, repeatTime string, offers []byte) *OfferJob {
	return &OfferJob{
		ch:           make(chan bool),
		OfferID:      offerID,
		repeatNumb:   repeatNumb,
		repeatTime:   repeatTime,
		repeatedNumb: 0,
		offers:       offers,
	}
}

func (offerJob *OfferJob) StartJob() {
	go offerJob.cronJobWorker()
	<-offerJob.ch
	_ = offerJob.removeJobByID()
}

func (offerJob *OfferJob) cronJobWorker() {
	for {
		spec := fmt.Sprintf("* %s * * *", offerJob.repeatTime)
		cronID, err := cronJob.AddFunc(spec, func() {
			if offerJob.repeatedNumb == offerJob.repeatNumb {
				offerJob.ch <- true
				return
			}
			SendOffer(offerJob.offers)
			offerJob.repeatedNumb++
		})

		if err != nil {
			fmt.Printf("Error cronJobWorker: %v", err)
			close(offerJob.ch)
			return
		}
		jobIDs[offerJob.OfferID] = cronID

		select {
		case <-offerJob.ch:
			close(offerJob.ch)
			return
		}
	}
}

func (offerJob *OfferJob) removeJobByID() error {
	if cronID, ok := jobIDs[offerJob.OfferID]; ok {
		cronJob.Remove(cronID)
		delete(jobIDs, offerJob.OfferID)
		return nil
	}
	return errors.New("Not Found CronJob")
}

func DeleteJobByID(offerID int) error {
	if cronID, ok := jobIDs[offerID]; ok {
		cronJob.Remove(cronID)
		delete(jobIDs, offerID)
		return nil
	}
	return nil
}

func GetRunningJobs() []Entry {
	var entries []Entry
	for _, entry := range cronJob.Entries() {
		for i, id := range jobIDs {
			if entry.ID == id {
				entries = append(entries, Entry{ID: uint64(i), Next: entry.Next, Prev: entry.Prev})
			}
		}
	}
	return entries
}
