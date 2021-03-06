package util

import (
	"fmt"
	"log"
	"os"
)

// This is meant to be a better logging package that uses the default logger,
// but doesn't just spam the crap out of everything.

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var logLevel LogLevel = INFO

func LoggingLevel() LogLevel {
	return logLevel
}

func DisableLogging() {
	logLevel = FATAL + 1
}

func SetOutput(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err.Error())
	}
	log.SetOutput(f)
}

func SetLoggingLevel(level LogLevel) {
	if level < 0 || level > FATAL {
		panic("Not a valid log level")
	}
	logLevel = level

	if logLevel <= DEBUG {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(log.LstdFlags)
	}
}

func Info(msg string) {
	if logLevel <= INFO {
		log.SetPrefix("[INFO] ")
		log.Output(2, msg)
	}
}

func Infof(msg string, args ...interface{}) {
	if logLevel <= INFO {
		log.SetPrefix("[INFO] ")
		log.Output(2, fmt.Sprintf(msg, args...))
	}
}

func Debug(msg string) {
	if logLevel <= DEBUG {
		log.SetPrefix("[DEBUG] ")
		log.Output(2, msg)
	}
}

func Debugf(msg string, args ...interface{}) {
	if logLevel <= DEBUG {
		log.SetPrefix("[DEBUG] ")
		log.Output(2, fmt.Sprintf(msg, args...))
	}
}

func Warn(msg string) {
	if logLevel <= WARN {
		log.SetPrefix("[WARN] ")
		log.Output(2, msg)
	}
}

func Warnf(msg string, args ...interface{}) {
	if logLevel <= WARN {
		log.SetPrefix("[WARN] ")
		log.Output(2, fmt.Sprintf(msg, args...))
	}
}

func Error(msg string) {
	if logLevel <= ERROR {
		log.SetPrefix("[ERROR] ")
		log.Print(msg)
	}
}

func Errorf(msg string, args ...interface{}) {
	if logLevel <= ERROR {
		log.SetPrefix("[ERROR] ")
		log.Printf(msg, args...)
	}
}

func Fatal(msg string) {
	if logLevel <= FATAL {
		log.SetPrefix("[FATAL] ")
		log.Print(msg)
	}
}

func Fatalf(msg string, args ...interface{}) {
	if logLevel <= FATAL {
		log.SetPrefix("[FATAL] ")
		log.Printf(msg, args...)
	}
}
