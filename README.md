# logrus-runtime-formatter

[![CircleCI](https://circleci.com/gh/banzaicloud/logrus-runtime-formatter.svg?style=svg)](https://circleci.com/gh/banzaicloud/logrus-runtime-formatter)

Golang `runtime` package based automatic function, line and package fields for logrus. For further information and motivation behind the project please read this [post](https://banzaicloud.com/blog/runtime-logging/).

### tl;dr:

While we have been working on [Pipeline](https://github.com/banzaicloud/pipeline) we needed **function, package and line number information** to our log messages. We use Logrus and we could not find any similar extension so we have open sourced a Logrus runtime Formatter which **automatically tags log messages with runtime/stack information** without code modification.

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

```
go test
PASS
ok  	github.com/banzaicloud/logrus-runtime-formatter	0.005s
```

## Benchmark
`make bench`

```
go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/banzaicloud/logrus-runtime-formatter
BenchmarkErrorRuntimeAndTextFormatter-4   	  500000	      3191 ns/op	  26.00 MB/s	     822 B/op	      15 allocs/op
BenchmarkErrorTextFormatter-4             	 1000000	      1556 ns/op	  43.06 MB/s	     454 B/op	      12 allocs/op
BenchmarkSmallRuntimeAndTextFormatter-4   	  500000	      3110 ns/op	  32.79 MB/s	     848 B/op	      13 allocs/op
BenchmarkSmallTextFormatter-4             	 1000000	      1435 ns/op	  59.90 MB/s	     480 B/op	      10 allocs/op
BenchmarkLargeRuntimeAndTextFormatter-4   	  100000	     11617 ns/op	  26.00 MB/s	    6445 B/op	      19 allocs/op
BenchmarkLargeTextFormatter-4             	  300000	      4984 ns/op	  57.38 MB/s	    1728 B/op	      12 allocs/op
BenchmarkSmallRuntimeAndJSONFormatter-4   	  200000	      5884 ns/op	  22.77 MB/s	    2368 B/op	      34 allocs/op
BenchmarkSmallJSONFormatter-4             	  500000	      3732 ns/op	  30.54 MB/s	    1648 B/op	      28 allocs/op
BenchmarkLargeRuntimeAndJSONFormatter-4   	  100000	     22673 ns/op	  18.96 MB/s	   11629 B/op	      87 allocs/op
BenchmarkLargeJSONFormatter-4             	  100000	     16407 ns/op	  24.99 MB/s	    6906 B/op	      78 allocs/op
PASS
ok  	github.com/banzaicloud/logrus-runtime-formatter	17.501s
------------------------------------------------------------
```
