# avl
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hslam/avl)](https://pkg.go.dev/github.com/hslam/avl)
[![Build Status](https://github.com/hslam/avl/workflows/build/badge.svg)](https://github.com/hslam/avl/actions)
[![codecov](https://codecov.io/gh/hslam/avl/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/avl)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/avl)](https://goreportcard.com/report/github.com/hslam/avl)
[![LICENSE](https://img.shields.io/github/license/hslam/avl.svg?style=flat-square)](https://github.com/hslam/avl/blob/master/LICENSE)

Package avl implements an AVL tree.

## [Benchmark](http://github.com/hslam/tree-benchmark "tree-benchmark")

<img src="https://raw.githubusercontent.com/hslam/tree-benchmark/master/tree-insert.png" width = "400" height = "300" alt="insert" align=center><img src="https://raw.githubusercontent.com/hslam/tree-benchmark/master/tree-delete.png" width = "400" height = "300" alt="delete" align=center>

<img src="https://raw.githubusercontent.com/hslam/tree-benchmark/master/tree-search.png" width = "400" height = "300" alt="search" align=center><img src="https://raw.githubusercontent.com/hslam/tree-benchmark/master/tree-iterate.png" width = "400" height = "300" alt="iterate" align=center>


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
	tree := avl.New()
	str := String("Hello World")
	tree.Insert(str)
	fmt.Println(tree.Search(str))
	tree.Delete(str)
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
	tree := avl.New()
	l := "MNOLKQPHIA"
	for _, v := range l {
		tree.Insert(String(v))
	}
	iter := tree.Min()
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


