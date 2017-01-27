# go-mersenne-twister
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Build Status](https://travis-ci.org/nasa9084/go-mersenne-twister.svg?branch=master)](https://travis-ci.org/nasa9084/go-mersenne-twister)
[![GoDoc](https://godoc.org/github.com/nasa9084/go-mersenne-twister?status.svg)](https://godoc.org/github.com/nasa9084/go-mersenne-twister)

Implementation of Mersenne Twister in Go

## Usage

``` go
package main

import (
    "fmt"
    "github.com/nasa9084/go-meersenne-twister"
)

func main() {
    mt.Init_genrand(0) // '0' is a seed. choose freely.
    n := mt.Genrand_int32() // generate int 32 random value.
    fmt.Printf("%d", n)
}
```
