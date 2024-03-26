package logger

import "log"

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

func Debug(content string) {
	Distribute(LevelDebug, content)
}

func Info(content string) {
	Distribute(LevelInfo, content)
}

func Warn(content string) {
	Distribute(LevelWarn, content)
}
func Error(content string) {
	Distribute(LevelError, content)
}
func Fatal(content string) {
	Distribute(LevelFatal, content)
}
func Panic(content string) {
	Distribute(LevelPanic, content)
}
