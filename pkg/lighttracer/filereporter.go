package lighttracer

import (
	"encoding/json"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/openzipkin/zipkin-go/reporter"
	"os"
)

type fileReporter struct {
	*os.File
}

func NewFileReporter(filename string) (reporter.Reporter, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &fileReporter{file}, nil
}

func (r *fileReporter) Send(s model.SpanModel) {
	if b, err := json.Marshal(s); err == nil {
		r.WriteString(string(b))
		r.WriteString("\n")
	}
}

func (r *fileReporter) Close() error {
	return r.Close()
}
