package main

import (
	"fmt"
	"in_memory_key_store/controller"
	"in_memory_key_store/dll"
	"in_memory_key_store/domain"
	"in_memory_key_store/repository"
)

func main() {

	cacheRepo := &repository.LruEvictionPolicyRepo{
		KeyMap:   make(map[string]*dll.DoublyLinkedList),
		RootNode: nil,
		EndNode:  nil,
	}
	cacheC := controller.CacheController{
		Cache: domain.Cache{
			Size: 1,
		},
		EvictionPolicyRepo: cacheRepo,
	}

	cacheC.Put("1", "hello")
	err, value := cacheC.Get("1")

	if err == nil {
		fmt.Println(value)
	} else {
		fmt.Println("error: %+v", err)
	}

	cacheC.Put("2", "hello0000")
	err2, value2 := cacheC.Get("1")

	if err2 == nil {
		fmt.Println(value2)
	} else {
		fmt.Println("error: %+v", err2)
	}

}
