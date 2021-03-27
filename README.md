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

AUTHOR:
   Manfred Touron <https://github.com/moul/converter>

COMMANDS:
   base32
   base32-decode
   base58
   base58-decode
   base64
   base64-decode
   csv-decode
   hex
   hex-decode
   hexbase32
   hexbase32-decode
   json
   json-decode
   lower
   md5
   rawurlbase64
   rawurlbase64-decode
   rev
   sha1
   title
   toml
   upper
   urlbase64
   urlbase64-decode
   xml
   xml-decode
   help, h              Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

[embedmd]:# (.tmp/examples.txt console)
```console
foo@bar:~$ echo -n "Hello World!" | converter _bytes-to-string rev
!dlroW olleH
foo@bar:~$ echo "Hello World!" | converter md5
[141 221 139 228 177 121 165 41 175 165 242 255 174 75 152 88]
foo@bar:~$ echo "Hello World!" | converter md5 md5
[98 213 234 111 247 90 250 37 61 11 160 58 20 171 41 82]
foo@bar:~$ echo "Hello World!" | converter md5 md5 md5
[133 150 171 125 251 139 53 229 243 216 47 103 80 243 191 9]
foo@bar:~$ echo "Hello World!" | converter _bytes-to-string rev _string-to-bytes md5 hex upper rev
26E80BC257BC2EB49316825A8DB8E0C9
foo@bar:~$ echo "Hello World!" | converter _bytes-to-string rev _string-to-bytes md5 hex upper rev base64-decode
[219 161 60 208 16 182 231 176 66 216 64 120 247 125 122 243 110 64 240 48 124 19 64 189]
foo@bar:~$ echo "Hello World!" | converter _bytes-to-string rev _string-to-bytes md5 hex upper rev base64-decode bytes-to-string
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
