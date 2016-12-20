package fastLRU_test

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"testing"

	"github.com/mushroomsir/fastLRU"
	"github.com/stretchr/testify/assert"
)

func TestFastLRU(t *testing.T) {
	t.Run("TestFastLRU with default Options should be", func(t *testing.T) {
		assert := assert.New(t)
		lru := fastLRU.New(fastLRU.Options{})
		for i := 0; i < 1000; i++ {
			lru.Add(strconv.Itoa(i), i)
		}
		assert.Equal(1000, lru.Count())

	})
	t.Run("TestFastLRU with Get should be", func(t *testing.T) {
		assert := assert.New(t)
		lru := fastLRU.New(fastLRU.Options{})
		for i := 0; i < 1000; i++ {
			lru.Add(strconv.Itoa(i), i)
		}
		for i := 0; i < 1000; i++ {
			val, _ := lru.Get(strconv.Itoa(i))
			assert.Equal(i, val.(int))
		}
	})
	t.Run("TestFastLRU with adjust should be", func(t *testing.T) {
		assert := assert.New(t)
		lru := fastLRU.New(fastLRU.Options{MaxSize: 100})
		for i := 0; i < 101; i++ {
			lru.Add(strconv.Itoa(i), i)
		}

		val, _ := lru.Get("100")
		assert.Equal(100, val.(int))
	})

	t.Run("TestFastLRU with empty key should be", func(t *testing.T) {
		assert := assert.New(t)
		lru := fastLRU.New(fastLRU.Options{MaxSize: 100})
		id := genID()
		val, _ := lru.Get(id)

		assert.Equal(nil, val)
	})
	t.Run("TestFastLRU with remove key should be", func(t *testing.T) {
		assert := assert.New(t)
		lru := fastLRU.New(fastLRU.Options{MaxSize: 100})
		id := genID()
		lru.Add(id, 1)
		val, _ := lru.Get(id)
		assert.Equal(1, val.(int))
		lru.Remove(id)
		val, _ = lru.Get(id)
		assert.Equal(nil, val)
	})
}
func genID() string {
	buf := make([]byte, 12)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(buf)
}
