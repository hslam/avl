# avl
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
	fmt.Println(t.Search(str).Item())
	t.Delete(str)
}

type String string

func (a String) Less(than avl.Item) bool {
	b, _ := than.(String)
	return a < b
}
```

#### Output
```
Hello World
```

### License
This package is licensed under a MIT license (Copyright (c) 2020 Meng Huang)

### Author
avl was written by Meng Huang.


