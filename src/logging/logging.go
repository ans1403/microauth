package logging

import (
	"log"
	"os"
)

type Logger interface {
	Info(v ...any)
	Warning(v ...any)
	Error(v ...any)
}

func NewLogger() Logger {
	return &logger{}
}

type logger struct{}

func (l *logger) Info(v ...any) {
	logger := log.New(os.Stdout, "[INFO] ", log.Ltime)
	logger.Println(v...)
}

func (l *logger) Warning(v ...any) {
	logger := log.New(os.Stdout, "[WARNING] ", log.Ltime)
	logger.Println(v...)
}

func (l *logger) Error(v ...any) {
	logger := log.New(os.Stdout, "[ERROR] ", log.Ldate)
	logger.Println(v...)
}
