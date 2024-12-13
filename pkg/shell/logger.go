package shell

import (
	"log"
)

type LoggerInterface interface {
	Debug(...interface{})
	Debugf(string, ...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})

	Trace(...interface{})
	Tracef(string, ...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
}

type Logger struct {
	//telnet.Logger
	LoggerInterface
}

func (l Logger) Debugf(str string, vargs ...interface{}) {
	log.Printf(str, vargs...)
}

func (l Logger) Errorf(str string, vargs ...interface{}) {
	log.Printf(str, vargs...)
}

func (l Logger) Tracef(str string, vargs ...interface{}) {
	log.Printf(str, vargs...)
}

func (l Logger) Warnf(str string, vargs ...interface{}) {
	log.Printf(str, vargs...)
}
