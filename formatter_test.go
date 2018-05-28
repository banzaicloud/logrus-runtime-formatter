package reflected

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func bar() {
	logrus.WithFields(logrus.Fields{"test": "field"}).Infoln("Hello world from bar function!")
}

func foo() {
	logrus.Debugln("Hello world from foo function!")
}

func TestReflected(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	childFormatter := logrus.JSONFormatter{}
	reflectedFormatter := &ReflectedFormatter{ChildFormatter: &childFormatter}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(reflectedFormatter)
	logrus.SetOutput(buffer)

	decoder := json.NewDecoder(buffer)

	foo()

	expectFunction(t, decoder, "foo")

	bar()

	expectFunction(t, decoder, "bar")
}

func expectFunction(t *testing.T, decoder *json.Decoder, expectedFunction string) {
	data := map[string]string{}
	err := decoder.Decode(&data)
	if err != nil {
		t.Fatal(err)
	}

	packageAndFunc := strings.Split(data[FunctionKey], ".")
	function := packageAndFunc[len(packageAndFunc)-1]

	if function != expectedFunction {
		t.Fatalf("Expected function: %s, got: %s", expectedFunction, function)
	}
}
