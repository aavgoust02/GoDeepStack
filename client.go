package GoDeepStack

type DeepStack struct {
	Addr   string
	Port   int
	ApiKey string
	SSL    bool
}

func NewDeepStackClient(addr string, port int, ssl bool, apiKey string) *DeepStack {
	return &DeepStack{
		Addr:   addr,
		Port:   port,
		SSL:    ssl,
		ApiKey: apiKey,
	}
}

func (d *DeepStack) NewObjectDetector() *ObjectDetector {
	od := &ObjectDetector{
		AbstractDetector: &AbstractDetector{
			Client: d,
		},
		confidence: 0.5,
	}
	od.AbstractDetector.Detector = od
	return od
}

func (d *DeepStack) NewFaceDetector() *FaceDetector {
	fd := &FaceDetector{
		AbstractDetector: &AbstractDetector{
			Client: d,
		},
		confidence: 0.0,
	}
	fd.AbstractDetector.Detector = fd
	return fd
}

func (d *DeepStack) NewFaceMatcher() *FaceMatcher {
	fm := &FaceMatcher{
		AbstractDetector: &AbstractDetector{
			Client: d,
		},
	}
	fm.AbstractDetector.Detector = fm
	return fm
}

func (d *DeepStack) NewImageEnhancer() *ImageEnhancer {
	ie := &ImageEnhancer{
		AbstractDetector: &AbstractDetector{
			Client: d,
		},
	}
	ie.AbstractDetector.Detector = ie
	return ie
}
