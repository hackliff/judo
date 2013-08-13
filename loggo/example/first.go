package main

import (
	"launchpad.net/loggo"
)

var first = loggo.GetLogger("first")

func FirstCritical(message string) {
	first.Criticalf(message)
}

func FirstError(message string) {
	first.Errorf(message)
}

func FirstWarning(message string) {
	first.Warningf(message)
}

func FirstInfo(message string) {
	first.Infof(message)
}

func FirstTrace(message string) {
	first.Tracef(message)
}