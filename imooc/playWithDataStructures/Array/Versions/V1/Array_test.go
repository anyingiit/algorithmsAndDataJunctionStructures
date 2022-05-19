package V1

import (
	"testing"
)

func TestArray_resize(t *testing.T) {
	const defaultArrayCap int = 10
	array := NewArray()
	if array.cap != defaultArrayCap {
		t.Errorf("check default array cap failed")
	}

	const checkCustomCap = 233
	arrayCustomCap := NewArrayCustomCap(checkCustomCap)
	if actually := arrayCustomCap.cap; actually != checkCustomCap {
		t.Errorf("check array custom cap failed: expect %d actually %d", checkCustomCap, actually)
	}

	array1 := NewArray()
	arrayExpectLen := 20
	for i := 0; i < 11; i++ {
		array1.Append(i)
	}
	if actually := array1.cap; actually != arrayExpectLen {
		t.Errorf("check array automatcally resize cap to bigger failed: expect %d actually %d", arrayExpectLen, actually)
	}

	arrayExpectLen = 10
	for i := 0; i < 7; i++ {
		_, err := array1.RemoveFirst(i)
		if err != nil {
			t.Errorf("check array automatcally resize cap to smailler failed: excute RemoveFirst failed, index %d", i)
		}
	}
	if actually := array1.cap; actually != arrayExpectLen {
		t.Errorf("check array automatcally resize cap to smailler failed: expect %d actually %d", arrayExpectLen, actually)
	}

}
