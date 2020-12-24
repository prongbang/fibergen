//+build wireinject

package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"innotechdev.co.th/azbil/ath-connected/internal/app/api/hello"
	"innotechdev.co.th/azbil/ath-connected/internal/app/api/login"
	"innotechdev.co.th/azbil/ath-connected/internal/app/api/monitor"
	"innotechdev.co.th/azbil/ath-connected/internal/app/api/role"
	"innotechdev.co.th/azbil/ath-connected/internal/app/database"
	"innotechdev.co.th/azbil/ath-connected/internal/app/thirdparty/azbilmqtt"
	"innotechdev.co.th/azbil/ath-connected/internal/app/thirdparty/emqx"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	"github.com/prongbang/fibergen/myproject/internal/app/api/user"
	//+fibergen:import wire:package
)

func CreateAPI(dbDriver database.Drivers, emqxSource emqx.DataSource, azbilMqttSource azbilmqtt.DataSource, enforcer *casbin.Enforcer) API {
	wire.Build(
		NewAPI,
		NewRouters,
		role.ProviderSet,
		hello.ProviderSet,
		login.ProviderSet,
		monitor.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		user.ProviderSet,
		//+fibergen:func wire:build
	)
	return nil
}
