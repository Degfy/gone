// Code generated by gone; DO NOT EDIT.
package internal
import (
    "web-mysql/internal/controller"
    "web-mysql/internal/middleware"
    "web-mysql/internal/module/demo"
    "web-mysql/internal/router"
    "github.com/gone-io/gone"
)

func Priest(cemetery gone.Cemetery) error {
    cemetery.Bury(controller.NewDemoController())
    cemetery.Bury(middleware.NewAuthorizeMiddleware())
    cemetery.Bury(middleware.NewPlantMiddleware())
    cemetery.Bury(demo.NewDb())
    cemetery.Bury(demo.NewDemoService())
    cemetery.Bury(router.NewAuth())
    cemetery.Bury(router.NewPubRouter())
	return nil
}
