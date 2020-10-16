# go-multicodec
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)


Periodically Self-Generated Go constants of the [compact codecs](https://github.com/multiformats/multicodec) used by many [multiformats](https://github.com/multiformats/multiformats).

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Motivation](#motivation)
- [Install](#install)
- [Usage](#usage)
- [Generator](#generator)
- [Related Efforts](#related-efforts)
- [Maintainers](#maintainers)
- [License](#license)

## Motivation

Diverging canonical table and implementation.
Usage of GitHub action

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
    multicodec.Cidv2
}
```

## Generator


## Related Efforts


## Maintainers

[@dennis-tra](https://github.com/dennis-tra).

## License

[MIT](LICENSE) © Dennis Trautwein