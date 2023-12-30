//go:build wireinject
// +build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/prongbang/mvp/test_test/internal/test_test/database"
	"github.com/prongbang/mvp/test_test/internal/test_test/api/login"
	"github.com/prongbang/mvp/test_test/internal/test_test/api/auth"
	"github.com/prongbang/mvp/test_test/internal/test_test/api/forgot"
	"github.com/prongbang/mvp/test_test/internal/test_test/api/otp"
	//+fibergen:import wire:package
)

func CreateAPI(dbDriver database.Drivers) API {
	wire.Build(
		NewAPI,
		NewRouters,
		login.ProviderSet,
		auth.ProviderSet,
		forgot.ProviderSet,
		otp.ProviderSet,
		//+fibergen:func wire:build
	)
	return nil
}