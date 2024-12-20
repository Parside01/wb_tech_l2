package visitor

import (
	"math"
	"testing"
)

func TestAreaCalcVisitor(t *testing.T) {
	tt := []struct {
		name   string
		shape  Shape
		expect float64
	}{
		{
			name:   "Circle",
			shape:  &Circle{Radius: 4},
			expect: 4 * 4 * math.Pi,
		},
		{
			name:   "Rectangle",
			shape:  &Rectangle{Left: 2, Right: 8},
			expect: 2 * 8,
		},
		{
			name:   "Square",
			shape:  &Square{Side: 4},
			expect: 4 * 4,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			visitor := &AreaCalc{}
			tc.shape.accept(visitor)
			if visitor.Area != tc.expect {
				t.Error("Expected", tc.expect, "got", visitor.Area)
			}
		})
	}
}
