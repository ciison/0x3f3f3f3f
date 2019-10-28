package test

import (
	"reflect"
	"testing"
)

func TestList1UnionList2(t *testing.T) {
	l1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 21}
	l2 := []int{7, 8, 9, 10, 11, 12, 13}
	ans := List1UnionList2(l1, l2)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 21}

	if reflect.DeepEqual(ans, expected) {
		t.Log("Accepted")
	} else {
		t.Errorf("expected %+v  but get %+v\n",expected, ans )
	}

}
