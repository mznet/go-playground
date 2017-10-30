package logs

import (
	"log"
)

// Info Level Logging
func Info(v ...interface{}) {
	log.Println(v...)
}


// Error Level Logging
func Error(v ...interface{}) {
	log.Println(v...)
}
