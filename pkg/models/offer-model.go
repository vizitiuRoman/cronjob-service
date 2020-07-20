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
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	CronExpression string `json:"cronExpression"`
	RepeatNumb     uint8  `json:"repeatNumb"`
	RepeatTime     string `json:"repeatTime"`
}

type companies struct {
	Idno            string `json:"idno"`
	CompaniesOffers map[string]struct {
		Email  string `json:"email"`
		Sum    string `json:"sum"`
		Period string `json:"period"`
	}
}

type action struct {
	I18n struct {
		Ro string `json:"ro"`
		Ru string `json:"ru"`
		En string `json:"en"`
	} `json:"i18n"`
	Value string `json:"value"`
}

type template struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Schema string `json:"schema"`
}

func (offer *Offer) Prepare() {
	offer.Title = html.EscapeString(strings.TrimSpace(offer.Title))
	offer.Category = html.EscapeString(strings.TrimSpace(offer.Category))
	offer.Banners = html.EscapeString(strings.TrimSpace(offer.Banners))
	offer.ImgSrc = html.EscapeString(strings.TrimSpace(offer.ImgSrc))
	offer.OfferData.StartDate = html.EscapeString(strings.TrimSpace(offer.OfferData.StartDate))
	offer.OfferData.EndDate = html.EscapeString(strings.TrimSpace(offer.OfferData.EndDate))
	offer.OfferData.CronExpression = html.EscapeString(strings.TrimSpace(offer.OfferData.CronExpression))
	offer.OfferData.RepeatTime = html.EscapeString(strings.TrimSpace(offer.OfferData.RepeatTime))
	offer.Template.Name = html.EscapeString(strings.TrimSpace(offer.Template.Name))
	offer.Template.Type = html.EscapeString(strings.TrimSpace(offer.Template.Type))
	offer.Template.Schema = html.EscapeString(strings.TrimSpace(offer.Template.Schema))
	for i, _ := range offer.Companies {
		offer.Companies[i].Idno = html.EscapeString(strings.TrimSpace(offer.Companies[i].Idno))
	}
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
		return errors.New("Required StartDate")
	}
	if offer.OfferData.EndDate == "" {
		return errors.New("Required EndDate")
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
		data, ok := company.CompaniesOffers["data"]
		if !ok {
			return errors.New("Required CompaniesOffers")
		}
		if data.Email == "" {
			return errors.New("Required CompaniesOffers Email")
		}
		if data.Sum == "" {
			return errors.New("Required CompaniesOffers Sum")
		}
		if data.Period == "" {
			return errors.New("Required CompaniesOffers Period")
		}
	}
	return nil
}
