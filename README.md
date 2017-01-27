# go-mersenne-twister
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

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
