package logs

import (
	"fmt"
	"time"
)

type LogLevel int

const (
	Info LogLevel = iota
	Debug
	Warning
	Error
)

type Logger struct {
	level      LogLevel
	timeFormat string
}

func NewLogger(level LogLevel, timeFormat string) *Logger {
	return &Logger{
		level:      level,
		timeFormat: timeFormat,
	}
}

func main() {
	getLooger()
}

func (l *Logger) Log(msg string) {
	// if l.level != Debug {
	// 	return
	// }
	switch l.level {
	case Debug,Info:
		fmt.Printf("[%s] %s", time.Now().Format(l.timeFormat), msg)
	case Warning:
		fmt.Printf("[Warning][%s] %s", time.Now().Format(l.timeFormat), msg)
	case Error:
		fmt.Printf("[Error][%s] %s", time.Now().Format(l.timeFormat), msg)
	default:
		fmt.Printf("%s", msg)
	}
}

func PrintLine(content string) {
	fmt.Println(content)
}

func PrintError(errorContent error) {
	fmt.Println(errorContent)
}

func PrintJson(data interface{}) {
	fmt.Println(fmt.Sprintf("data -> %+v", data))
}

func getLooger() {
	logger := NewLogger(Info, time.RFC3339)
	logger.Log("printing log...")
}
