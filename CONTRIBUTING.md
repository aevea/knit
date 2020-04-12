# Contributing to Merge Master

## Table of Contents
- [Requirements](#requirements)
- [Commit style](#commit-style)
- [API Service](#API-service)
- [Tooling](#tooling)
- [Troubleshooting](#troubleshooting)

## Requirements

- Go (check pinned version in go.mod)
- [Mage](https://magefile.org/)

**How to install Mage:**
```
go install github.com/magefile/mage
```


## Commit style

This project uses [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/). This is automatically enforced by our tool [Commitsar](https://github.com/outillage/commitsar). Please make sure your commits follow this standard and are descriptive.


## API Service

The API is created using [oto](https://github.com/pacedotdev/oto). Please check out [THIS](https://pace.dev/blog/2020/02/26/tech-stack-at-pace.html) article to understand how to use it.

All definition files are housed under api/definitions. To generate API endpoint defintions from this please run: `make generate_oto`.

#### Troubleshooting
- Can't find `oto` binary: please run `make install_deps`.


## Tooling

For 3rd party tools such as `oto` we use the [go tools file approach](https://github.com/golang/go/issues/25922). For consideration in the future is this project: https://github.com/go-modules-by-example/index/tree/master/017_using_gobin

**Essentially boils down to this:**
- add your tool as a dependency to `tools.go`
- add the dependency to Install in `mage.go` and run `mage install`

**Reasons**

By using this approach we can version tools used in Merge master as a normal Go dependency.

## Troubleshooting

#### Can't find `mage`

Make sure your PATH points to GOBIN:
```
export PATH=${PATH}:`go env GOPATH`/bin
```