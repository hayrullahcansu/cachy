package logging

import "testing"

func TestInstance(t *testing.T) {
	if Instance() == nil {
		t.Error("Logger instance can't crate")
	}
}
