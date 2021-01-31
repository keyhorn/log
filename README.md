# log

[![Test Status](https://github.com/keyhorn/log/workflows/test/badge.svg)](https://github.com/keyhorn/log/actions?query=workflow%3Atest)
[![Lint Status](https://github.com/keyhorn/log/workflows/lint/badge.svg)](https://github.com/keyhorn/log/actions?query=workflow%3Alint)
[![Coverage Status](https://coveralls.io/repos/github/keyhorn/log/badge.svg?branch=main)](https://coveralls.io/github/keyhorn/log?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/keyhorn/log)](https://goreportcard.com/report/github.com/keyhorn/log)
[![BCH compliance](https://bettercodehub.com/edge/badge/keyhorn/log?branch=main)](https://bettercodehub.com/)
[![Documentation](https://godoc.org/github.com/keyhorn/log?status.svg)](http://godoc.org/github.com/keyhorn/log)
[![License](https://img.shields.io/github/license/keyhorn/log.svg?maxAge=2592000)](https://github.com/keyhorn/log/LICENSE)
[![Release](https://img.shields.io/github/release/keyhorn/log.svg?label=Release)](https://github.com/keyhorn/log/releases)

logging library for Go

## Install

Use `go get` to install this library.

```shell
go get -u github.com/keyhorn/log
```

## API Document

See [GoDoc](https://godoc.org/github.com/keyhorn/log) for full doument.

## Usage

```golang
package main

import (
    "github.com/keyhorn/log"
)

func main() {
    var logger = log.New(log.StdoutWriter())

    logger.Print("Hello, log file!")
}
```

## License

This library is licensed under MIT license. See [LICENSE](https://github.com/keyhorn/log/LICENSE) for details.
