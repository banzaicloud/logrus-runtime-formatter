package runtime

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/sirupsen/logrus"
)

func foo() {
	logrus.Debugln("Hello world from foo function!")
}

func bar() {
	log := logrus.WithFields(logrus.Fields{"test": "field"})
	log.Infoln("Hello world from bar function!")
}

type A struct{}

func (A) valueFunc() {
	logrus.Infoln("Hello world from valueFunc function!")
}

func (*A) pointerFunc() {
	logrus.Infoln("Hello world from pointerFunc function!")
}

func TestRuntimeFormatter(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	childFormatter := logrus.JSONFormatter{}
	formatter := &Formatter{ChildFormatter: &childFormatter}
	formatter.Line = true
	formatter.Package = true
	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(buffer)

	decoder := json.NewDecoder(buffer)

	foo()

	expectFunction(t, decoder, "github.com/banzaicloud/logrus-runtime-formatter", "foo", "12")

	bar()

	expectFunction(t, decoder, "github.com/banzaicloud/logrus-runtime-formatter", "bar", "17")

	A{}.valueFunc()

	expectFunction(t, decoder, "github.com/banzaicloud/logrus-runtime-formatter.A", "valueFunc", "23")

	(&A{}).pointerFunc()

	expectFunction(t, decoder, "github.com/banzaicloud/logrus-runtime-formatter.(*A)", "pointerFunc", "27")
}

func expectFunction(t *testing.T, decoder *json.Decoder, expectedPackage string, expectedFunction string, expectedLine string) {
	data := map[string]string{}
	err := decoder.Decode(&data)
	if err != nil {
		t.Fatal(err)
	}

	_package := data[PackageKey]
	function := data[FunctionKey]
	line := data[LineKey]

	if _package != expectedPackage {
		t.Fatalf("Expected package: %s, got: %s", expectedPackage, _package)
	}
	if function != expectedFunction {
		t.Fatalf("Expected function: %s, got: %s", expectedFunction, function)
	}
	if line != expectedLine {
		t.Fatalf("Expected line: %s, got: %s", expectedLine, line)
	}
}
