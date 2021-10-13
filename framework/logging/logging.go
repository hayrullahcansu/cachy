package logging

import (
	"io"
	"log"
	"os"
	"sync"
)

var _instance *Logger

var _once sync.Once

func initialLogger() {

	// arguments stuffs
	// config.SetConfigFilePath(*configPath)
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	_instance = &Logger{
		WarningLogger: log.New(multiWriter, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLogger:    log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger:   log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func Instance() *Logger {
	_once.Do(initialLogger)
	return _instance
}

func Info(args ...interface{}) {
	Instance().InfoLogger.Println(args...)
}

func Warning(args ...interface{}) {
	Instance().WarningLogger.Println(args...)
}

func Error(args ...interface{}) {
	Instance().ErrorLogger.Println(args...)
}

func Fatal(args ...interface{}) {
	Instance().ErrorLogger.Println(args...)
	os.Exit(-1)
}

func Infof(format string, args ...interface{}) {
	Instance().InfoLogger.Printf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	Instance().WarningLogger.Printf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	Instance().ErrorLogger.Printf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	Instance().ErrorLogger.Printf(format, args...)
	os.Exit(-1)
}
