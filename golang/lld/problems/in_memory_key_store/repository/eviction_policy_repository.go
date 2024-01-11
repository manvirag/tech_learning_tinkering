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
	keyMap   map[string]*dll.DoublyLinkedList
	rootNode *dll.DoublyLinkedList
	endNode  *dll.DoublyLinkedList
}

func (luc *LruEvictionPolicyRepo) Get(key string) (error, interface{}) {
	if node, ok := luc.keyMap[key]; ok {
		luc.moveToFront(node)
		return nil, node.Value
	} else {
		return errors.New("not found key in cache"), nil
	}
}

func (luc *LruEvictionPolicyRepo) Put(cache domain.Cache, key, value string) error {
	if node, exists := luc.keyMap[key]; exists {
		node.Value = value
		luc.moveToFront(node)
	} else {
		newNode := &dll.DoublyLinkedList{Key: key, Value: value, Previous: nil, Next: nil}
		if len(luc.keyMap) == cache.Size {
			luc.evictLRUNode()
		}
		luc.addToFront(&cache, newNode)
		luc.keyMap[key] = newNode
	}
	return nil
}

func (luc *LruEvictionPolicyRepo) moveToFront(node *dll.DoublyLinkedList) {
	if node == luc.rootNode {
		return
	}

	node.Previous.Next = node.Next
	if node.Next != nil {
		node.Next.Previous = node.Previous
	} else {
		luc.endNode = node.Previous
	}

	node.Previous = nil
	node.Next = luc.rootNode
	luc.rootNode.Previous = node
	luc.rootNode = node
}

func (luc *LruEvictionPolicyRepo) addToFront(cache *domain.Cache, node *dll.DoublyLinkedList) {
	if luc.rootNode == nil {
		luc.rootNode = node
		luc.endNode = node
	} else {
		node.Next = luc.rootNode
		luc.rootNode.Previous = node
		luc.rootNode = node
	}
	if len(luc.keyMap) > cache.Size {
		luc.evictLRUNode()
	}
}

func (luc *LruEvictionPolicyRepo) evictLRUNode() {
	if luc.endNode == nil {
		return
	}

	delete(luc.keyMap, luc.endNode.Key)

	if luc.endNode.Previous != nil {
		luc.endNode.Previous.Next = nil
		luc.endNode = luc.endNode.Previous
	} else {
		luc.rootNode = nil
		luc.endNode = nil
	}
}

func (luc *LruEvictionPolicyRepo) Remove(key string) error {
	if valuePointer, ok := luc.keyMap[key]; ok {
		if valuePointer.Previous != nil {
			valuePointer.Previous.Next = valuePointer.Next
			valuePointer.Next.Previous = valuePointer.Previous
		} else {
			luc.rootNode = luc.rootNode.Next
		}
		delete(luc.keyMap, key)
		return nil
	} else {
		return errors.New("not found key in cache")
	}
}
