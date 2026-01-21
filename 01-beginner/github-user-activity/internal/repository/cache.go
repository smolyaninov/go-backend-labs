package repository

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type JSONCache[T any] struct {
	Dir string
	TTL time.Duration
}

func NewJSONCache[T any](dir string, ttl time.Duration) *JSONCache[T] {
	return &JSONCache[T]{Dir: dir, TTL: ttl}
}

func (c *JSONCache[T]) Load(key string) ([]T, error) {
	path := c.cacheFilePath(key)

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if time.Since(info.ModTime()) > c.TTL {
		return nil, errors.New("cache expired")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var items []T
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (c *JSONCache[T]) Save(key string, items []T) error {
	if err := os.MkdirAll(c.Dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(c.cacheFilePath(key), data, 0644)
}

func (c *JSONCache[T]) cacheFilePath(key string) string {
	return filepath.Join(c.Dir, key+".json")
}
