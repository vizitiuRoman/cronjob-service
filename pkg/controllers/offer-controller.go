package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	. "github.com/cronjobs-service/pkg/models"
	. "github.com/cronjobs-service/pkg/offer-job"
	. "github.com/cronjobs-service/pkg/utils"
	"github.com/gorilla/mux"
)

func StartOfferCron(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var offer Offer
	err = json.Unmarshal(body, &offer)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = offer.Validate()
	if err != nil {
		ERROR(w, http.StatusBadRequest, err)
		return
	}

	go RunJob(int(offer.ID), offer.Name, offer.StartDate, offer.EndDate)

	JSON(w, http.StatusOK, "")
}

func StopOffer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	offerID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, errors.New(http.StatusText(http.StatusUnprocessableEntity)))
		return
	}

	StopJob(int(offerID))

	JSON(w, http.StatusOK, "")
}