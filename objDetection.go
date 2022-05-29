package GoDeepStack

import (
	"GoDeepStack/util"
	"fmt"
)

type ObjectDetector struct {
	*AbstractDetector
	client     *DeepStack
	confidence float64
	ImagePath  string
}

type Object struct {
	XMax       float64 `json:"x_max"`
	XMin       float64 `json:"x_min"`
	YMin       float64 `json:"y_min"`
	Confidence float64 `json:"confidence"`
	Label      string  `json:"label"`
	YMax       float64 `json:"y_max"`
}

type Objects struct {
	Predictions []Object `json:"predictions"`
	Success     bool     `json:"success"`
}

func (od *ObjectDetector) Endpoint() string {
	c := od.AbstractDetector.Client
	protocol := "http"
	if c.SSL {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s:%d/v1/vision/detection", protocol, c.Addr, c.Port)
}

func (od *ObjectDetector) SetMinConfidence(confidence float64) *ObjectDetector {
	od.confidence = confidence
	return od
}

func (od *ObjectDetector) Detect(imagePath string) *Objects {
	files := util.NewEmptyFiles()
	files.Files["image"] = imagePath
	params := util.NewEmptyParameters()
	params.Parameters["min_confidence"] = fmt.Sprintf("%f", od.confidence)
	var result *Objects
	od.SendRequest(files, params, &result)
	return result
}
