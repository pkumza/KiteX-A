package zlog

import (
	"fmt"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
)

type Zlog struct {
	logFile *os.File
}

func New() klog.FormatLogger {
	z := &Zlog{}
	logDir := getLogDir()
	logFile, err := os.Create(logDir + "/app.log")
	if err != nil {
		panic(err)
	}
	z.logFile = logFile
	return z
}

var (
	EnvLogDir     = "KITEX_LOG_DIR"
	DefaultLogDir = "log"
)

// getLogDir gets dir of log file.
func getLogDir() string {
	if logDir := os.Getenv(EnvLogDir); logDir != "" {
		return logDir
	}
	return DefaultLogDir
}

func (z *Zlog) Tracef(format string, v ...interface{}) {
	fmt.Fprintf(z.logFile, "TRACE "+format+"\n", v...)
}
func (z *Zlog) Debugf(format string, v ...interface{}) {
	fmt.Fprintf(z.logFile, "DEBUG "+format+"\n", v...)
}
func (z *Zlog) Infof(format string, v ...interface{}) {
	fmt.Fprintf(z.logFile, "INFO "+format+"\n", v...)
}
func (z *Zlog) Noticef(format string, v ...interface{}) {
	fmt.Fprintf(z.logFile, "NOTICE "+format+"\n", v...)
}
func (z *Zlog) Warnf(format string, v ...interface{}) {
	fmt.Fprintf(z.logFile, "WARN "+format+"\n", v...)
}
func (z *Zlog) Errorf(format string, v ...interface{}) {
	fmt.Fprintf(z.logFile, "ERROR "+format+"\n", v...)
}
func (z *Zlog) Fatalf(format string, v ...interface{}) {
	fmt.Fprintf(z.logFile, "FATAL "+format+"\n", v...)
}
