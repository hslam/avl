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
	t := avl.New(func(a, b interface{}) bool {
		return a.(string) < b.(string)
	})
	str := "Hello World"
	t.Insert(str)
	fmt.Println(t.Search(str).Value)
	t.Delete(str)
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


