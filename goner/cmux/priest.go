package cmux

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/config"
	"github.com/gone-io/gone/goner/logrus"
)

func Priest(cemetery gone.Cemetery) error {
	_ = logrus.Priest(cemetery)
	_ = config.Priest(cemetery)
	newServer, id := NewServer()
	gone.CheckAndBury(cemetery, newServer, id)
	return nil
}
