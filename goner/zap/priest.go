package gone_zap

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/config"
	"github.com/gone-io/gone/goner/tracer"
)

func Priest(cemetery gone.Cemetery) error {
	t := cemetery.GetTomById(gone.IdGoneLogger)
	if t != nil && t.GetGoner().(gone.Logger) != gone.GetSimpleLogger() {
		_, ok := t.GetGoner().(*sugar)
		if !ok {
			t.GetGoner().(gone.Logger).Warn("logger is loaded, zap logger not used")
		}
		return nil
	}

	cemetery.BuryOnce(NewZapLogger())
	_ = config.Priest(cemetery)
	_ = tracer.Priest(cemetery)
	return cemetery.ReplaceBury(NewSugar())
}
