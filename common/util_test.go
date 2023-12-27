package common

import (
	"reflect"
	"testing"
)

//TestCheckInNumberSlice
func TestCheckInNumberSlice(t *testing.T) {
	arr := []uint64{1, 2, 3, 5, 6}
	var target uint64 = 3
	expect := true

	res := CheckInNumberSlice(target, arr)

	if expect != res {
		t.Errorf("expect: %v, res: %v", expect, res)
	}
}


//TestDelEleInSlice
func TestDelEleInSlice(t *testing.T) {
	arr := []uint64{1, 2, 3, 4, 5, 6}
	var target uint64 = 3
	expect := []uint64{1, 2, 4, 5, 6}

	res := DelEleInSlice(target, arr)

	if reflect.DeepEqual(expect, res) == false {
		t.Errorf("expect: %#v, res: %#v", expect, res)
	}
}