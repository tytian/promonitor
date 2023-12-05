package server

import (
	"encoding/json"
	"fmt"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

const (
	logDir = "/var/log/monitor"
)

type MyJsonFormatter struct {
	Time  string `json:"time"`
	File  string `json:"file"`
	Line  int    `json:"line"`
	Level string `json:"level"`
	Info  string `json:"info"`
	Msg   string `json:"msg"`
}

func (m *MyJsonFormatter) Format(entry *log.Entry) ([]byte, error) {
	logrusJF := &(log.JSONFormatter{})
	logrusJF.TimestampFormat = "2006-01-02 15:04:05.000"
	bytes, _ := logrusJF.Format(entry)
	json.Unmarshal(bytes, &m)
	if _, file, no, ok := runtime.Caller(8); ok {
		m.File = file
		m.Line = no
	}
	str := fmt.Sprintf("[%s] %s %s:%d %s\n", m.Level, m.Time, m.File, m.Line, m.Msg)
	return []byte(str), nil
}

func InitLog(level string) {
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.Fatal("log conf only allow [debug, info, warn, error], please check your configuration")
	}
	os.MkdirAll(logDir, os.ModePerm)
	path := filepath.Base(os.Args[0])
	var logname = logDir + "/" + path + ".log"
	log.SetOutput(ioutil.Discard)
	log.AddHook(lfshook.NewHook(
		lfshook.PathMap{
			log.InfoLevel:  logname,
			log.DebugLevel: logname,
			log.WarnLevel:  logname,
			log.ErrorLevel: logname,
			log.FatalLevel: logname,
			log.PanicLevel: logname,
		},
		&MyJsonFormatter{}),
	)
}
