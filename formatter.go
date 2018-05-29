package runtime

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// FunctionKey holds the function field
const FunctionKey = "function"

// PackageKey holds the package field
const PackageKey = "package"

// LineKey holds the line field
const LineKey = "line"

// Formatter decorates log entries with function name and package name (optional) and line number (optional)
type Formatter struct {
	ChildFormatter logrus.Formatter
	// When true, line number will be tagged to fields as well
	Line bool
	// When true, package name will be tagged to fields as well
	Package bool
}

// Format the current log entry by adding the function name and line number of the caller.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	function, line := f.getCurrentPosition(entry)

	packageEnd := strings.LastIndex(function, ".")
	functionName := function[packageEnd+1:]

	data := logrus.Fields{FunctionKey: functionName}
	if f.Line {
		data[LineKey] = line
	}
	if f.Package {
		data[PackageKey] = function[:packageEnd]
	}
	for k, v := range entry.Data {
		data[k] = v
	}
	entry.Data = data

	// entry.Data[FunctionKey] = function
	// if f.LineNumber {
	// 	entry.Data[LineKey] = line
	// }

	return f.ChildFormatter.Format(entry)
}

func (f *Formatter) getCurrentPosition(entry *logrus.Entry) (string, string) {
	skip := 6
	if len(entry.Data) == 0 {
		skip = 8
	}
	function, _, line, _ := runtime.Caller(skip)
	lineNumber := ""
	if f.Line {
		lineNumber = fmt.Sprintf("%d", line)
	}
	return runtime.FuncForPC(function).Name(), lineNumber
}
