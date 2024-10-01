package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

// var Logger *logrus.Logger

func GetLogger() *logrus.Logger {
	Logger := logrus.New()
	Logger.SetReportCaller(true)
	Logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}

	Logger.SetOutput(io.Discard)

	Logger.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	Logger.SetLevel(logrus.DebugLevel)

	return Logger
}

// func Init() {
// 	Logger := logrus.New()
// 	Logger.SetReportCaller(true)
// 	Logger.Formatter = &logrus.TextFormatter{
// 		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
// 			filename := path.Base(frame.File)
// 			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
// 		},
// 		DisableColors: false,
// 		FullTimestamp: true,
// 	}

// 	err := os.MkdirAll("logs", 0644)
// 	if err != nil {
// 		panic(err)
// 	}

// 	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
// 	if err != nil {
// 		panic(err)
// 	}

// 	Logger.SetOutput(io.Discard)

// 	Logger.AddHook(&writerHook{
// 		Writer:    []io.Writer{allFile, os.Stdout},
// 		LogLevels: logrus.AllLevels,
// 	})

// 	Logger.SetLevel(logrus.DebugLevel)

// Logger = Logger
// }
