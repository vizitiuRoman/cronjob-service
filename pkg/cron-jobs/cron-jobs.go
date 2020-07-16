package cron_jobs

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

type Job struct {
	OfferJob func(time.Duration, string, chan bool, chan uint32)
}

var (
	cronJob = gocron.NewScheduler()
	jobs    = make(map[uint32]Job)
)

func sendOffer(name string) {
	fmt.Println("Offer Name", name)
}

func worker(offerID chan uint32) {
	for {
		select {
		case id := <-offerID:
			fmt.Println("Offer ID", id)
		}
	}
}

func RunJob(id uint32, name string) {
	ch := make(chan bool)
	offerID := make(chan uint32)

	jobs[id] = struct {
		OfferJob func(time.Duration, string, chan bool, chan uint32)
	}{
		OfferJob: offerJob,
	}

	go offerJob(time.Second*3, name, ch, offerID)

	fmt.Println(jobs)
	<-ch
	fmt.Println(<-ch)
}

func offerJob(t time.Duration, name string, ch chan bool, offerID chan uint32) {
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
