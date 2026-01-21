package convert

import (
	"fmt"
	"math"
)

const kelvinOffset = 273.15

func Temperature(value float64, from, to string) (float64, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, ErrInvalidValue
	}

	k, err := toKelvin(value, from)
	if err != nil {
		return 0, err
	}
	out, err := fromKelvin(k, to)
	if err != nil {
		return 0, err
	}
	return out, nil
}

func toKelvin(v float64, from string) (float64, error) {
	switch from {
	case "K":
		return v, nil
	case "C":
		return v + kelvinOffset, nil
	case "F":
		return (v-32)*5.0/9.0 + kelvinOffset, nil
	default:
		return 0, fmt.Errorf("%w: %s", ErrUnknownUnit, from)
	}
}

func fromKelvin(k float64, to string) (float64, error) {
	switch to {
	case "K":
		return k, nil
	case "C":
		return k - kelvinOffset, nil
	case "F":
		return (k-kelvinOffset)*9.0/5.0 + 32, nil
	default:
		return 0, fmt.Errorf("%w: %s", ErrUnknownUnit, to)
	}
}
