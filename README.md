# avl
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hslam/avl)](https://pkg.go.dev/github.com/hslam/avl)
[![Build Status](https://travis-ci.org/hslam/avl.svg?branch=master)](https://travis-ci.org/hslam/avl)
[![codecov](https://codecov.io/gh/hslam/avl/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/avl)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/avl)](https://goreportcard.com/report/github.com/hslam/avl)
[![LICENSE](https://img.shields.io/github/license/hslam/avl.svg?style=flat-square)](https://github.com/hslam/avl/blob/master/LICENSE)

Package avl implements an AVL tree.

## Get started

### Install
```
go get github.com/hslam/avl
```
### Import
```
import "github.com/hslam/avl"
```
### Usage
#### Example
```go
package main

import (
	"fmt"
	"github.com/hslam/avl"
)

func main() {
	t := avl.New()
	str := String("Hello World")
	t.Insert(str)
	fmt.Println(t.Search(str))
	t.Delete(str)
}

type String string

func (a String) Less(b avl.Item) bool {
	return a < b.(String)
}
```

#### Output
```
Hello World
```

#### Iterator Example
```go
package main

import (
	"fmt"
	"github.com/hslam/avl"
)

func main() {
	t := avl.New()
	l := "MNOLKQPHIA"
	for _, v := range l {
		t.Insert(String(v))
	}
	iter := t.Min()
	for iter != nil {
		fmt.Printf("%s\t", iter.Item())
		iter = iter.Next()
	}
}

type String string

func (a String) Less(b avl.Item) bool {
	return a < b.(String)
}
```
#### AVL Tree
<img src="https://raw.githubusercontent.com/hslam/avl/master/avl.png" alt="avl" align=center>

#### Output
```
A	H	I	K	L	M	N	O	P	Q
```

### License
This package is licensed under a MIT license (Copyright (c) 2020 Meng Huang)

### Author
avl was written by Meng Huang.


