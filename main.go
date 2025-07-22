package main

import "cache/evictionPolicy/LRU"

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	cache := NewCache()
	lru := LRU.NewLRUCache(10)
	cache.SetEvictionPolicy(lru, 10)
	_ = cache.setKey(2, 2)
	value, _ := cache.getKey(2)
	println("Value : ", value)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
