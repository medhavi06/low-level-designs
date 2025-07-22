package main

import "errors"

type Cache struct {
	evictionPolicy IEvictionPolicy
	capacity       int
}

func (cache *Cache) getKey(key int) (int, error) {
	value, err := cache.evictionPolicy.Get(key)
	if err != nil {
		return -1, errors.New("something went wrong")
	}
	return value, nil
}

func (cache *Cache) setKey(key, value int) error {
	err := cache.evictionPolicy.Set(key, value)
	if err != nil {
		return errors.New("things went wrong, please retry")
	}
	return nil
}

func NewCache() *Cache {
	return &Cache{}
}

func (cache *Cache) SetEvictionPolicy(policy IEvictionPolicy, capacity int) {
	cache.evictionPolicy = policy
	cache.capacity = capacity
}
