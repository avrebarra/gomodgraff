![GitHub Logo](gomodgraff.png)

# gomodgraff [![GoDoc](https://godoc.org/github.com/shrotavre/gomodgraff?status.svg)](http://godoc.org/github.com/shrotavre/gomodgraff) [![Go Report Card](https://goreportcard.com/badge/shrotavre/gomodgraff)](https://goreportcard.com/report/github.com/shrotavre/gomodgraff)

Gomodgraff is a utility to draw relationship between packages inside a Go module/project as dot diagram.

## Installation
You can either build the binary yourself using `make build` or you can just
prebuilt binary by downloading the latest releases from [release page](https://github.com/shrotavre/gomodgraff/releases).

You can use these install scripts to download the latest version:

```sh
# install latest release to /usr/local/bin/
curl https://i.jpillora.com/avrebarra/gomodgraff! | *remove_this* bash
```

```sh
# install specific version
curl https://i.jpillora.com/avrebarra/gomodgraff@{version} | *remove_this* bash
```

## Usage

~~~ bash

# Show helps
$ gomodgraff -help

# Basic usages run in current directory
$ gomodgraff -only-internal
digraph sample {
"&cmd" -> "&modgraff";
"&main" -> "&cmd";
}

# Basic usages run in current directory, pipe to dot, and output a png
$ gomodgraff -only-internal | dot -Tpng -o gomodgraff.png
~~~

