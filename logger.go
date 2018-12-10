package main

import (
	logrus "github.com/sirupsen/logrus"
)

// Info - Print a Info
func Info(message ...interface{}) {
	logrus.Println(message)
}

// Panic - Print a Panic
func Panic(err error) {
	logrus.Panicln(err)
}
