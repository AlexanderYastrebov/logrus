package logrus_test

import (
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func ExampleJSONFormatter_CallerPrettyfier() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Out = os.Stdout
	l.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			const (
				maximumCallerDepth = 25
				minimumCallerDepth = 4
			)
			pcs := make([]uintptr, maximumCallerDepth)
			depth := runtime.Callers(minimumCallerDepth, pcs)
			frames := runtime.CallersFrames(pcs[:depth])

			for frame, again := frames.Next(); again; frame, again = frames.Next() {
				if frame.PC == f.PC {
					// get parent of f frame
					frame, _ := frames.Next()
					f = &frame
					break
				}
			}

			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	}

	myfunc := func() {
		l.Info("example of custom format caller")
	}
	myfunc()

	// Output:
	// {"file":"example_custom_caller_test.go","func":"ExampleJSONFormatter_CallerPrettyfier","level":"info","msg":"example of custom format caller"}
}
