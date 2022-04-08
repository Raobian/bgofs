package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var l *logrus.Logger

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

func init() {
	l = logrus.New()
}

func SetLevel(level Level) {
	switch level {
	case DEBUG:
		l.SetLevel(logrus.DebugLevel)
		break
	case INFO:
		l.SetLevel(logrus.InfoLevel)
		break
	case WARN:
		l.SetLevel(logrus.WarnLevel)
		break
	case ERROR:
		l.SetLevel(logrus.ErrorLevel)
		break
	case FATAL:
		l.SetLevel(logrus.FatalLevel)
		break
	case PANIC:
		l.SetLevel(logrus.PanicLevel)
		break
	default:
		break
	}
}

func SetReportCaller() {
	l.SetReportCaller(true)
}

func SetOutput(file *os.File) {
	l.Out = file
}

func SetJsonFormat() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func Dbug(i ...interface{}) {
	l.Debug(i...)
}

func Dinfo(i ...interface{}) {
	l.Info(i...)
}

func Dwarn(i ...interface{}) {
	l.Warn(i...)
}

func Derror(i ...interface{}) {
	l.Error(i...)
}

func Dfatal(i ...interface{}) {
	l.Fatal(i...)
}

func Dpanic(i ...interface{}) {
	l.Panic(i...)
}

func DBUG(f string, i ...interface{}) {
	l.Debugf(f, i...)
}

func DINFO(f string, i ...interface{}) {
	l.Infof(f, i...)
}

func DWARN(f string, i ...interface{}) {
	l.Warnf(f, i...)
}

func DERROR(f string, i ...interface{}) {
	l.Errorf(f, i...)
}

func DFATAL(f string, i ...interface{}) {
	l.Fatalf(f, i...)
}

func DPANIC(f string, i ...interface{}) {
	l.Panicf(f, i...)
}
