package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type OfferModel interface {
	Prepare()
	Validate() error
}

type Offer struct {
	ID        uint64        `json:"id"`
	Name      string        `json:"name"`
	StartDate time.Duration `json:"startDate"`
	EndDate   time.Duration `json:"endDate"`
}

func (offer *Offer) Prepare() {
	offer.Name = html.EscapeString(strings.TrimSpace(offer.Name))
}

func (offer *Offer) Validate() error {
	if offer.Name == "" {
		return errors.New("Required Name")
	}
	if offer.StartDate == 0 {
		return errors.New("Required EndDate")
	}
	if offer.EndDate == 0 {
		return errors.New("Required EndDate")
	}
	return nil
}
