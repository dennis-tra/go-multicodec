# 🚧 Archived

This project has been moved upstream to [multiformats/go-multicodec](https://github.com/multiformats/go-multicodec/pull/37/commits/4694ef8846ef034239c573af2ef13a54d22d2a9f).

# go-multicodec

![Multicodecs update](https://github.com/dennis-tra/go-multicodec/workflows/Multicodecs%20update/badge.svg) [![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme) [![Go Report Card](https://goreportcard.com/badge/dennis-tra/go-multicodec)](https://goreportcard.com/report/dennis-tra/go-multicodec)

> Periodically self-generated Go constants of [multicodecs](https://github.com/multiformats/multicodec) used by the [multiformats](https://github.com/multiformats/multiformats) projects.

## Table of Contents

- [Motivation](#motivation)
- [Workflow](#workflow)
- [Install](#install)
- [Usage](#usage)
- [Generator](#generator)
- [Maintainers](#maintainers)
- [License](#license)

## Motivation

Consolidation of multiple constant definitions. E.g.

- [multiformats/go-multihash/multihash.go#L38](https://github.com/multiformats/go-multihash/blob/6f1ea18f1da5f7735ea31b5e2011da61c409e37f/multihash.go#L38) or [not obvious](https://github.com/multiformats/go-multihash/issues/53#issuecomment-313360164) constant [generation](https://github.com/multiformats/go-multihash/blob/master/multihash.go#L78)
- [multiformats/go-cid/cid.go#L52](https://github.com/ipfs/go-cid/blob/e530276a7008f5973e7da6640ed305ecc5825d27/cid.go#L52)
- maybe more?

## Workflow

Every night a few minutes after midnight (UTC) a GitHub-Action fetches the latest [multicodecs table](https://raw.githubusercontent.com/multiformats/multicodec/master/table.csv), generates the constants, commits possible changes and creates a pull request. It will update the same pull request if subsequent runs find different changes and it won't create a pull request if no changes were detected.

## Install

`go-multicodec` is a standard Go module which can be installed with:

```sh
go get github.com/dennis-tra/go-multicodec
```

## Usage

```go
package main

import "github.com/dennis-tra/go-multicodec"

func main() {
    _ = multicodec.Sha2_256
}
```

The corresponding `name` value for each codec from the [multicodecs table](https://raw.githubusercontent.com/multiformats/multicodec/master/table.csv) can be accessed via its `String()` method. E.g. `multicodec.Sha2_256.String()` will return `sha2-256`.

## Generator

To generate the constants yourself checkout the repository like

```shell
git clone https://github.com/dennis-tra/go-multicodec.git
```

Then run `go generate` in the root of the repository. Currently this will invoke (see [`init.go`](./init.go)):

```shell
go run ./gen/gen.go
gofmt -w codec.go
stringer -type=Codec -linecomment
```

Note: You may need to install `stringer` via `go install golang.org/x/tools/cmd/stringer`.

## Maintainers

[@dennis-tra](https://github.com/dennis-tra).

## License

[MIT](LICENSE) © Dennis Trautwein
