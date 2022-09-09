//go:build !windows
// +build !windows

package logrus_test

import (
	"os"

	"github.com/sirupsen/logrus"
)

type getCallerHook struct{}

func (h *getCallerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *getCallerHook) Fire(e *logrus.Entry) error {
	// get caller here by walking frames
	e.Data["caller"] = "acaller"
	return nil
}

// An example on how to use a hook
func Example_hook() {
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)                     // default
	log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	log.Hooks.Add(&getCallerHook{})
	log.Out = os.Stdout

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Error("The ice breaks!")

	// Output:
	// level=info msg="A group of walrus emerges from the ocean" animal=walrus size=10
	// level=warning msg="The group's number increased tremendously!" number=122 omg=true
	// level=error msg="The ice breaks!" number=100 omg=true
}
