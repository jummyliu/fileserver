# fileserver
A simple file server writed with go, like python `http.server`.

It will create file server in current directory by default (http://0.0.0.0:8000).

# Installation
To install this package, you need to install Go and set your Go workspace first.

1. The first need Go installed (version 1.11+ is required), then you can use the below Go command to install fileserver.

```shell
$ go get -u github.com/jummyliu/fileserver
```

# Usage

```shell
$ fileserver -h
  -addr string
        addr (default "0.0.0.0")
  -h    help
  -path string
        basic path (default ".")
  -port int
        port (default 8000)
```
