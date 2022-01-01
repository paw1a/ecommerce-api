package logging

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
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

var entry *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{entry}
}

func init() {
	l := logrus.New()
	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	entry = logrus.NewEntry(l)
}
