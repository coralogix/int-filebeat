package fshttp

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/publisher"
)

type logEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	App       string    `json:"app"`
	Pod       string    `json:"pod"`
}

func buildLogEntry(e publisher.Event) logEntry {
	msg, _ := e.Content.GetValue("message")
	app, _ := e.Content.GetValue("kubernetes.label.app")
	pod, _ := e.Content.GetValue("kubernetes.pod.name")

	return logEntry{
		Timestamp: e.Content.Timestamp,
		Message:   fmt.Sprintf("%v", msg),
		App:       fmt.Sprintf("%v", app),
		Pod:       fmt.Sprintf("%v", pod),
	}
}
