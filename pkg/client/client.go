package client

import (
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
)

func SendOfferToMBB(offers []byte) {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.SetRequestURI(os.Getenv("CB_PJ_URL"))
	req.Header.SetContentType("application/json")
	req.SetBody(offers)
	req.Header.SetMethodBytes([]byte("POST"))

	err := fasthttp.Do(req, res)
	if err != nil {
		fmt.Printf("Error SendOffer: %s", err)
	}
	fmt.Println(string(res.Body()))
}
