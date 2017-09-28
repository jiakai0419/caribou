package sliceutil

import (
	"testing"
	"reflect"
)

func TestSplitN(t *testing.T) {
	actual := SplitN([]int64{1024, 2048, 3, 7, 9, 22, 88, 55, 13, 17}, 4)
	expected := [][]int64 {
		{1024, 2048, 3, 7},
		{9, 22, 88, 55},
		{13, 17},
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected:%v, actual:%v", expected, actual)
	}
}
