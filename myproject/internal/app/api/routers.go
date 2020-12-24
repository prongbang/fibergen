package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/innotechdevops/apirole/pkg/apirole"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	//+fibergen:import routers:package
)

type Routers interface {
	core.Routers
}

type routers struct {
	RoleRoute apirole.Router
	UserRoute user.Router
	//+fibergen:struct routers
}

func (r *routers) Initials(app *fiber.App) {
	r.RoleRoute.Initial(app)
	r.UserRoute.Initial(app)
	//+fibergen:func initials
}

func NewRouters(
	roleRoute apirole.Router,
	userRoute user.Router,
	//+fibergen:func new:routers
) Routers {
	return &routers{
		RoleRoute: roleRoute,
		UserRoute: userRoute,
		//+fibergen:return &routers
	}
}
