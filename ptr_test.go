package ptr_test

import (
	"testing"

	"github.com/Inuart/ptr"
)

func TestPtr(t *testing.T) {
	testRetrieve(t, nil, 0)
	testRetrieve(t, nil, 1)
	testRetrieve(t, ptr.To(1), 2)

	testRetrieve(t, nil, "")
	testRetrieve(t, nil, "a")
	testRetrieve(t, ptr.To("a"), "b")
}

func testRetrieve[T comparable](t *testing.T, p *T, sample T) {
	var pv, zero T
	if p != nil {
		pv = *p
	}

	_, ok := ptr.Val(p)
	if p == nil && ok {
		t.Error("p is nil, should return (zero, false)")
	}
	if p != nil && !ok {
		t.Errorf("p is not nil, should return (%v, true)", pv)
	}

	v := ptr.ValOr(p, sample)
	if p == nil && v != sample {
		t.Errorf("p is nil, should return %v", sample)
	}

	v = ptr.ValOrZero(p)
	if p == nil && v != zero {
		t.Error("p is nil, should return 0")
	}
	if p != nil && v == zero {
		t.Errorf("p is not nil, should return %v", pv)
	}
}
