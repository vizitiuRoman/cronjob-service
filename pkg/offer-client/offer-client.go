package offer_client

import (
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
)

var ChatBotURL = os.Getenv("CB_PJ_URL")

func SendOffer(offers []byte) {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.SetRequestURI(ChatBotURL)
	req.Header.SetContentType("application/json")
	req.SetBody(offers)
	req.Header.SetMethodBytes([]byte("POST"))

	err := fasthttp.Do(req, res)
	if err != nil {
		fmt.Printf("Error SendOffer: %s", err)
	}

	bodyBytes := res.Body()
	fmt.Println(string(bodyBytes))
}
