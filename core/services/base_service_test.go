package services

import "testing"

func TestNewBaseService(t *testing.T) {
	if NewBaseService() == nil {
		t.Error("BaseService Instance can't create")
	}
}
