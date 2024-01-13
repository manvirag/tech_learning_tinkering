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
	Cache              domain.Cache
	EvictionPolicyRepo repository.EvictionPolicyRepo
}

func (cc *CacheController) Put(key string, value interface{}) error {
	return cc.EvictionPolicyRepo.Put(cc.Cache, key, value)
}

func (cc *CacheController) Get(key string) (error, interface{}) {
	return cc.EvictionPolicyRepo.Get(key)
}

func (cc *CacheController) Remove(key string) error {
	return cc.EvictionPolicyRepo.Remove(key)
}
