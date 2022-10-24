package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")
	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some Program")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Println("LOG MAIL: Logging in Go!")
	fmt.Println("Will you see this ?")
}
