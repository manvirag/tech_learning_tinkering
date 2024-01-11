package controller

import (
	"in_memory_key_store/domain"
	"in_memory_key_store/repository"
)

type ICacheController interface {
	Put(key string, value interface{}) error
	Get(key string) (error, interface{})
	Remove(key string) error
}

type CacheController struct {
	cache              domain.Cache
	evictionPolicyRepo repository.EvictionPolicyRepo
}

func (cc *CacheController) Put(key string, value interface{}) error {
	return cc.evictionPolicyRepo.Put(cc.cache, key, value)
}

func (cc *CacheController) Get(key string) (error, interface{}) {
	return cc.evictionPolicyRepo.Get(key)
}

func (cc *CacheController) Remove(key string) error {
	return cc.evictionPolicyRepo.Remove(key)
}
