package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/mvp/test_test/pkg/core"
	"github.com/prongbang/mvp/test_test/internal/test_test/api/login"
	"github.com/prongbang/mvp/test_test/internal/test_test/api/auth"
	"github.com/prongbang/mvp/test_test/internal/test_test/api/forgot"
	"github.com/prongbang/mvp/test_test/internal/test_test/api/otp"
	//+fibergen:import routers:package
)

type Routers interface {
	core.Routers
}

type routers struct {
	LoginRoute login.Router
	AuthRoute auth.Router
	ForgotRoute forgot.Router
	OtpRoute otp.Router
	//+fibergen:struct routers
}

func (r *routers) Initials(app *fiber.App) {
	r.LoginRoute.Initial(app)
	r.AuthRoute.Initial(app)
	r.ForgotRoute.Initial(app)
	r.OtpRoute.Initial(app)
	//+fibergen:func initials
}

func NewRouters(
	loginRoute login.Router,
	authRoute auth.Router,
	forgotRoute forgot.Router,
	otpRoute otp.Router,
	//+fibergen:func new:routers
) Routers {
	return &routers{
		LoginRoute: loginRoute,
		AuthRoute: authRoute,
		ForgotRoute: forgotRoute,
		OtpRoute: otpRoute,
		//+fibergen:return &routers
	}
}