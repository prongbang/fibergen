# FiberGen

Generate a Clean Architecture for REST API with support for the Fiber Web Framework in Golang

[![Build Status](http://img.shields.io/travis/prongbang/fibergen.svg)](https://travis-ci.org/prongbang/fibergen)
[![Codecov](https://img.shields.io/codecov/c/github/prongbang/fibergen.svg)](https://codecov.io/gh/prongbang/fibergen)
[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/fibergen)](https://goreportcard.com/report/github.com/prongbang/fibergen)

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/prongbang)

## Install

```shell
go install github.com/prongbang/fibergen@latest
```

## Requirement

1. New Project

`-new`  project name

`-mod`  module name

```shell
fibergen -new test -mod github.com/prongbang
```

Example:

- go.mod

```
module github.com/prongbang/test
```

Structure

```
test
├── go.mod
├── go.sum
└── internal
    └── test
        ├── api
        │   ├── feature/domain
        │   ├── api.go
        │   ├── routers.go
        │   ├── wire.go
        │   └── wire_gen.go
        └── database
            └── drivers.go
```

2. mark `+fibergeen`

- wire.go

```golang
//+build wireinject

package api

import (
	"github.com/google/wire"
	//+fibergen:import wire:package
)

func CreateAPI(dbDriver database.Drivers) API {
	wire.Build(
		NewAPI,
		NewRouters,
		//+fibergen:func wire:build
	)
	return nil
}
```

- routers.go

```golang
package api

import (
	"github.com/gofiber/fiber/v2"
	//+fibergen:import routers:package
)

type Routers interface {
	core.Routers
}

type routers struct {
	//+fibergen:struct routers
}

func (r *routers) Initials(app *fiber.App) {
	//+fibergen:func initials
}

func NewRouters(
	//+fibergen:func new:routers
) Routers {
	return &routers{
		//+fibergen:return &routers
	}
}
```

## How to use

`-f`  feature name

```shell script
fibergen -f user
```
OR

```shell script
$ cd project/internal/app/api && fibergen -f user
```

## Output

```
user
├── datasource.go
├── handler.go
├── provider.go
├── repository.go
├── router.go
├── usecase.go
├── user.go
└── validate.go
```

## CRUD

- define spec `auth.json`

```json
{
   "accessToken": "JWT",
   "expired": 1234567,
}
```

- gen

```shell
fibergen -crud auth -s spec/auth.json -d mariadb
```
