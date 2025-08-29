# 👻 Goth

[![Test & Build](https://github.com/katallaxie/fiber-goth/actions/workflows/main.yml/badge.svg)](https://github.com/katallaxie/fiber-goth/actions/workflows/main.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/katallaxie/fiber-goth.svg)](https://pkg.go.dev/github.com/katallaxie/fiber-goth)
[![Go Report Card](https://goreportcard.com/badge/github.com/katallaxie/fiber-goth)](https://goreportcard.com/report/github.com/katallaxie/fiber-goth)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)

A [fiber](https://gofiber.io/) :rocket: middleware to integrate authentication to your application. It uses lightweight `adapters` and `providers` interfaces to integrate with multi-providers. 

## Installation

```bash
$ go get github.com/katallaxie/fiber-goth/v3
```

## Providers

* GitHub (github.com, Enterprise, and Enterprise Cloud)
* Microsoft Entra ID

## CSRF

The middleware supports CSRF protection. It is added via the following package.

```golang
import "github.com/katallaxie/fiber-goth/v3/csrf"

app := fiber.New()
app.Use(csrf.New())
```

The CSRF protection depends on the session middleware.

## Examples

See [examples](https://github.com/katallaxie/fiber-goth/tree/master/examples) to understand the provided interfaces

## License

[MIT](/LICENSE)
