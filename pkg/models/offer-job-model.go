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
	removeJobByID()
}

type OfferJob struct {
	done         chan bool
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
		done:         make(chan bool),
		OfferID:      offerID,
		repeatNumb:   repeatNumb,
		repeatTime:   repeatTime,
		repeatedNumb: 0,
		offers:       offers,
	}
}

func (offerJob *OfferJob) StartJob() {
	go offerJob.cronJobWorker()
}

func (offerJob *OfferJob) cronJobWorker() {
	for {
		spec := fmt.Sprintf("* %s * * *", offerJob.repeatTime)
		cronID, err := cronJob.AddFunc(spec, func() {
			if offerJob.repeatedNumb == offerJob.repeatNumb {
				offerJob.done <- true
				return
			}
			SendOfferToMBB(offerJob.offers)
			offerJob.repeatedNumb++
		})
		if err != nil {
			fmt.Printf("Error cronJobWorker: %v", err)
			offerJob.done <- true
			return
		}
		jobIDs[offerJob.OfferID] = cronID
		select {
		case <-offerJob.done:
			offerJob.removeJobByID()
			close(offerJob.done)
			return
		}
	}
}

func (offerJob *OfferJob) removeJobByID() {
	if cronID, ok := jobIDs[offerJob.OfferID]; ok {
		cronJob.Remove(cronID)
		delete(jobIDs, offerJob.OfferID)
	}
}

func DeleteJobByID(offerID int) error {
	if cronID, ok := jobIDs[offerID]; ok {
		cronJob.Remove(cronID)
		delete(jobIDs, offerID)
		return nil
	}
	return errors.New(fmt.Sprintf("Not Found id %v", offerID))
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
