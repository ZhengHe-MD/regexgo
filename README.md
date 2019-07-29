# Regexgo [![CircleCI](https://circleci.com/gh/ZhengHe-MD/regexgo.svg?style=svg)](https://circleci.com/gh/ZhengHe-MD/regexgo)

A simple regular expression engine implemented in Go. It supports concatenation, union (|) and closure (*) operations as well as grouping. It follows Ken Thompson's algorithm for constructing an NFA from a regular expression.

This is a rewrite of [regexjs](https://github.com/deniskyashif/regexjs). Check out the original author's beautiful [blog](https://deniskyashif.com/implementing-a-regular-expression-engine/) for more info.

## Install

```sh
$ go get github.com/ZhengHe-MD/regexgo
```

## Usage

```go
package main

import (
	"fmt"
	. "github.com/ZhengHe-MD/regexgo"
)

func main() {
	r := Compile("(a|b)*c")
	cases := []string{"ac", "abc", "aabababbc", "aaab"}
	for _, c := range cases {
		fmt.Printf("match %s %v\n", c, MatchString(r, c, &MatchOptions{DFS}))
	}
}
```

## References

* [Implementing a Regular Expression Engine](https://deniskyashif.com/implementing-a-regular-expression-engine/)
* [Udacity: CS212 Week3](https://sites.google.com/site/udacitymirrorcs212/syllabus)





