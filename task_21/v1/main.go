package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	_ "github.com/rs/zerolog/log"
)

type Logger interface {
	Info(message string)
	Error(message string)
}

type CustomLogger struct{}

func (c CustomLogger) Info(message string) {
	fmt.Println("INFO", message)
}
func (c CustomLogger) Error(message string) {
	fmt.Println("Error", message)
}

type ZeroLoggerAdapter struct{}

func (zl *ZeroLoggerAdapter) Info(message string) {
	log.Info().Msg(message)
}
func (zl *ZeroLoggerAdapter) Error(message string) {
	log.Error().Msg(message)
}

func main() {
	loggers := []Logger{&ZeroLoggerAdapter{}, &CustomLogger{}}

	for _, logger := range loggers {
		logger.Info("This is INFO message")
		logger.Error("This is ERROR message")
	}
}
