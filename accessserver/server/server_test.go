package server

import (
	"math"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var toAdd uint32 = math.MaxUint32

	atomic.AddUint32(&toAdd, 1)
	t.Log(toAdd)
	if toAdd != 0 {
		t.Error("难道不因为是0吗?")
	}
}
