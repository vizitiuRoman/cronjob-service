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
	ID        uint64      `json:"id"`
	Title     string      `json:"title"`
	Category  string      `json:"category"`
	Banners   string      `json:"banners"`
	Actions   []action    `json:"actions"`
	OfferData offerData   `json:"offerData"`
	Companies []companies `json:"companies"`
	Template  template    `json:"Template"`
	ImgSrc    string      `json:"imgSrc"`
}

type offerData struct {
	StartDate      string
	EndDate        string
	CronExpression string
	RepeatNumb     uint8
	RepeatTime     string
}

type companies struct {
	Idno            string
	CompaniesOffers map[string]struct {
		Email  string
		Sum    string
		Period string
	}
}

type action struct {
	I18n struct {
		Ro string
		Ru string
		En string
	}
	Value string
}

type template struct {
	Name   string
	Type   string
	Schema string
}

func (offer *Offer) Prepare() {
	offer.Title = html.EscapeString(strings.TrimSpace(offer.Title))
	offer.Category = html.EscapeString(strings.TrimSpace(offer.Category))
	offer.Banners = html.EscapeString(strings.TrimSpace(offer.Banners))
	offer.ImgSrc = html.EscapeString(strings.TrimSpace(offer.ImgSrc))
}

func (offer *Offer) Validate() error {
	if offer.ID == 0 {
		return errors.New("Required ID")
	}
	if offer.Title == "" {
		return errors.New("Required Title")
	}
	if offer.Category == "" {
		return errors.New("Required Category")
	}
	if offer.ImgSrc == "" {
		return errors.New("Required Image")
	}
	if offer.OfferData.StartDate == "" {
		return errors.New("Required Image")
	}
	if offer.OfferData.EndDate == "" {
		return errors.New("Required Image")
	}
	if offer.OfferData.CronExpression == "" {
		return errors.New("Required CronExression")
	}
	if offer.OfferData.RepeatTime == "" {
		return errors.New("Required RepeatTime")
	}
	if offer.Template.Name == "" {
		return errors.New("Required Template Name")
	}
	if offer.Template.Type == "" {
		return errors.New("Required Template Type")
	}
	if offer.Template.Schema == "" {
		return errors.New("Required Template Schema")
	}
	for _, company := range offer.Companies {
		if company.Idno == "" {
			return errors.New("Required Company Idno")
		}
		if _, ok := company.CompaniesOffers["data"]; !ok {
			return errors.New("Required CompaniesOffers")
		}
		if company.CompaniesOffers["data"].Email == "" {
			return errors.New("Required CompaniesOffers email")
		}
		if company.CompaniesOffers["data"].Sum == "" {
			return errors.New("Required CompaniesOffers Sum")
		}
		if company.CompaniesOffers["data"].Period == "" {
			return errors.New("Required CompaniesOffers Period")
		}
	}
	return nil
}
