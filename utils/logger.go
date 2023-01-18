package utils

import (
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	Info  *color.Color
	Error *color.Color
}

func NewLogger() *Logger {
	return &Logger{
		Info:  color.New(color.FgGreen),
		Error: color.New(color.FgRed),
	}
}

func TimeStamp() string {
	return time.Now().Format("15:04:05")
}
