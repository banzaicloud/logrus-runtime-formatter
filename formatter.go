package main

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

const FunctionKey = "function"
const LineKey = "line"

type ReflectFormatter struct {
	ChildFormatter logrus.Formatter
	LineNumber     bool
}

// Format the current log entry by adding the function name and line number of the caller.
func (f *ReflectFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	function, line := f.getCurrentPosition(entry)

	fields := logrus.Fields{FunctionKey: function}
	if f.LineNumber {
		fields[LineKey] = line
	}
	newEntry := entry.WithFields(fields)
	entry.Data = newEntry.Data

	// entry.Data[FunctionKey] = function
	// if f.LineNumber {
	// 	entry.Data[LineKey] = line
	// }

	return f.ChildFormatter.Format(entry)
}

func (f *ReflectFormatter) getCurrentPosition(entry *logrus.Entry) (string, string) {
	skip := 6
	if len(entry.Data) == 0 {
		skip = 8
	}
	function, _, line, _ := runtime.Caller(skip)
	lineNumber := ""
	if f.LineNumber {
		lineNumber = fmt.Sprintf("%d", line)
	}
	return runtime.FuncForPC(function).Name(), lineNumber
}
