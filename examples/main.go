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
