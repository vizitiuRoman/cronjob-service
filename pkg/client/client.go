package client

import (
	"os"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

var chatBotURL = os.Getenv("CB_PJ_URL")

func SendOfferToMBB(offers []byte) {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.SetRequestURI(chatBotURL)
	req.Header.SetContentType("application/json")
	req.SetBody(offers)
	req.Header.SetMethodBytes([]byte("POST"))

	err := fasthttp.Do(req, res)
	if err != nil {
		zap.S().Errorf("Error to send offer: %v", err)
	}
	zap.S().Info(string(res.Body()))
}
