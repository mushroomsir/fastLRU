fastLRU
=====
The fastest fastLRU

[![Build Status](http://img.shields.io/travis/mushroomsir/fastLRU.svg?style=flat-square)](https://travis-ci.org/mushroomsir/fastLRU)
[![Coverage Status](http://img.shields.io/coveralls/mushroomsir/fastLRU.svg?style=flat-square)](https://coveralls.io/r/mushroomsir/fastLRU)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/mushroomsir/fastLRU/master/LICENSE)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mushroomsir/fastLRU)
## Installation

```bash
go get github.com/mushroomsir/fastLRU
```
## Examples
Try into github.com/mushroomsir/fastLRU directory, then:
```go
go run examples/main.go
```
Complete example:
```go
package main

import "github.com/mushroomsir/fastLRU"
import "fmt"

func main() {
	lru := fastLRU.New(fastLRU.Options{})

	key := "top10news"
	lru.Add(key, "xxx")

	value, _ := lru.Get(key)
	fmt.Printf("%s\n", value)

	count := lru.Count()
	fmt.Printf("%d\n", count)

	lru.Remove(key)
}
```
## API
Package provides the fastest LRU algorithm.
```go
  import "github.com/mushroomsir/fastLRU"
```
### type LRUCache
```go
type LRUCache struct {
    //unexported fields
}
```

#### func New
New create a LRUCache with options.
```go
func New(opts Options) (lru *LRUCache)
```
#### func Add
Add one new cache item to storage
```go
func (l *LRUCache) Add(key string, value interface{})
```
#### func Get
Get one cache item by key 
```go
func (l *LRUCache) Get(key string) (val interface{}, ok bool)
```
#### func Remove
Remove cache item by key
```go
func (l *LRUCache) Remove(key string) 
```
#### func Count
Count get total count of cache items
```go
func (l *LRUCache) Count() int 
```