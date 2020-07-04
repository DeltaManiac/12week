package lru_cache

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	capacity := 10
	cache := Constructor(10)
	if cache.cap != capacity {
		t.Errorf("Cache capacity error, got: %v want: %v ", cache.cap, capacity)
	}
	if cache.cache == nil {
		t.Errorf("Cache map error, got: %v want: not nil ", cache.cache)
	}
}

func TestPut(t *testing.T) {
	c := Constructor(3)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	if c.cap != 3 {
		t.Errorf("Cache capacity error, got: %v want: %v ", c.cap, 3)
	}
	if c.cache[1].data != 1 {
		t.Errorf("Cache Item error, got: %v want: %v ", c.cache[1].data, 1)
	}
	if c.cache[2].data != 2 {
		t.Errorf("Cache Item error, got: %v want: %v ", c.cache[2].data, 2)
	}
	if c.cache[3].data != 3 {
		t.Errorf("Cache Item error, got: %v want: %v ", c.cache[3].data, 3)
	}
	c.Put(4, 4)
	if c.cap != 3 {
		t.Errorf("Cache capacity error, got: %v want: %v ", c.cap, 3)
	}
	if c.cache[4].data != 4 {
		t.Errorf("Cache Item error, got: %v want: %v ", c.cache[3].data, 4)
	}
	if c.cache[1] != nil {
		t.Errorf("Cache Item error, got: %v want: %v ", c.cache[1], nil)
	}
	c.Put(5, 5)
	if c.cap != 3 {
		t.Errorf("Cache capacity error, got: %v want: %v ", c.cap, 3)
	}
	if c.cache[5].data != 5 {
		t.Errorf("Cache Item error, got: %v want: %v ", c.cache[3].data, 5)
	}
	if c.cache[2] != nil {
		t.Errorf("Cache Item error, got: %v want: %v ", c.cache[2], nil)
	}
	c.Put(4, 9)
	if c.cap != 3 {
		t.Errorf("Cache capacity error, got: %v want: %v ", c.cap, 3)
	}
	if c.cache[4].data != 9 {
		t.Errorf("Cache Item error, got: %v want: %v ", c.cache[4].data, 9)
	}
}
func TestGet(t *testing.T) {
	c := Constructor(3)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	x := c.Get(3)
	if x != 3 {
		t.Errorf("Cache Item retrieval error, got: %v want: %v ", c.Get(3), 3)
	}
	x = c.Get(1)
	if x != 1 {
		t.Errorf("Cache Item retrieval error, got: %v want: %v ", c.Get(1), 1)
	}
	c.Put(4, 4)
	x = c.Get(2)
	if x != -1 {
		t.Errorf("Cache Item retrieval error, got: %v want: %v ", c.Get(2), -1)
	}
	x = c.Get(1)
	if x != 1 {
		t.Errorf("Cache Item retrieval error, got: %v want: %v ", c.Get(1), 1)
	}
}
