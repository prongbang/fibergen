# FiberGen 🚀

[![Codecov](https://img.shields.io/codecov/c/github/prongbang/fibergen.svg)](https://codecov.io/gh/prongbang/fibergen)
[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/fibergen)](https://goreportcard.com/report/github.com/prongbang/fibergen)
[![Go Reference](https://pkg.go.dev/badge/github.com/prongbang/fibergen.svg)](https://pkg.go.dev/github.com/prongbang/fibergen)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Version](https://img.shields.io/github/v/release/prongbang/fibergen)](https://github.com/prongbang/fibergen/releases)

> Generate Clean Architecture for REST API with support for Fiber Web Framework and gRPC in Golang. Speed up your development with automatic code generation.


## ✨ Features

- 🏗️ **Clean Architecture** - Automatically generates layered architecture structure
- 🔌 **Fiber Framework Support** - Optimized for the Fiber web framework
- 🌐 **gRPC Support** - Built-in support for gRPC services (**in-progress**)
- 🔄 **CRUD Generation** - Generate CRUD operations from JSON specifications
- 🧩 **Modular Design** - Feature-based modules for better organization
- 🔧 **Wire Integration** - Dependency injection with Google Wire
- ⚡ **Fast Development** - Speed up your development workflow

## 📦 Installation

Latest version:
```shell
go install github.com/prongbang/fibergen@v1.3.0
```

For older projects:
```shell
go install github.com/prongbang/fibergen@v1.2.5
```

## 🚀 Quick Start

### 1. Create a New Project

Generate a new project with module structure:

```shell
fibergen -new test_project -mod github.com/prongbang
```

Parameters:
- `-new`: Project name
- `-mod`: Module name

This creates the following structure:

```
test_project/
├── go.mod
├── go.sum
├── wire.go
├── wire_gen.go
└── internal/
    ├── app/
    │   ├── app.go
    │   ├── grpc/
    │   │   ├── featurename/
    │   │   ├── grpc.go
    │   │   └── servers.go
    │   └── api/
    │       ├── featurename/
    │       ├── api.go
    │       └── routers.go
    └── database/
        └── drivers.go
```

### 2. Mark Generation Points

Add `+fibergen` markers to your code:

**wire.go**
```go
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

**routers.go**
```go
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

### 3. Generate Features

Generate a new feature module:

```shell
fibergen -f user
```

This creates:
```
user/
├── datasource.go
├── handler.go
├── provider.go
├── repository.go
├── router.go
├── usecase.go
└── user.go
```

## 🔄 CRUD Generation

Generate CRUD operations from JSON specifications:

### 1. Define Spec File

Create `spec/auth.json`:
```json
{
    "accessToken": "JWT",
    "expired": 1234567,
    "date": "2024-10-15T14:30:00Z"
}
```

### 2. Generate CRUD

```shell
fibergen -crud auth -s spec/auth.json
```

This generates complete CRUD operations based on your JSON structure.

## 📁 Generated Structure

Each feature generates the following components:

### 1. DataSource Layer
`datasource.go` - Database operations
```go
type DataSource interface {
    // Generated methods
}

type dataSource struct {
    DbDriver database.Drivers
}

func NewDataSource(dbDriver database.Drivers) DataSource {
    return &dataSource{
        DbDriver: dbDriver,
    }
}
```

### 2. Repository Layer
`repository.go` - Business logic repository
```go
type Repository interface {
    // Generated methods
}

type repository struct {
    Ds DataSource
}

func NewRepository(ds DataSource) Repository {
    return &repository{
        Ds: ds,
    }
}
```

### 3. UseCase Layer
`usecase.go` - Business logic
```go
type UseCase interface {
    // Generated methods
}

type useCase struct {
    Repo Repository
}

func NewUseCase(repo Repository) UseCase {
    return &useCase{
        Repo: repo,
    }
}
```

### 4. Handler Layer
`handler.go` - HTTP handlers
```go
type Handler interface {
    // Generated methods
}

type handler struct {
    Uc UseCase
}

func NewHandler(uc UseCase) Handler {
    return &handler{
        Uc: uc,
    }
}
```

### 5. Router Configuration
`router.go` - Route definitions
```go
type Router interface {
    Initial(app *fiber.App)
}

type router struct {
    Handle Handler
}

func (r *router) Initial(app *fiber.App) {
    // Generated routes
}

func NewRouter(handle Handler) Router {
    return &router{Handle: handle}
}
```

## 🔧 Advanced Usage

### Custom Markers

You can use various markers to customize generation:

- `//+fibergen:import wire:package` - Import packages
- `//+fibergen:func wire:build` - Wire build functions
- `//+fibergen:struct routers` - Struct fields
- `//+fibergen:func initials` - Function implementations
- `//+fibergen:func new:routers` - Constructor parameters
- `//+fibergen:return &routers` - Return values

### gRPC Support

Generate gRPC services alongside REST APIs:

```shell
in-progress
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 💖 Support

If you find this tool helpful, please consider buying me a coffee:

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/prongbang)

## 🔗 Related Projects

- [Fiber](https://github.com/gofiber/fiber) - Express-inspired web framework
- [Wire](https://github.com/google/wire) - Compile-time dependency injection
- [gRPC](https://grpc.io/) - High performance RPC framework

---
