package timeutil

import "fmt"
import "time"

func FormatConciseDate(t time.Time) string {
	return fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
}
