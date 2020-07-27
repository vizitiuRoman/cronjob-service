package controllers

import (
	"bytes"
	"encoding/json"
	"strconv"

	. "github.com/cronjob-service/pkg/models"
	. "github.com/cronjob-service/pkg/utils"
	"github.com/valyala/fasthttp"
)

func StartOfferJob(ctx *fasthttp.RequestCtx) {
	var offers []Offer
	err := json.Unmarshal(ctx.PostBody(), &offers)
	if err != nil {
		ERROR(ctx, fasthttp.StatusUnprocessableEntity, err)
		return
	}

	for _, offer := range offers {
		offersBytes := new(bytes.Buffer)
		err := json.NewEncoder(offersBytes).Encode(offers)
		if err != nil {
			ERROR(ctx, fasthttp.StatusUnprocessableEntity, err)
			return
		}

		offer.Prepare()
		err = offer.Validate()
		if err != nil {
			ERROR(ctx, fasthttp.StatusBadRequest, err)
			return
		}

		offerJob := NewOfferJob(
			int(offer.ID),
			offer.OfferData.RepeatNumb,
			offer.OfferData.RepeatTime,
			offersBytes.Bytes(),
		)
		go offerJob.StartJob()
	}

	JSON(ctx, fasthttp.StatusOK, true)
}

func UpdateOfferJob(ctx *fasthttp.RequestCtx) {
	var offers []Offer
	err := json.Unmarshal(ctx.PostBody(), &offers)
	if err != nil {
		ERROR(ctx, fasthttp.StatusUnprocessableEntity, err)
		return
	}

	for _, offer := range offers {
		offersBytes := new(bytes.Buffer)
		err := json.NewEncoder(offersBytes).Encode(offers)
		if err != nil {
			ERROR(ctx, fasthttp.StatusUnprocessableEntity, err)
			return
		}

		offer.Prepare()
		err = offer.Validate()
		if err != nil {
			ERROR(ctx, fasthttp.StatusBadRequest, err)
			return
		}

		err = DeleteJobByID(int(offer.ID))
		if err != nil {
			ERROR(ctx, fasthttp.StatusNotFound, err)
			return
		}

		offerJob := NewOfferJob(
			int(offer.ID),
			offer.OfferData.RepeatNumb,
			offer.OfferData.RepeatTime,
			offersBytes.Bytes(),
		)
		go offerJob.StartJob()
	}

	JSON(ctx, fasthttp.StatusOK, true)
}

func DeleteOfferJob(ctx *fasthttp.RequestCtx) {
	offerID, err := strconv.ParseInt(ctx.UserValue("id").(string), 10, 64)
	if err != nil {
		ERROR(ctx, fasthttp.StatusUnprocessableEntity, err)
		return
	}

	err = DeleteJobByID(int(offerID))
	if err != nil {
		ERROR(ctx, fasthttp.StatusNotFound, err)
		return
	}
	JSON(ctx, fasthttp.StatusOK, offerID)
}

func GetJobs(ctx *fasthttp.RequestCtx) {
	entries := GetRunningJobs()
	JSON(ctx, fasthttp.StatusOK, entries)
}
