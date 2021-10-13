package caching

import (
	"testing"
	"time"
)

func TestInstance(t *testing.T) {
	if Instance() == nil {
		t.Error("Cache Instance can't create")
	}
	testInstance := getTestInstanceWithSomeData()
	testKey1 := testInstance.Get("test_key1")
	testKey2 := testInstance.Get("test_key2")
	if testKey1 == nil {
		t.Errorf("test_key1 object cannot found in cache. It must be exists")
	}
	if testKey1 != nil && testKey1.Value != "test_key1_value" {
		t.Errorf("test_key1 object value is not equals")
	}
	if testKey2.Dead > 0 {
		t.Errorf("test_key2 object dead value cannot bigger than 0")
	}
}

func TestExpireFeature(t *testing.T) {
	testInstance := getTestInstanceWithSomeData()
	testKey3 := testInstance.Get("test_key3")
	if testKey3 == nil {
		t.Errorf("testKey3 object cannot found in cache. It must be exists")
	}
	time.Sleep(time.Duration(5) * time.Second)
	testKey3 = testInstance.Get("test_key3")
	if testKey3 != nil {
		t.Errorf("testKey3 object found in cache. It must not be exists")
	}
}

func getTestInstanceWithSomeData() *Cache {
	testInstance := Instance()
	testInstance.SetWithTimeSpan("test_key1", "test_key1_value", time.Now().Add(time.Duration(300)*time.Second).UnixNano())
	testInstance.Set("test_key2", 100)
	testInstance.SetWithTimeSpan("test_key3", "test_key3_value", time.Now().Add(time.Duration(2)*time.Second).UnixNano())
	return testInstance
}
