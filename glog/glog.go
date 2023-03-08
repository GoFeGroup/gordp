package glog

import (
	"fmt"
	"log"
	"os"
)

type LEVEL int

const (
	DEBUG LEVEL = iota
	INFO
	WARN
	ERROR
	NONE
)

var (
	logger *log.Logger
	level  = DEBUG
)

func init() {
	logger = log.New(os.Stdout, "", 0)
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func SetLogger(l *log.Logger) {
	logger = l
}

func SetLevel(l LEVEL) {
	level = l
}

func logPrint(prefix string, v ...any) {
	var arr []interface{}
	arr = append(arr, prefix)
	arr = append(arr, v...)
	_ = logger.Output(3, fmt.Sprintln(arr...))
}

func Debug(v ...any) {
	if level <= DEBUG {
		logPrint("[DEBUG]", v)
	}
}

func Debugf(f string, v ...interface{}) {
	if level <= DEBUG {
		logPrint("[DEBUG]", fmt.Sprintf(f, v...))
	}
}

func Info(v ...any) {
	if level <= DEBUG {
		logPrint("[INFO]", v)
	}
}

func Infof(f string, v ...interface{}) {
	if level <= DEBUG {
		logPrint("[INFO]", fmt.Sprintf(f, v...))
	}
}

func Warn(v ...any) {
	if level <= DEBUG {
		logPrint("[WARN]", v)
	}
}

func Warnf(f string, v ...interface{}) {
	if level <= DEBUG {
		logPrint("[WARN]", fmt.Sprintf(f, v...))
	}
}

func Error(v ...any) {
	if level <= DEBUG {
		logPrint("[ERROR]", v)
	}
}

func Errorf(f string, v ...interface{}) {
	if level <= DEBUG {
		logPrint("[ERROR]", fmt.Sprintf(f, v...))
	}
}
