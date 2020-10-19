# go-multicodec

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme) [![Go Report Card](https://goreportcard.com/badge/dennis-tra/go-multicodec)](https://goreportcard.com/report/dennis-tra/go-multicodec)

Periodically self-generated Go constants of [multicodecs](https://github.com/multiformats/multicodec) used by many of the [multiformats](https://github.com/multiformats/multiformats) projects.

## Table of Contents

- [Motivation](#motivation)
- [Workflow](#workflow)
- [Install](#install)
- [Usage](#usage)
- [Generator](#generator)
- [Maintainers](#maintainers)
- [License](#license)

## Motivation

Multiple constant definitions. E.g.

- [multiformats/go-multihash/multihash.go#L38](https://github.com/multiformats/go-multihash/blob/6f1ea18f1da5f7735ea31b5e2011da61c409e37f/multihash.go#L38) or [not obvious](https://github.com/multiformats/go-multihash/issues/53#issuecomment-313360164) constant [generation](https://github.com/multiformats/go-multihash/blob/master/multihash.go#L78)
- [multiformats/go-cid/cid.go#L52](https://github.com/ipfs/go-cid/blob/e530276a7008f5973e7da6640ed305ecc5825d27/cid.go#L52)
- maybe more?

## Workflow

This repo contains the [multiformats/multicodec](https://github.com/multiformats/multicodec) repo as a submodule. Every night a few minutes after midnight a GitHub-Action updates the HEAD of the submodule to the most recent `master` commit, runs the constants generator, commits possible changes and creates a pull request. It will update the same pull request if subsequent runs find different changes and it won't create a pull request if no changes were detected (GitHub-Action: [peter-evans/create-pull-request@v3](https://github.com/peter-evans/create-pull-request)).

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
    _ = multicodec.CIDV2
}
```

## Generator

To generate the constants yourself checkout the repository like

```shell
git clone --recursive https://github.com/dennis-tra/go-multicodec.git
```

The `--recursive` flag makes sure you're also fetching the submodule. Then run `go generate` in the root of the repository. Currently this will invoke (see [`init.go`](./init.go)):

```shell
go run ./gen/
go fmt ./...
```

## Maintainers

[@dennis-tra](https://github.com/dennis-tra).

## License

[MIT](LICENSE) © Dennis Trautwein