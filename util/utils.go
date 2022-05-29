package util

type Parameters struct {
	Parameters map[string]string
}

type Files struct {
	Files map[string]string
}

func NewEmptyFiles() *Files {
	return &Files{
		Files: make(map[string]string),
	}
}

func NewEmptyParameters() *Parameters {
	return &Parameters{
		Parameters: make(map[string]string),
	}
}
