package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdapter(t *testing.T) {
	tests := map[string]struct {
		squarePeg    *SquarePeg
		holeRadius   float64
		expectedFits bool
	}{
		"The peg fits": {
			squarePeg:    &SquarePeg{2},
			holeRadius:   10.0,
			expectedFits: true,
		},
		"The peg does not fit": {
			squarePeg:    &SquarePeg{100},
			holeRadius:   10.0,
			expectedFits: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			hole := RoundHole{test.holeRadius}
			adaptedPeg := PegAdapter{test.squarePeg}
			actualFits := hole.Fits(&adaptedPeg)

			assert.Equal(t, test.expectedFits, actualFits)
		})
	}
}
