package casbinx

import (
	"fmt"
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
	fmt.Println("Policy api loaded.")
	return fibercasbinrest.NewDefault(e, viper.GetString("jwt.secret"))
}
