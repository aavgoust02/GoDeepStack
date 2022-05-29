package GoDeepStack

import (
	"fmt"
	"github.com/aavgoust02/GoDeepStack/util"
)

type ImageEnhancer struct {
	*AbstractDetector
	client    *DeepStack
	ImagePath string
}

type Image struct {
	Success bool   `json:"success"`
	Base64  string `json:"base64"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

func (ie *ImageEnhancer) Endpoint() string {
	c := ie.AbstractDetector.Client
	protocol := "http"
	if c.SSL {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s:%d/v1/vision/enhance", protocol, c.Addr, c.Port)
}

func (ie *ImageEnhancer) Enhance(imagePath string) *Image {
	files := util.NewEmptyFiles()
	files.Files["image"] = imagePath
	var result *Image
	ie.SendRequest(files, util.NewEmptyParameters(), &result)
	return result
}
