package sampling

import (
	"math"
	"math/rand"
)

func BoltzmannSampling(logits ...float64) int {
	var deno = 0.0
	var es []float64
	for _, logit := range logits {
		e := math.Exp(logit)
		deno += e
		es = append(es, e)
	}

	var probs []float64
	for _, e := range es {
		probs = append(probs, e / deno)
	}

	return sampling(probs...)
}

func sampling(probs ...float64) int {
	var roll = rand.Float64()
	var acc = 0.0
	for i, prob := range probs {
		acc += prob
		if roll < acc {
			return i
		}
	}
	return len(probs) - 1
}
