package services

import "testing"

func TestNewCacheService(t *testing.T) {
	if NewCacheService() == nil {
		t.Error("CacheService Instance can't create")
	}
}
