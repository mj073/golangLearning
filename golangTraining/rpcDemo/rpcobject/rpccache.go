package rpcobject

import (
	"sync"
	"errors"
)

type (
	RPCCacheService struct{
		cache map[string]string
		request *Request
		mu *sync.RWMutex
	}
	CacheItem struct {
		Key string
		Value string
	}
	Request struct {
		Get uint32
		Put uint32
		Delete uint32
		Clear uint32
	}
)
func NewRPCCache() *RPCCacheService{
	return &RPCCacheService{
		cache: make(map[string]string),
		request: &Request{},
		mu: &sync.RWMutex{},
	}
}

func (r *RPCCacheService) Get(key string, resp *CacheItem) error{
	r.mu.RLock()
	defer r.mu.RUnlock()

	if val,ok := r.cache[key]; ok {
		*resp = CacheItem{key,val}
	}else {
		return errors.New("Not Found")
	}
	r.request.Get++
	return nil
}
func (r *RPCCacheService) Put(c *CacheItem, ack *bool) error{
	r.mu.Lock()
	defer r.mu.Unlock()


	r.cache[c.Key] = c.Value
	*ack = true

	r.request.Put++
	return nil
}
func (r *RPCCacheService) Delete(key string, ack *bool) error{
	r.mu.Lock()
	defer r.mu.Unlock()

	if _,ok := r.cache[key]; ok{
		delete(r.cache,key)
		*ack = true
	}else{
		return errors.New("Not found")
	}

	r.request.Delete++
	return nil
}
func (r *RPCCacheService) Clear(skip bool, ack *bool) error{
	r.mu.Lock()
	defer r.mu.Unlock()

	r.cache = make(map[string]string)
	*ack = true

	r.request.Clear++
	return nil
}
func (r *RPCCacheService) Stats(skip bool, req *Request) error{
	*req = *r.request
	return nil
}