package logging

import (
	"log"
	"os"
)

var (
	ErrFile   *log.Logger
	DebugFile *log.Logger
)

func init() {
	var errsLogFile, debugLogFile *os.File
	var err error
	if errsLogFile, err = os.OpenFile("errs.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm); err != nil {
		log.Fatalln("Init failed, reason:", err)
	}
	if debugLogFile, err = os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm); err != nil {
		log.Fatalln("Init failed, reason:", err)
	}
	ErrFile = log.New(errsLogFile, "ERROR", log.LstdFlags)
	DebugFile = log.New(debugLogFile, "DEBUG", log.LstdFlags)
}
