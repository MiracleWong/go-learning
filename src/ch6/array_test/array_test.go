package operator_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1, 2, 3}
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[0], arr[1], arr[2])
	t.Log(arr1[1])
	t.Log(arr2, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr4 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr4); i++ {
		t.Log(arr4[i])
	}
	for idx, e := range arr4 {
		t.Log(idx, e)
	}
	for _, e := range arr4 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr5 := [...]int{1, 2, 3, 4, 5}
	arr5_sec1 := arr5[0:3]
	arr5_sec2 := arr5[0:]
	arr5_sec3 := arr5[1:len(arr5)]
	arr5_sec4 := arr5[:len(arr5)]
	t.Log(arr5_sec1)
	t.Log(arr5_sec2)
	t.Log(arr5_sec3)
	t.Log(arr5_sec4)
}
