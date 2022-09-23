# Patrick

[![Build](https://github.com/blackhorseya/patrick/actions/workflows/build.yaml/badge.svg)](https://github.com/blackhorseya/patrick/actions/workflows/build.yaml)
[![codecov](https://codecov.io/gh/blackhorseya/patrick/branch/main/graph/badge.svg?token=qkP4xlJB12)](https://codecov.io/gh/blackhorseya/patrick)
[![Go Report Card](https://goreportcard.com/badge/github.com/blackhorseya/patrick)](https://goreportcard.com/report/github.com/blackhorseya/patrick)
[![Go Reference](https://pkg.go.dev/badge/github.com/blackhorseya/patrick)](https://pkg.go.dev/github.com/blackhorseya/patrick)
[![Release](https://img.shields.io/github/release/blackhorseya/patrick)](https://github.com/blackhorseya/patrick/releases/latest)
[![GitHub license](https://img.shields.io/github/license/blackhorseya/patrick)](https://github.com/blackhorseya/patrick/blob/main/LICENSE)

Patrick provides its own program that will create your application.

- [Patrick](#patrick)
  - [Installation](#installation)
  - [Usage](#usage)

## Installation

```shell
go install github.com/blackhorseya/patrick/cmd/patrick@latest
```

## Usage

First, you can show help the following command.

```shell
patrick -h
```

### patrick init

The `patrick init [app name]` command will create your initial application code for you.

#### Initializing a module

**If you already have a module, skip this step**

If you want to initialize a new Go module:

1. Create new directory
2. `cd` into that directory
3. run `go mod init <MOD_NAME>`

e.g.

```shell
cd $HOME/project
mkdir myapp
cd myapp
go mod init github.com/blackhorseya/myapp
```
