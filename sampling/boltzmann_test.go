package sampling

import (
	. "fmt"
	"testing"
)

func TestBoltzmannSampling1(t *testing.T) {
	counter := make([]int, 5)
	for i := 0; i < 1000000; i++ {
		counter[BoltzmannSampling(0.3, 0.4, 0.2, 0.05, 0.05)]++
	}
	expected := []int{218932, 241958, 198098, 170505, 170505}
	for i, cnt := range counter {
		Println("i:", i, "expected:", expected[i], "actual:", cnt, "a/e:", float64(cnt) / float64(expected[i]))
	}
}

func TestBoltzmannSampling2(t *testing.T) {
	counter := make([]int, 8)
	for i := 0; i < 1000000; i++ {
		counter[BoltzmannSampling(0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1)]++
	}
	expected := []int{125000, 125000, 125000, 125000, 125000, 125000, 125000, 125000}
	for i, cnt := range counter {
		Println("i:", i, "expected:", expected[i], "actual:", cnt, "a/e:", float64(cnt) / float64(expected[i]))
	}
}

func TestBoltzmannSampling3(t *testing.T) {
	counter := make([]int, 3)
	for i := 0; i < 1000000; i++ {
		counter[BoltzmannSampling(1.0, 0.1, 0.01)]++
	}
	expected := []int{562383, 228648, 208968}
	for i, cnt := range counter {
		Println("i:", i, "expected:", expected[i], "actual:", cnt, "a/e:", float64(cnt) / float64(expected[i]))
	}
}
