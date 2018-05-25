package main

import (
	"runtime"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

//////

type ReflectFormatter struct {
	ChildFormatter logrus.Formatter
}

func (f *ReflectFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	frame := getCurrentFrame(entry)
	fields := logrus.Fields{"function": frame.Function, "line": frame.Line}
	newEntry := entry.WithFields(fields)
	entry.Data = newEntry.Data
	return f.ChildFormatter.Format(entry)
}

func getCurrentFrame(entry *logrus.Entry) *runtime.Frame {
	skip := 6
	if len(entry.Data) == 0 {
		skip = 8
	}
	pc, _, _, _ := runtime.Caller(skip)
	frames := runtime.CallersFrames([]uintptr{pc})
	frame, _ := frames.Next()
	return &frame
}

//////

func init() {
	textFormatter := prefixed.TextFormatter{}
	textFormatter.FullTimestamp = true
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&ReflectFormatter{ChildFormatter: &textFormatter})
}

func bar() {
	foo()
	logrus.WithFields(logrus.Fields{"test": "field"}).WithFields(logrus.Fields{"ez": "mas"}).Errorln("Hello world from bar function!")
	println()
	logrus.WithFields(logrus.Fields{"test": "field"}).Warnln("Hello world from bar function!")
	println()
}

func foo() {
	logrus.Infoln("Hello world from foo function!")
	println()
	logrus.Debugln("Hello world from foo debugged function!")
	println()
}

func main() {
	bar()
	foo()
}
