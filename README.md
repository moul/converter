# converter
:scissors: multiformat data conversion

[![Build Status](https://travis-ci.org/moul/converter.svg?branch=master)](https://travis-ci.org/moul/converter)
[![GoDoc](https://godoc.org/github.com/moul/converter?status.svg)](https://godoc.org/github.com/moul/converter)
[![Coverage Status](https://coveralls.io/repos/moul/converter/badge.svg?branch=master&service=github)](https://coveralls.io/github/moul/converter?branch=master)

## Usage

`converter` can be used in the CLI as follow:

```console
$ converter --list-filters
Available filters:
- md5
- sha1
- base64-encode
- base64-decode
- base32-encode
- base32-decode
- hex-encode
- hex-decode
- xml-encode
- xml-decode
- json-encode
- json-decode
- toml-encode
- csv-decode
- fetch
- sleep-100ms
- sleep-1s
- sleep-2s
- sleep-5s
- sleep-10s
- sleep-1m
- reverse
- upper
- lower
- split-lines
- to-unix
- parse-ansi-date
- parse-rfc339-date
- parse-rfc822-date
- parse-rfc850-date
- parse-rfc1123-date
- parse-unix-date
- bytes-to-string
- string-to-bytes
- int-to-string
- string-to-int
- string-to-float
- float-to-string
```

```console
$ echo 'Hello World!' | converter reverse
!dlroW olleH
$ echo 'Hello World!' | converter md5
8ddd8be4b179a529afa5f2ffae4b9858
$ echo 'Hello World!' | converter md5 md5
b87408ae303f7ca8d4834e5ac3143d06
$ echo 'Hello World!' | converter md5 md5 md5
710f24df02eb8e151074364ea23e1a39
$ echo 'Hello World!' | converter reverse md5 upper reverse
26E80BC257BC2EB49316825A8DB8E0C9
$ echo 'Hello World!' | converter reverse md5 upper reverse base64-decode
[219 161 60 208 16 182 231 176 66 216 64 120 247 125 122 243 110 64 240 48 124 19 64 189]
$ echo 'Hello World!' | converter reverse md5 upper reverse base64-decode bytes-to-string
ۡ<���B�@x�}z�n@�0|@�
```

## Using as a Golang library

See [GoDoc](https://godoc.org/github.com/moul/converter) for usage and examples.

## Changelog

### v1.0.0 (2015-12-02)

* First release
* Chaining support
* Pipe support
* Basic streaming support
* Helper to convert a basic converter to streaming converter
* Automtically detect intermediary converters if needed
* Available filters:
  * base32-decode
  * base32-encode
  * base64-decode
  * base64-encode
  * bytes-to-string
  * csv-decode
  * fetch
  * float-to-string
  * hex-decode
  * hex-encode
  * int-to-string
  * json-decode
  * json-encode
  * lower
  * md5
  * parse-ansi-date
  * parse-date
  * parse-rfc1123-date
  * parse-rfc339-date
  * parse-rfc822-date
  * parse-rfc850-date
  * parse-unix-date
  * parse-unix-timestamp
  * reverse
  * sha1
  * sleep-100ms
  * sleep-10s
  * sleep-1m
  * sleep-1s
  * sleep-2s
  * sleep-5s
  * split-lines
  * string-to-bytes
  * string-to-float
  * string-to-int
  * time-to-string
  * to-unix
  * toml-encode
  * upper
  * xml-decode
  * xml-encode

## License

MIT
