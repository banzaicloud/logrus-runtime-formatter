package runtime

import (
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

// smallFields is a small size data set for benchmarking
var smallFields = logrus.Fields{
	"foo":   "bar",
	"baz":   "qux",
	"one":   "two",
	"three": "four",
}

// largeFields is a large size data set for benchmarking
var largeFields = logrus.Fields{
	"foo":       "bar",
	"baz":       "qux",
	"one":       "two",
	"three":     "four",
	"five":      "six",
	"seven":     "eight",
	"nine":      "ten",
	"eleven":    "twelve",
	"thirteen":  "fourteen",
	"fifteen":   "sixteen",
	"seventeen": "eighteen",
	"nineteen":  "twenty",
	"a":         "b",
	"c":         "d",
	"e":         "f",
	"g":         "h",
	"i":         "j",
	"k":         "l",
	"m":         "n",
	"o":         "p",
	"q":         "r",
	"s":         "t",
	"u":         "v",
	"w":         "x",
	"y":         "z",
	"this":      "will",
	"make":      "thirty",
	"entries":   "yeah",
}

var errorFields = logrus.Fields{
	"foo": fmt.Errorf("bar"),
	"baz": fmt.Errorf("qux"),
}

func BenchmarkErrorRuntimeAndTextFormatter(b *testing.B) {
	doBenchmark(b, &Formatter{ChildFormatter: &logrus.TextFormatter{DisableColors: true}}, errorFields)
}

func BenchmarkErrorTextFormatter(b *testing.B) {
	doBenchmark(b, &logrus.TextFormatter{DisableColors: true}, errorFields)
}

func BenchmarkSmallRuntimeAndTextFormatter(b *testing.B) {
	doBenchmark(b, &Formatter{ChildFormatter: &logrus.TextFormatter{DisableColors: true}}, smallFields)
}

func BenchmarkSmallTextFormatter(b *testing.B) {
	doBenchmark(b, &logrus.TextFormatter{DisableColors: true}, smallFields)
}

func BenchmarkLargeRuntimeAndTextFormatter(b *testing.B) {
	doBenchmark(b, &Formatter{ChildFormatter: &logrus.TextFormatter{DisableColors: true}}, largeFields)
}

func BenchmarkLargeTextFormatter(b *testing.B) {
	doBenchmark(b, &logrus.TextFormatter{DisableColors: true}, largeFields)
}

func BenchmarkSmallRuntimeAndJSONFormatter(b *testing.B) {
	doBenchmark(b, &Formatter{ChildFormatter: &logrus.JSONFormatter{}}, smallFields)
}

func BenchmarkSmallJSONFormatter(b *testing.B) {
	doBenchmark(b, &logrus.JSONFormatter{}, smallFields)
}

func BenchmarkLargeRuntimeAndJSONFormatter(b *testing.B) {
	doBenchmark(b, &Formatter{ChildFormatter: &logrus.JSONFormatter{}}, largeFields)
}

func BenchmarkLargeJSONFormatter(b *testing.B) {
	doBenchmark(b, &logrus.JSONFormatter{}, largeFields)
}

func doBenchmark(b *testing.B, formatter logrus.Formatter, fields logrus.Fields) {
	logger := logrus.New()

	entry := &logrus.Entry{
		Time:    time.Time{},
		Level:   logrus.InfoLevel,
		Message: "message",
		Data:    fields,
		Logger:  logger,
	}
	var d []byte
	var err error
	for i := 0; i < b.N; i++ {
		d, err = formatter.Format(entry)
		if err != nil {
			b.Fatal(err)
		}
		b.SetBytes(int64(len(d)))
	}
}
