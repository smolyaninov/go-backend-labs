package convert

import (
	"fmt"
	"math"
)

var weightToGram = map[string]float64{
	"mg": 0.001,
	"g":  1,
	"kg": 1000,

	"oz": 28.349523125,
	"lb": 453.59237,
}

func Weight(value float64, from, to string) (float64, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, ErrInvalidValue
	}
	if value < 0 {
		return 0, ErrNegativeValue
	}

	fromFactor, ok := weightToGram[from]
	if !ok {
		return 0, fmt.Errorf("%w: %s", ErrUnknownUnit, from)
	}
	toFactor, ok := weightToGram[to]
	if !ok {
		return 0, fmt.Errorf("%w: %s", ErrUnknownUnit, to)
	}

	grams := value * fromFactor
	out := grams / toFactor
	return out, nil
}
