package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

const pattern = "[%-5s] %-30s [%20s:%4d] - %s\n"

func Init(logDirectory string, service string) {

	logfile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logDirectory, service), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("FATAL", "Open log file error:", err)
	}

	log.SetOutput(logfile)
	log.SetFlags(log.LstdFlags | log.LUTC)
}

func getInfo() (funcName string, fileName string, line int) {
	var pc uintptr
	var file string
	pc, file, line, _ = runtime.Caller(3)
	funcName = runtime.FuncForPC(pc).Name()
	funcName = funcName[strings.LastIndex(funcName, "/")+1:]
	fileName = file[strings.LastIndex(file, "/")+1:]
	return
}

func fatalf(message string) {
	funcName, fileName, line := getInfo()
	log.Fatalf(pattern, "FATAL", funcName, fileName, line, message)
}

func panicf(message string) {
	funcName, fileName, line := getInfo()
	log.Panicf(pattern, "PANIC", funcName, fileName, line, message)
}

func printf(level string, message string) {
	funcName, fileName, line := getInfo()
	log.Printf(pattern, level, funcName, fileName, line, message)
}

func Fatal(message string) {
	fatalf(message)
}

func Panic(message string) {
	panicf(message)
}

func Error(message string) {
	printf("ERROR", message)
}

func Warn(message string) {
	printf("WARN", message)
}

func Info(message string) {
	printf("INFO", message)
}

func Debug(message string) {
	printf("DEBUG", message)
}

func Trace(message string) {
	printf("TRACE", message)
}
