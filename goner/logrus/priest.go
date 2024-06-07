package logrus

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/config"
	"github.com/gone-io/gone/goner/tracer"
)

func Priest(cemetery gone.Cemetery) error {
	_ = tracer.Priest(cemetery)
	_ = config.Priest(cemetery)

	theLogger, id, _ := NewLogger()
	return cemetery.ReplaceBury(theLogger, id)
}
