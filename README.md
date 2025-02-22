# FiberGen

Generate a Clean Architecture for REST API with support for the Fiber Web Framework in Golang

[![Build Status](http://img.shields.io/travis/prongbang/fibergen.svg)](https://travis-ci.org/prongbang/fibergen)
[![Codecov](https://img.shields.io/codecov/c/github/prongbang/fibergen.svg)](https://codecov.io/gh/prongbang/fibergen)
[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/fibergen)](https://goreportcard.com/report/github.com/prongbang/fibergen)

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/prongbang)

## Install

```shell
go install github.com/prongbang/fibergen@v1.3.0
```

or old project

```shell
go install github.com/prongbang/fibergen@v1.2.5
```

## Requirement

1. New Project

`-new`  project name

`-mod`  module name

```shell
fibergen -new test_project -mod github.com/prongbang
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
├── wire.go
├── wire_gen.go
└── internal
     ├── app
     │   ├── app.go
     │	 ├── grpc
     │   │   ├── featurename
     │   │   ├── grpc.go
     │   │   └── servers.go
     │   └── api
     │       ├── featurename
     │       ├── api.go
     │       └── routers.go
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
		New,
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
$ fibergen -f user
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
└── user.go
```

## CRUD Generate by .json file

- define spec `auth.json`

```json
{
   "accessToken": "JWT",
   "expired": 1234567,
   "date": "2024-10-15T14:30:00Z"
}
```

- gen

```shell
fibergen -crud auth -s spec/auth.json
```
