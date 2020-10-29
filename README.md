# FiberGen

Golang Generate Clean Architecture for REST API support Fiber Web Framework

[![Build Status](http://img.shields.io/travis/prongbang/fibergen.svg)](https://travis-ci.org/prongbang/fibergen)
[![Codecov](https://img.shields.io/codecov/c/github/prongbang/fibergen.svg)](https://codecov.io/gh/prongbang/fibergen)
[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/fibergen)](https://goreportcard.com/report/github.com/prongbang/fibergen)

## Install

```shell script
$ go install github.com/prongbang/fibergen
```

## How to use

`-f`  feature name

```shell script
$ fibergen -f user
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
└── user.go
```