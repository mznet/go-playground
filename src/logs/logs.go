package logs

import (
	"log"
)

func Initialize() {

}

// Debug Level Logging
func Debug(v ...interface{}) {
	log.Println(v...)
}

// Info Level Logging
func Info(v ...interface{}) {
	log.Println(v...)
}

// Warn Level Logging
func Warn(v ...interface{}) {
	log.Println(v...)
}

// Error Level Logging
func Error(v ...interface{}) {
	log.Println(v...)
}

// Fatal Level Logging
func Fatal(v ...interface{}) {
	log.Println(v...)
}

