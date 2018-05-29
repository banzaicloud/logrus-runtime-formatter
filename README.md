# logrus-runtime-formatter
Golang `runtime` package based automatic function, line and package fields for logrus.

## Usage

You have to wrap your desired Formatter with the runtime.Formatter and it will do the job:

```go
package main

import (
	"github.com/sirupsen/logrus"
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
)

var log = logrus.New()

func init() {
	childFormatter := logrus.JSONFormatter{}
	runtimeFormatter := &runtime.Formatter{ChildFormatter: &childFormatter}
	log.Formatter = runtimeFormatter
	log.Level = logrus.DebugLevel
}

func main() {
	log.WithFields(logrus.Fields{
		"prefix": "main",
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	log.WithFields(logrus.Fields{
		"prefix":      "sensor",
		"temperature": -4,
	}).Info("Temperature changes")
}
```

## Test
`make test`

## Benchmark
`make bench`

```
go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/banzaicloud/logrus-reflect-formatter
BenchmarkErrorRuntimeAndTextFormatter-4   	  500000	      3092 ns/op	  29.42 MB/s	     822 B/op	      15 allocs/op
BenchmarkErrorTextFormatter-4               	 1000000	      1582 ns/op	  42.34 MB/s	     454 B/op	      12 allocs/op
BenchmarkSmallRuntimeAndTextFormatter-4   	  500000	      2971 ns/op	  37.02 MB/s	     848 B/op	      13 allocs/op
BenchmarkSmallTextFormatter-4               	 1000000	      1401 ns/op	  61.38 MB/s	     480 B/op	      10 allocs/op
BenchmarkLargeRuntimeAndTextFormatter-4   	  100000	     11263 ns/op	  27.52 MB/s	    6446 B/op	      19 allocs/op
BenchmarkLargeTextFormatter-4               	  300000	      4849 ns/op	  58.97 MB/s	    1728 B/op	      12 allocs/op
BenchmarkSmallRuntimeAndJSONFormatter-4   	  300000	      5663 ns/op	  25.07 MB/s	    2368 B/op	      34 allocs/op
BenchmarkSmallJSONFormatter-4               	  500000	      3698 ns/op	  30.82 MB/s	    1648 B/op	      28 allocs/op
BenchmarkLargeRuntimeAndJSONFormatter-4   	  100000	     22549 ns/op	  19.42 MB/s	   11629 B/op	      87 allocs/op
BenchmarkLargeJSONFormatter-4               	  100000	     15895 ns/op	  25.79 MB/s	    6906 B/op	      78 allocs/op
```