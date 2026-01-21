package convert_test

import (
	"errors"
	"math"
	"testing"

	"go-backend-labs/01-beginner/unit-converter/internal/convert"
)

func almostEqual(a, b, eps float64) bool {
	return math.Abs(a-b) <= eps
}

func TestLength(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		from  string
		to    string
		want  float64
	}{
		{"m to cm", 1, "m", "cm", 100},
		{"km to m", 2.5, "km", "m", 2500},
		{"mi to km", 1, "mi", "km", 1.609344},
		{"in to mm", 1, "in", "mm", 25.4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convert.Length(tt.value, tt.from, tt.to)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestWeight(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		from  string
		to    string
		want  float64
	}{
		{"kg to g", 1.5, "kg", "g", 1500},
		{"lb to kg", 1, "lb", "kg", 0.45359237},
		{"oz to g", 2, "oz", "g", 56.69904625},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convert.Weight(tt.value, tt.from, tt.to)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if !almostEqual(got, tt.want, 1e-12) {
				t.Fatalf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestTemperature(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		from  string
		to    string
		want  float64
	}{
		{"C to F", 0, "C", "F", 32},
		{"F to C", 212, "F", "C", 100},
		{"C to K", 0, "C", "K", 273.15},
		{"K to C", 273.15, "K", "C", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convert.Temperature(tt.value, tt.from, tt.to)
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	t.Run("unknown unit", func(t *testing.T) {
		_, err := convert.Length(1, "m", "nope")
		if err == nil {
			t.Fatal("expected error")
		}
		if !errors.Is(err, convert.ErrUnknownUnit) {
			t.Fatalf("expected ErrUnknownUnit, got: %v", err)
		}
	})

	t.Run("negative value disallowed for length", func(t *testing.T) {
		_, err := convert.Length(-1, "m", "cm")
		if !errors.Is(err, convert.ErrNegativeValue) {
			t.Fatalf("expected ErrNegativeValue, got: %v", err)
		}
	})

	t.Run("negative value allowed for temperature", func(t *testing.T) {
		got, err := convert.Temperature(-40, "C", "F")
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if !almostEqual(got, -40, 1e-9) {
			t.Fatalf("got %v want -40", got)
		}
	})
}
