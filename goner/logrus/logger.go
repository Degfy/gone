package logrus

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/tracer"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func NewLogger() (gone.Goner, gone.GonerId, gone.GonerOption) {
	log := &logger{
		Logger: logrus.StandardLogger(),
	}
	log.ResetLog()
	return log, gone.IdGoneLogger, gone.IsDefault(new(gone.Logger))
}

type logger struct {
	gone.Flag

	*logrus.Logger
	Tracer           tracer.Tracer `gone:"gone-tracer"`
	ConfLevel        string        `gone:"config,log.level,default=info"`
	ConfReportCaller bool          `gone:"config,log.report-caller,default=true"`
	ConfOutput       string        `gone:"config,log.output,default=stdout"`
}

func (log *logger) ResetLog() {
	if log.Tracer != nil {
		log.Formatter = &DefaultFormatter{
			GetTraceId: func() string {
				return log.Tracer.GetTraceId()
			},
		}
	} else {
		log.Formatter = &DefaultFormatter{
			GetTraceId: func() string {
				return "Init"
			},
		}
	}
	log.ReportCaller = log.ConfReportCaller
	log.Level = parseLogLevel(log.ConfLevel)
	log.Out = parseOutput(log.ConfOutput)

}
func (log *logger) AfterRevive() gone.AfterReviveError {
	log.ResetLog()
	return nil
}

func parseLogLevel(level string) logrus.Level {
	if level == "" {
		level = "info"
	}
	var l logrus.Level
	err := l.UnmarshalText([]byte(level))
	if err != nil {
		return logrus.InfoLevel
	}
	return l
}

func parseOutput(output string) io.Writer {
	switch output {
	case "stdout", "":
		return os.Stdout
	case "stderr":
		return os.Stderr
	default:
		f, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(gone.ToError(err))
		}
		return f
	}
}
