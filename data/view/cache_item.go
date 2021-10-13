package view

import "time"

type CacheItemView struct {
	Key      string      `json:"key"`
	Value    interface{} `json:"value"`
	ExpireAt string      `json:"expire_at"`
}

func NewCacheItemView(key string, value interface{}, dead int64) *CacheItemView {

	expire_at := time.Unix(0, dead)
	var expire_at_formatted string
	if dead > 0 {
		expire_at_formatted = expire_at.Format("02-01-2006 15:04:05")
	}

	return &CacheItemView{
		Key:      key,
		Value:    value,
		ExpireAt: expire_at_formatted,
	}
}
