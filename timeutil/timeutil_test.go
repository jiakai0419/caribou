package timeutil

import (
	"testing"
	"time"
)

func TestFormatConciseDate(t *testing.T) {
	tm, _ := time.Parse(time.RFC3339, "2017-05-29T16:49:52+00:00")
	expected := "20170529"
	actual := FormatConciseDate(tm)
	if expected != actual {
		t.Errorf("expected:%s, actual:%s", expected, actual)
	}
}
