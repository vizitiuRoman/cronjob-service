package models

import (
	"errors"
	"html"
	"strings"
)

type OfferModel interface {
	Prepare()
	Validate() error
}

type Offer struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

func (offer *Offer) Prepare() {
	offer.Name = html.EscapeString(strings.TrimSpace(offer.Name))
}

func (offer *Offer) Validate() error {
	if offer.Name == "" {
		return errors.New("Required Name")
	}
	return nil
}
