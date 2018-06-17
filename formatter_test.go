package runtime

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func foo() {
	logrus.Debug("Hello world from foo function!")
}

func bar() {
	log := logrus.WithFields(logrus.Fields{"test": "field"})
	log.Infoln("Hello world from bar function!")
	log.Infof("Hello world from bar function!")
}

type A struct{}

func (A) valueFunc() {
	logrus.Print("Hello world from valueFunc function!")
}

func (*A) pointerFunc() {
	logrus.Printf("Hello world from pointerFunc function!")
}

func (*A) ReflectedFunc(msg string) {
	logrus.Printf("Hello world from ReflectedFunc function: %s", msg)
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

	expectFunction(t, decoder, "github.com/banzaicloud/logrus-runtime-formatter", "foo", "13")

	bar()

	expectFunction(t, decoder, "github.com/banzaicloud/logrus-runtime-formatter", "bar", "18")
	expectFunction(t, decoder, "github.com/banzaicloud/logrus-runtime-formatter", "bar", "19")

	a := A{}

	a.valueFunc()

	expectFunction(t, decoder, "github.com/banzaicloud/logrus-runtime-formatter.A", "valueFunc", "25")

	(&a).pointerFunc()

	expectFunction(t, decoder, "github.com/banzaicloud/logrus-runtime-formatter.(*A)", "pointerFunc", "29")

	switch method := reflect.ValueOf(&a).MethodByName("ReflectedFunc").Interface().(type) {
	case func(string):
		method("hello world")
	}

	expectFunction(t, decoder, "github.com/banzaicloud/logrus-runtime-formatter.(*A)", "ReflectedFunc", "33")
}

func TestFunctionInFunctionFormatter(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	childFormatter := logrus.JSONFormatter{}
	formatter := &Formatter{ChildFormatter: &childFormatter}
	formatter.Line = true
	formatter.Package = true
	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(buffer)

	decoder := json.NewDecoder(buffer)

	funcInFunc()

	expectFunction(t, decoder, "github.com/banzaicloud/logrus-runtime-formatter", "baz", "118")

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

func baz() {
	logrus.Debug("Hello world from baz function!")
}

func funcInFunc() {
	baz()
}
