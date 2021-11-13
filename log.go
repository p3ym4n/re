package re

import "log"

type LoggerFunc func(Error)

var logger LoggerFunc = func(bag Error) {
	kind := bag.Kind()
	if kind == KindInvalid || kind == KindNotFound {
		return
	}
	log.Println(bag.AsMap())
}

func SetLogger(loggerFunc LoggerFunc) {
	logger = loggerFunc
}

func Log(err Error) {
	logger(err)
}
