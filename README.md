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
- 🛠️ **Swagger CRUD Generation** - Generate Swagger CRUD operations
- 🧩 **Modular Design** - Feature-based modules for better organization
- 🔧 **Wire Integration** - Dependency injection with Google Wire
- ⚡ **Fast Development** - Speed up your development workflow

## 📦 Installation

Latest version:
```shell
go install github.com/prongbang/fibergen@v1.4.3
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
.
├── Makefile
├── cmd
│     └── api
│         └── main.go
├── configuration
│     ├── configuration.go
│     ├── development.yml
│     ├── environment.go
│     └── production.yml
├── deployments
│     ├── Dockerfile
│     └── api-prod.yml
├── docs
│     └── apispec
│         ├── docs.go
│         ├── swagger.json
│         └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│     ├── app
│     │     ├── api
│     │     │     ├── api.go
│     │     │     ├── example
│     │     │     │     ├── datasource.go
│     │     │     │     ├── example.go
│     │     │     │     ├── handler.go
│     │     │     │     ├── permission.go
│     │     │     │     ├── provider.go
│     │     │     │     ├── repository.go
│     │     │     │     ├── router.go
│     │     │     │     └── usecase.go
│     │     │     └── routers.go
│     │     └── app.go
│     ├── database
│     │     ├── db.go
│     │     ├── drivers.go
│     │     ├── mariadb.go
│     │     ├── mongodb.go
│     │     ├── wire.go
│     │     └── wire_gen.go
│     ├── middleware
│     │     ├── api_key.go
│     │     ├── jwt.go
│     │     └── on_request.go
│     ├── pkg
│     │     ├── casbinx
│     │     │     └── casbinx.go
│     │     ├── response
│     │     │     └── response.go
│     │     └── validator
│     │         └── validator.go
│     └── shared
│           └── example
│               ├── datasource.go
│               ├── example.go
│               ├── provider.go
│               └── repository.go
│
├── pkg
│     ├── core
│     │     ├── common.go
│     │     ├── flag.go
│     │     ├── handler.go
│     │     ├── header.go
│     │     ├── jwt.go
│     │     ├── paging.go
│     │     ├── params.go
│     │     ├── request.go
│     │     ├── response.go
│     │     ├── router.go
│     │     └── sorting.go
│     ├── multipartx
│     │     └── multipartx.go
│     ├── requestx
│     │     └── request.go
│     ├── schema
│     │     └── sql.go
│     ├── streamx
│     │     └── streamx.go
│     ├── structx
│     │     └── structx.go
│     └── typex
│         └── typex.go
├── policy
│     ├── model.conf
│     └── policy.csv
├── spec
│     └── promotion.json
├── wire.go
└── wire_gen.go
```

### 2. Generate Features Prototype

Generate a new feature module:

```shell
fibergen -f user
```

This creates:
```
test-project/internal/app/api/promotion
├── datasource.go
├── handler.go
├── permission.go
├── promotion.go
├── provider.go
├── repository.go
├── router.go
└── usecase.go
```

### 3. Generate Features CRUD and Swagger

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

- SQL Builder

```shell
fibergen -f auth -s spec/auth.json -d mariadb -orm sqlbuilder
```

- Bun

```shell
fibergen -f auth -s spec/auth.json -d mariadb -orm bun
```

This generates complete CRUD operations based on your JSON structure.

```
test-project/internal/app/api/promotion
├── datasource.go
├── handler.go
├── permission.go
├── promotion.go
├── provider.go
├── repository.go
├── router.go
└── usecase.go
```

### 4. Generate Shared Prototype

```shell
fibergen -sh promotion
```
This generates shared prototype

```
test-project/internal/shared/promotion
├── datasource.go
├── promotion.go
├── provider.go
└── repository.go
```

### 5. Generate Shared CRUD

- SQL Builder

```shell
fibergen -sh promotion -s spec/promotion.json -d maridb -orm sqlbuilder
```

- Bun

```shell
fibergen -sh promotion -s spec/promotion.json -d maridb -orm bun
```

This generates shared CRUD operations based on your JSON structure.

```
test-project/internal/shared/promotion
├── datasource.go
├── promotion.go
├── provider.go
└── repository.go
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
