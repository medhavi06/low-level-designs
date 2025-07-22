package main

type IEvictionPolicy interface {
	Get(key int) (int, error)
	Set(key, value int) error
}
