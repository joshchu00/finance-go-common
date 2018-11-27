package logger

import (
	"fmt"
	"log"
	"os"
)

func Init(logDirectory string, service string) {

	logfile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logDirectory, service), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("FATAL", "Open log file error:", err)
	}

	log.SetOutput(logfile)
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
}
