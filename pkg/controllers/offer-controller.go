package controllers

import (
	"encoding/json"
	"strconv"

	. "github.com/cronjob-service/pkg/models"
	. "github.com/cronjob-service/pkg/offer-job"
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
		offer.Prepare()
		err = offer.Validate()
		if err != nil {
			ERROR(ctx, fasthttp.StatusBadRequest, err)
			return
		}

		offerJob := NewOfferJob(int(offer.ID),
			offer.Title,
			offer.OfferData.RepeatNumb,
			offer.OfferData.RepeatTime,
		)
		go StartJob(offerJob)
	}

	JSON(ctx, fasthttp.StatusOK, offers)
}

func GetJobs(ctx *fasthttp.RequestCtx) {
	runningJobs := GetRunningJobs()
	JSON(ctx, fasthttp.StatusOK, runningJobs)
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
