package caching

type CacheItem struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value,omitempty"`
	Dead  int64       `json:"dead"`
}

func NewCacheItem(key string, value interface{}, dead int64) *CacheItem {
	return &CacheItem{
		Key:   key,
		Value: value,
		Dead:  dead,
	}
}
