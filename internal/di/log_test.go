package di

import (
	"goweb/pkg/log"
	"testing"
)

func TestLog(t *testing.T) {
	logExamples()

	log.Init(&log.Config{
		Format: "json",
		Output: "stdout",
		Level:  "warn",
	})

	logExamples()
}

func logExamples() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Error("The ice breaks!")
}
