package main

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

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
	textFormatter := prefixed.TextFormatter{}
	textFormatter.FullTimestamp = true
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&ReflectFormatter{ChildFormatter: &textFormatter})

	bar()
	foo()
}
