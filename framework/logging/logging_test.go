package logging

import "testing"

func TestLogger(t *testing.T) {
	if Instance() == nil {
		t.Error("Logger instance can't crate")
	}
}
