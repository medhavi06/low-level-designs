package LFU

type Node struct {
	key       int
	value     int
	frequency int
	prev      *Node
	next      *Node
}

type Cache struct {
	store map[int]*Node
}

func NewLFUCache(store map[int]*Node) *Cache {
	return &Cache{store: store}
}

func (lru *Cache) Get(key int) (int, error) {
	return -1, nil
}

func (lru *Cache) Set(key, value int) error {
	//return success or failure along with error
	return nil
}
