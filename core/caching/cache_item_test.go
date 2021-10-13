package caching

import (
	"errors"
	"testing"
)

func TestNewCacheItem(t *testing.T) {
	tests := map[string]struct {
		inputKey   string
		inputValue interface{}
		inputDead  int64
		output     CacheItem
		err        error
	}{
		"successful conversion 1": {
			inputKey:   "test_1",
			inputValue: 10,
			inputDead:  123423425,
			output: CacheItem{
				Key:   "test_1",
				Value: 10,
				Dead:  123423425,
			},
			err: nil,
		},
		"successful conversion 2": {
			inputKey:   "test_1",
			inputValue: 10,
			inputDead:  123423425,
			output: CacheItem{
				Key:   "test_1",
				Value: 10,
				Dead:  123423425,
			},
			err: errors.New("exception"),
		},
	}
	for _, test := range tests {
		a, b := test.output, NewCacheItem(test.inputKey, test.inputValue, test.inputDead)
		if a.Key != b.Key {
			t.Errorf("New Cache Item not works expected")
		}
		if a.Value != b.Value {
			t.Errorf("New Cache Item not works expected")
		}
		if a.Dead != b.Dead {
			t.Errorf("New Cache Item not works expected")
		}
	}
}
