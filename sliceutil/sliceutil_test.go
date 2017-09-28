package sliceutil

import (
	"testing"
	"fmt"
)

//func TestFormatConciseDate(t *testing.T) {
//	tm, _ := time.Parse(time.RFC3339, "2017-05-29T16:49:52+00:00")
//	expected := "20170529"
//	actual := FormatConciseDate(tm)
//	if expected != actual {
//		t.Errorf("expected:%s, actual:%s", expected, actual)
//	}
//}

func TestSplitN(t *testing.T) {
	actual := SplitN([]int64{1024, 2048, 3, 7, 9, 22, 88, 55, 13, 17}, 4)
	fmt.Println(actual)
}
