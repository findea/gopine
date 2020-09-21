package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Level         string
	Format        string
	Output        string
	RotationCount uint
	RotationTime  string
}

const (
	FormatJson = "json"
	FormatText = "text"

	OutputStdout = "stdout"
	OutputStderr = "stderr"

	RotationTimeDay  = "day"
	RotationTimeHour = "hour"
)

func Init(config *Config) error {
	// level
	if config.Level != "" {
		level, err := logrus.ParseLevel(config.Level)
		if err != nil {
			return err
		}
		logrus.SetLevel(level)
	}

	// format
	switch config.Format {
	case FormatJson:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case FormatText:
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	}

	// output
	switch {
	case config.Output == OutputStdout:
		logrus.SetOutput(os.Stdout)
	case config.Output == OutputStderr:
		logrus.SetOutput(os.Stderr)
	case config.Output != "":
		hook, err := newLfsHook(config)
		if err != nil {
			return err
		}
		logrus.AddHook(hook)
		logrus.SetOutput(new(NopWriter))
	}

	return nil
}
