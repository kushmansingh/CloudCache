package cloudcache

import (
	"testing"
)

func TestMemoryCacheGet(t *testing.T) {
	key := "key"

	cases := []struct {
		Description string
		Cache       Cacher
		Error       error
	}{
		{
			Description: "Not found.",
			Cache:       MemoryCache(),
			Error:       ErrNotFound,
		},
		{
			Description: "Found.",
			Cache: &MemoryCacher{
				Memory: map[string]interface{} {
					key: "value",
				},
			},
			Error: nil,
		},
	}
	for _, c := range cases {
		t.Run(c.Description, func(*testing.T) {
			_, err := c.Cache.Get(key)
			if err != nil {
				if err != c.Error {
					t.Errorf("Expected %s but got %s", c.Error, err)
				}
			}
		})
	}
}
