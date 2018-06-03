package fiscalog

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	TraceLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	LogFile       *os.File
)

func init() {
	logFilePath := "./fiscaluno.log"
	LogFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open log file: ", err)
	}

	TraceLogger = log.New(io.MultiWriter(LogFile, ioutil.Discard),
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	WarningLogger = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	ErrorLogger = log.New(io.MultiWriter(LogFile, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}

func Trace(context string) {
	TraceLogger.Println(context)
	defer LogFile.Close()
}

func Info(context string) {
	InfoLogger.Println(context)
	defer LogFile.Close()
}

func Warning(context string) {
	WarningLogger.Println(context)
	defer LogFile.Close()
}

func Error(context interface{}) {
	ErrorLogger.Println(context)
	defer LogFile.Close()
}
