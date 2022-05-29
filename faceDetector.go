package GoDeepStack

import (
	"GoDeepStack/util"
	"fmt"
)

type FaceDetector struct {
	*AbstractDetector
	confidence float64
	ImagePath  string
}

type Face struct {
	XMax       int     `json:"x_max"`
	YMax       int     `json:"y_max"`
	XMin       int     `json:"x_min"`
	Confidence float64 `json:"confidence"`
	YMin       int     `json:"y_min"`
}

type Faces struct {
	Predictions []Face `json:"predictions"`
	Success     bool   `json:"success"`
}

func (fd *FaceDetector) Endpoint() string {
	c := fd.AbstractDetector.Client
	protocol := "http"
	if c.SSL {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s:%d/v1/vision/face", protocol, c.Addr, c.Port)
}

func (fd *FaceDetector) SetMinConfidence(confidence float64) *FaceDetector {
	fd.confidence = confidence
	return fd
}

func (fd *FaceDetector) Detect(imagePath string) *Faces {
	files := util.NewEmptyFiles()
	files.Files["image"] = imagePath
	var result *Faces
	fd.SendRequest(files, util.NewEmptyParameters(), &result)
	return result
}
