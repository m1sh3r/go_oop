package main

import (
	"testing"
	"unsafe"
)

func TestGoodStructAlignment(t *testing.T) {
	badSize := unsafe.Sizeof(Bad{})
	goodSize := unsafe.Sizeof(Good{})

	if goodSize > badSize {
		t.Errorf("good size = %d, bad size = %d", goodSize, badSize)
	}
}
