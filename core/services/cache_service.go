package services

import "github.com/hayrullahcansu/cachy/core/caching"

type CacheService struct {
	*BaseService
}

func NewCacheService() *CacheService {
	return &CacheService{
		BaseService: NewBaseService(),
	}
}

func (s *CacheService) GetCacheEntry(key string) *caching.CacheItem {
	item := caching.Instance().Get(key)
	if item != nil {
		return item
	} else {
		return nil
	}
}
func (s *CacheService) GetCacheEntries() []*caching.CacheItem {
	items := caching.Instance().List()
	var result []*caching.CacheItem = make([]*caching.CacheItem, 0, len(items))
	if items != nil {
		result = append(result, items...)
	}
	return result
}

func (s *CacheService) SetCacheEntryLimitless(key string, data interface{}) *caching.CacheItem {
	return s.setCacheEntry(key, data, 0)
}
func (s *CacheService) SetCacheEntry(key string, data interface{}, dead int64) *caching.CacheItem {
	return s.setCacheEntry(key, data, dead)
}

func (s *CacheService) DeleteCacheEntry(key string) *caching.CacheItem {
	return caching.Instance().Delete(key)
}
func (s *CacheService) FlushEntries(key string) {
	caching.Instance().Flush()
}

func (s *CacheService) setCacheEntry(key string, data interface{}, dead int64) *caching.CacheItem {
	item := caching.Instance().SetWithTimeSpan(key, data, dead)
	if item != nil {
		return item
	} else {
		return nil
	}
}
