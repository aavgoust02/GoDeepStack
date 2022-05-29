package GoDeepStack

import (
	"github.com/aavgoust02/GoDeepStack/util"
	"github.com/imroc/req/v3"
	"log"
	"net/url"
)

type Detector interface {
	Endpoint() string
}

type AbstractDetector struct {
	Detector
	Client *DeepStack
}

func (a *AbstractDetector) SendRequest(files *util.Files, additionalParameters *util.Parameters, out interface{}) interface{} {
	client := req.C().R()
	if a.Client.ApiKey != "" {
		var apiKey url.Values
		apiKey.Add("api_key", a.Client.ApiKey)
		client.SetFormDataFromValues(apiKey)
	}
	request := client.SetFiles(files.Files).SetFormData(additionalParameters.Parameters).SetResult(&out)
	response, err := request.Post(a.Endpoint())
	util.Check(err)
	if !response.IsSuccess() {
		log.Fatalf("Unable to run query [%d]", response.StatusCode)
	}
	return out
}
