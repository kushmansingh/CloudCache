package cloudcache

import (
	"errors"
)

var (
	ErrNotFound = errors.New("cloudcache: item not found")
)

type Cacher interface {
	Get(string) (interface{}, error)
	Put(string, interface{}) error
}

type MemoryCacher struct {
	Memory map[string]interface{}
}

func MemoryCache() Cacher {
	return &MemoryCacher{
		Memory: map[string]interface{}{},
	}
}

func (c *MemoryCacher) Get(key string) (interface{}, error) {
	val, ok := c.Memory[key]
	if !ok {
		return nil, ErrNotFound
	}
	return val, nil
}

func (c *MemoryCacher) Put(key string, value interface{}) error {
	c.Memory[key] = value
	return nil
}
