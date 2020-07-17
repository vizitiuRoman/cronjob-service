package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	. "github.com/cronjob-service/pkg/models"
	. "github.com/cronjob-service/pkg/offer-job"
	. "github.com/cronjob-service/pkg/utils"
	"github.com/gorilla/mux"
)

func StartOfferJob(w http.ResponseWriter, r *http.Request) {
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

	go StartJob(int(offer.ID), offer.Name, offer.StartDate, offer.EndDate)

	JSON(w, http.StatusCreated, offer.ID)
}

func DeleteOfferJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	offerID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, errors.New(http.StatusText(http.StatusUnprocessableEntity)))
		return
	}

	err = DeleteJobByID(int(offerID))
	if err != nil {
		ERROR(w, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
		return
	}

	JSON(w, http.StatusOK, offerID)
}

func GetJobs(w http.ResponseWriter, r *http.Request) {
	runningJobs := GetRunningJobs()
	JSON(w, http.StatusOK, runningJobs)
}
