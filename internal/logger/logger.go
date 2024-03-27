package logger

import (
	"fmt"
	"log"
)

const (
	LevelDebug int = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func Distribute(level int, content string) {
	switch level {
	case LevelDebug:
		log.Println("debug", content)
	case LevelInfo:
		log.Println("info", content)
	case LevelWarn:
		log.Println("warn", content)
	case LevelFatal:
		log.Println("fatal", content)
	case LevelPanic:
		log.Println("panic", content)
	}
}

func Debug(content ...interface{}) {
	Distribute(LevelDebug, fmt.Sprintln(content))
}

func Info(content interface{}) {
	Distribute(LevelInfo, fmt.Sprintln(content))
}

func Warn(content interface{}) {
	Distribute(LevelWarn, fmt.Sprintln(content))
}
func Error(content string) {
	Distribute(LevelError, content)
}
func Fatal(content string) {
	Distribute(LevelFatal, content)
}
func Panic(content ...string) {
	Distribute(LevelPanic, fmt.Sprintln(content))
}
