![](https://counter.gofiber.io/badge/bitmaskit/go-portscanner?unique=true)

## USE AT OWN RISK 
## DON'T DO ANYTHING ILLEGAL

## go-portscanner
Port scanner written in Go

### Requirements
Installed Go

### Usage

```shell script
$ git clone git@github.com:bitmaskit/go-portscanner
$ cd go-portscanner
$ go build .
```

Addr - address you want to scan

From - starting port

To - ending port

W - number of workers(threads). More threads = faster scan*
```shell script
$ ./go-portscanner -addr=localhost -from=10 -to=444 -w=1000
```

*Based on PC configuration
