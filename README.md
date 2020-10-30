# go-map
> Library for working with [Magic Attribute Protocol](https://github.com/rohenaz/MAP) in Go and used in conjunction with a Tape from [go-bob](https://github.com/bitcoinschema/go-bob)

[![Release](https://img.shields.io/github/release-pre/BitcoinSchema/go-map.svg?logo=github&style=flat&v=3)](https://github.com/BitcoinSchema/go-map/releases)
[![Build Status](https://travis-ci.com/BitcoinSchema/go-map.svg?branch=master&v=3)](https://travis-ci.com/BitcoinSchema/go-map)
[![Report](https://goreportcard.com/badge/github.com/BitcoinSchema/go-map?style=flat&v=3)](https://goreportcard.com/report/github.com/BitcoinSchema/go-map)
[![codecov](https://codecov.io/gh/BitcoinSchema/go-map/branch/master/graph/badge.svg?v=3)](https://codecov.io/gh/BitcoinSchema/go-map)
[![Go](https://img.shields.io/github/go-mod/go-version/BitcoinSchema/go-map?v=3)](https://golang.org/)
[![Sponsor](https://img.shields.io/badge/sponsor-BitcoinSchema-181717.svg?logo=github&style=flat&v=3)](https://github.com/sponsors/BitcoinSchema)
[![Donate](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat&v=3)](https://gobitcoinsv.com/#sponsor)

<br/>

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

<br/>

## Installation

**go-map** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/bitcoinschema/go-map
```

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/bitcoinschema/go-map)

[![GoDoc](https://godoc.org/github.com/bitcoinschema/go-map?status.svg&style=flat)](https://pkg.go.dev/github.com/bitcoinschema/go-map)

### Features
- [NewFromTape()](bob.go)
- Support Commands:
  - [SET](magic.go)
  - [ADD](magic.go)
  - [DELETE](magic.go)
  - [REMOVE](magic.go)
  - [SELECT](magic.go)

<details>
<summary><strong><code>Package Dependencies</code></strong></summary>
<br/>

- [bitcoinschema/go-bob](https://github.com/bitcoinschema/go-bob)
</details>

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
all                    Runs multiple commands
clean                  Remove previous builds and any test cache data
clean-mods             Remove all the Go mod cache
coverage               Shows the test coverage
godocs                 Sync the latest tag with GoDocs
help                   Show this help message
install                Install the application
install-go             Install the application (Using Native Go)
lint                   Run the Go lint application
release                Full production release (creates release in Github)
release                Runs common.release then runs godocs
release-snap           Test the full release (build binaries)
release-test           Full production test release (everything except deploy)
replace-version        Replaces the version in HTML/JS (pre-deploy)
tag                    Generate a new tag and push (tag version=0.0.0)
tag-remove             Remove a tag if found (tag-remove version=0.0.0)
tag-update             Update an existing tag to current commit (tag-update version=0.0.0)
test                   Runs vet, lint and ALL tests
test-short             Runs vet, lint and tests (excludes integration tests)
test-travis            Runs all tests via Travis (also exports coverage)
test-travis-short      Runs unit tests via Travis (also exports coverage)
uninstall              Uninstall the application (and remove files)
vet                    Run the Go vet application
```
</details>

<br/>

## Examples & Tests
All unit tests and [examples](examples) run via [Travis CI](https://travis-ci.com/bitcoinschema/go-map) and uses [Go version 1.15.x](https://golang.org/doc/go1.15). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```shell script
make test
```

Run tests (excluding integration tests)
```shell script
make test-short
```

<br/>

## Benchmarks
Run the Go benchmarks:
```shell script
make bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

<br/>

## Usage
Checkout all the [examples](examples)!

```go
    import "github.com/bitcoinschema/go-bob"
    import "github.com/bitcoinschema/go-map"

    line := "<BOB formatted json string>"

    bobData := &bob.Tx{}
    if err := json.Unmarshal(line, bobData); err != nil {
      fmt.Println("Error:", err)
      return
    }

    for _, out := range bobData.Out {
      for _, tape := range out.Tape {
        mapData, err := magic.NewFromTape(tape)
        log.Printf("MAP TYPE is %s", mapData["type"])
      }
    }


```

<br/>

## Maintainers
| [<img src="https://github.com/rohenaz.png" height="50" alt="MrZ" />](https://github.com/rohenaz) | [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:---:|:---:|
| [Satchmo](https://github.com/rohenaz) | [MrZ](https://github.com/mrz1836) |

<br/>

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:!
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:.
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/BitcoinSchema) :clap:
or by making a [**bitcoin donation**](https://gobitcoinsv.com/#sponsor) to ensure this journey continues indefinitely! :rocket:


### Credits
[Siggi](https://github.com/icellan) for creating [BAP](https://github.com/icellan/bap) :clap:

<br/>

## License

![License](https://img.shields.io/github/license/BitcoinSchema/go-map.svg?style=flat&v=3)