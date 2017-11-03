package logs

import (
	"log"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"runtime"
	"strconv"
	"fmt"
)


func Initialize() {
	log.SetOutput(&lumberjack.Logger{
		Filename:	"./tmp/production.log",
		MaxSize:	10,
		MaxBackups:	5,
		MaxAge:		10,
	})
}

// Debug Level Logging
func Debug(v ...interface{}) {
	shortFile := "???:-1"
	if _, file, line, ok := runtime.Caller(1); ok {
		shortFile = getShortFile(file, line)
	}
	log.Println(shortFile + " [Debug] " + fmt.Sprint(v...))
}

// Info Level Logging
func Info(v ...interface{}) {
	shortFile := "???:-1"
	if _, file, line, ok := runtime.Caller(1); ok {
		shortFile = getShortFile(file, line)
	}
	log.Println(shortFile + " [INFO] " + fmt.Sprint(v...))
}

// Warn Level Logging
func Warn(v ...interface{}) {
	shortFile := "???:-1"
	if _, file, line, ok := runtime.Caller(1); ok {
		shortFile = getShortFile(file, line)
	}
	log.Println(shortFile + " [WARN] " + fmt.Sprint(v...))
}

// Error Level Logging
func Error(v ...interface{}) {
	shortFile := "???:-1"
	if _, file, line, ok := runtime.Caller(1); ok {
		shortFile = getShortFile(file, line)
	}
	log.Println(shortFile + " [ERROR] " + fmt.Sprint(v...))
}

// Fatal Level Logging
func Fatal(v ...interface{}) {
	shortFile := "???:-1"
	if _, file, line, ok := runtime.Caller(1); ok {
		shortFile = getShortFile(file, line)
	}
	log.Println(shortFile + " [FATAL] " + fmt.Sprint(v...))
}

func getShortFile(file string, line int) (path string) {
	short := file
	for i := len(file) -1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	return short + ":" + strconv.Itoa(line)
}