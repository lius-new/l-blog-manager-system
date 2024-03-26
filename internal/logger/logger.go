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
