package structural

import (
	"fmt"
	"strings"
)

type ILogger interface {
	LogMessage(message string)
}

type Logger struct {
}

func (logger *Logger) LogMessage(message string) {
	fmt.Println(message)
}

type UpperCaseLogger struct {
	Logger
}

func (logger *UpperCaseLogger) LogMessage(message string) {
	logger.Logger.LogMessage(strings.ToUpper(message))
}
