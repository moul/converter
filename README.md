# converter

:scissors: multiformat data conversion

[![Build Status](https://travis-ci.org/moul/converter.svg?branch=master)](https://travis-ci.org/moul/converter)
[![GoDoc](https://godoc.org/github.com/moul/converter?status.svg)](https://godoc.org/github.com/moul/converter)
[![Coverage Status](https://coveralls.io/repos/moul/converter/badge.svg?branch=master&service=github)](https://coveralls.io/github/moul/converter?branch=master)

![dictionary](https://raw.githubusercontent.com/moul/converter/master/assets/dictionary.png)

## Usage

`converter` can be used in the CLI as follow:

[embedmd]:# (.tmp/usage.txt console)
```console
foo@bar:~$ converter -h
NAME:
   converter - A new cli application

USAGE:
   converter [global options] command [command options] [arguments...]

VERSION:
    ()

AUTHOR:
   Manfred Touron <https://github.com/moul/converter>

COMMANDS:
   md5                   []byte  ->  string
   sha1                  []byte  ->  string
   base58-encode         []byte  ->  string
   base58-decode         string  ->  []byte
   base64-encode         []byte  ->  string
   base64-decode         string  ->  []byte
   urlbase64-encode      []byte  ->  string
   urlbase64-decode      string  ->  []byte
   rawurlbase64-encode   []byte  ->  string
   rawurlbase64-decode   string  ->  []byte
   base32-encode         []byte  ->  string
   base32-decode         string  ->  []byte
   hex-encode            []byte  ->  string
   hex-decode            string  ->  []byte
   xml-encode            interface{}  ->  []byte
   xml-decode            []byte  ->  interface{}
   json-encode           interface{}  ->  []byte
   json-decode           []byte  ->  interface{}
   toml-encode           []byte  ->  interface{}
   csv-decode            string  ->  [][]string
   fetch                 string  ->  []byte
   sleep-100ms           interface{}  ->  interface{}
   sleep-1s              interface{}  ->  interface{}
   sleep-2s              interface{}  ->  interface{}
   sleep-5s              interface{}  ->  interface{}
   sleep-10s             interface{}  ->  interface{}
   sleep-1m              interface{}  ->  interface{}
   reverse               string  ->  string
   upper                 string  ->  string
   lower                 string  ->  string
   split-lines           []byte  ->  []byte
   to-unix               time.Time  ->  int64
   parse-ansi-date       string  ->  time.Time
   parse-rfc339-date     string  ->  time.Time
   parse-rfc822-date     string  ->  time.Time
   parse-rfc850-date     string  ->  time.Time
   parse-rfc1123-date    string  ->  time.Time
   parse-unix-date       string  ->  time.Time
   parse-date            string  ->  time.Time
   time-to-string        time.Time  ->  string
   parse-unix-timestamp  int64  ->  time.Time
   bytes-to-string       []byte  ->  string
   string-to-bytes       string  ->  []byte
   int-to-string         int64  ->  string
   string-to-int         string  ->  int64
   string-to-float       string  ->  float64
   float-to-string       float64  ->  string
   help, h               Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

[embedmd]:# (.tmp/examples.txt console)
```console
foo@bar:~$ echo -n "Hello World!" | converter reverse
!dlroW olleH
foo@bar:~$ echo "Hello World!" | converter md5
8ddd8be4b179a529afa5f2ffae4b9858
foo@bar:~$ echo "Hello World!" | converter md5 md5
b87408ae303f7ca8d4834e5ac3143d06
foo@bar:~$ echo "Hello World!" | converter md5 md5 md5
710f24df02eb8e151074364ea23e1a39
foo@bar:~$ echo "Hello World!" | converter reverse md5 upper reverse
26E80BC257BC2EB49316825A8DB8E0C9
foo@bar:~$ echo "Hello World!" | converter reverse md5 upper reverse base64-decode
[219 161 60 208 16 182 231 176 66 216 64 120 247 125 122 243 110 64 240 48 124 19 64 189]
foo@bar:~$ echo "Hello World!" | converter reverse md5 upper reverse base64-decode bytes-to-string
Û¡<Ð¶ç°BØ@x÷}zón@ð0|@½
```

## Using with Docker

```console
$ date | docker run --rm moul/converter md5 sha1
67a74306b06d0c01624fe0d0249a570f4d093747
```

## Using as a Golang library

See [GoDoc](https://godoc.org/github.com/moul/converter) for usage and examples.

## Credit

This project is inspired by the [transformer](https://github.com/jbenet/transformer/) by the venerable [Juan Benet](https://github.com/jbenet)

## License

MIT
