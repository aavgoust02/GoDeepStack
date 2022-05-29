package GoDeepStack

import (
	"fmt"
	"github.com/aavgoust02/GoDeepStack/util"
)

type FaceMatcher struct {
	*AbstractDetector
	client           *DeepStack
	ImagePath        string
	AnotherImagePath string
}

type FaceMatch struct {
	Similarity float64 `json:"similarity"`
	Success    bool    `json:"success"`
}

func (fm *FaceMatcher) Endpoint() string {
	c := fm.AbstractDetector.Client
	protocol := "http"
	if c.SSL {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s:%d/v1/vision/face/match", protocol, c.Addr, c.Port)
}

func (fm *FaceMatcher) Match(image1 string, image2 string) *FaceMatch {
	files := util.NewEmptyFiles()
	files.Files["image1"] = image1
	files.Files["image2"] = image2
	var result *FaceMatch
	fm.SendRequest(files, util.NewEmptyParameters(), &result)
	return result
}
