package template

type internalPkgCasbinxTemplate struct {
}

func (i *internalPkgCasbinxTemplate) Text() []byte {
	return []byte(`package casbinx

import (
	"log"
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	fibercasbinrest "github.com/prongbang/fiber-casbinrest"
	"github.com/spf13/viper"
)

func New() fiber.Handler {
	e, err := casbin.NewEnforcer(viper.Get("casbin.model"), viper.Get("casbin.policy"))
	if err != nil {
		panic(err)
	}
	log.Println("Policy API loaded.")
	return fibercasbinrest.NewDefault(e, viper.GetString("jwt.secret"))
}
`)
}

func InternalPkgCasbinxTemplate() Template {
	return &internalPkgCasbinxTemplate{}
}
