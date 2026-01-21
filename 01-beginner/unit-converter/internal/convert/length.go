package convert

import (
	"fmt"
	"math"
)

var lengthToMeter = map[string]float64{
	"mm": 0.001,
	"cm": 0.01,
	"m":  1,
	"km": 1000,

	"in": 0.0254,
	"ft": 0.3048,
	"yd": 0.9144,
	"mi": 1609.344,
}

func Length(value float64, from, to string) (float64, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, ErrInvalidValue
	}
	if value < 0 {
		return 0, ErrNegativeValue
	}

	fromFactor, ok := lengthToMeter[from]
	if !ok {
		return 0, fmt.Errorf("%w: %s", ErrUnknownUnit, from)
	}
	toFactor, ok := lengthToMeter[to]
	if !ok {
		return 0, fmt.Errorf("%w: %s", ErrUnknownUnit, to)
	}

	meters := value * fromFactor
	out := meters / toFactor
	return out, nil
}
