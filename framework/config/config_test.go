package config

import "testing"

func TestInstance(t *testing.T) {
	if Instance() == nil {
		t.Error("Config instance can't crate")
	}
}
