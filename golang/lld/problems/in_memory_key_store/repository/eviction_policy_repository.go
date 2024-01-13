package repository

import (
	"errors"
	"in_memory_key_store/dll"
	"in_memory_key_store/domain"
)

type EvictionPolicyRepo interface {
	Get(key string) (error, interface{})
	Put(cache domain.Cache, key string, value interface{}) error
	Remove(key string) error
}

type LruEvictionPolicyRepo struct {
	KeyMap   map[string]*dll.DoublyLinkedList
	RootNode *dll.DoublyLinkedList
	EndNode  *dll.DoublyLinkedList
}

func (luc *LruEvictionPolicyRepo) Get(key string) (error, interface{}) {
	if node, ok := luc.KeyMap[key]; ok {
		luc.moveToFront(node)
		return nil, node.Value
	} else {
		return errors.New("not found key in cache"), nil
	}
}

func (luc *LruEvictionPolicyRepo) Put(cache domain.Cache, key string, value interface{}) error {
	if node, exists := luc.KeyMap[key]; exists {
		node.Value = value
		luc.moveToFront(node)
	} else {
		newNode := &dll.DoublyLinkedList{Key: key, Value: value, Previous: nil, Next: nil}
		if len(luc.KeyMap) == cache.Size {
			luc.evictLRUNode()
		}
		luc.addToFront(&cache, newNode)
		luc.KeyMap[key] = newNode
	}
	return nil
}

func (luc *LruEvictionPolicyRepo) moveToFront(node *dll.DoublyLinkedList) {
	if node == luc.RootNode {
		return
	}

	node.Previous.Next = node.Next
	if node.Next != nil {
		node.Next.Previous = node.Previous
	} else {
		luc.EndNode = node.Previous
	}

	node.Previous = nil
	node.Next = luc.RootNode
	luc.RootNode.Previous = node
	luc.RootNode = node
}

func (luc *LruEvictionPolicyRepo) addToFront(cache *domain.Cache, node *dll.DoublyLinkedList) {
	if luc.RootNode == nil {
		luc.RootNode = node
		luc.EndNode = node
	} else {
		node.Next = luc.RootNode
		luc.RootNode.Previous = node
		luc.RootNode = node
	}
	if len(luc.KeyMap) > cache.Size {
		luc.evictLRUNode()
	}
}

func (luc *LruEvictionPolicyRepo) evictLRUNode() {
	if luc.EndNode == nil {
		return
	}

	delete(luc.KeyMap, luc.EndNode.Key)

	if luc.EndNode.Previous != nil {
		luc.EndNode.Previous.Next = nil
		luc.EndNode = luc.EndNode.Previous
	} else {
		luc.RootNode = nil
		luc.EndNode = nil
	}
}

func (luc *LruEvictionPolicyRepo) Remove(key string) error {
	if valuePointer, ok := luc.KeyMap[key]; ok {
		if valuePointer.Previous != nil {
			valuePointer.Previous.Next = valuePointer.Next
			valuePointer.Next.Previous = valuePointer.Previous
		} else {
			luc.RootNode = luc.RootNode.Next
		}
		delete(luc.KeyMap, key)
		return nil
	} else {
		return errors.New("not found key in cache")
	}
}
