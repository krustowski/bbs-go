package main

import (
	"github.com/reiver/go-telnet"

	"log"
)

type Logger interface{
	Debug(...interface{})
	Debugf(string, ...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})

	Trace(...interface{})
	Tracef(string, ...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
}

type logger struct{
	//telnet.Logger
	Logger
}

func (l logger) Debugf(str string, vargs ...interface{}) {
	log.Printf(str, vargs...)
}

func (l logger) Errorf(str string, vargs ...interface{}) {
	log.Printf(str, vargs...)
}

func (l logger) Tracef(str string, vargs ...interface{}) {
	log.Printf(str, vargs...)
}

func (l logger) Warnf(str string, vargs ...interface{}) {
	log.Printf(str, vargs...)
}
