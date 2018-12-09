package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

const pattern = "[%-5s] %-30s [%20s:%4d] - %s"

func Init(logDirectory string, service string) {

	logfile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logDirectory, service), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("FATAL", "Open log file error:", err)
	}

	log.SetOutput(logfile)
	log.SetFlags(log.LstdFlags | log.LUTC)
}

func write(level string, message string) {

	pc, file, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	funcName = funcName[strings.LastIndex(funcName, "/")+1:]
	fileName := file[strings.LastIndex(file, "/")+1:]

	log.Printf(pattern, level, funcName, fileName, line, message)
}

func Fatal(message string) {
	write("FATAL", message)
}

func Error(message string) {
	write("ERROR", message)
}

func Warn(message string) {
	write("WARN", message)
}

func Info(message string) {
	write("INFO", message)
}

func Debug(message string) {
	write("DEBUG", message)
}

func Trace(message string) {
	write("TRACE", message)
}
