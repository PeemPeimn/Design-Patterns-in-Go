package composite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComposite(t *testing.T) {
	tests := map[string]struct {
		component     Component
		expectedPrice int
	}{
		"Nested box": {
			component: &Box{
				[]Component{
					&Product{10},
					&Box{
						[]Component{
							&Product{20},
							&Product{30},
						},
					},
					&Product{40},
				},
			},
			expectedPrice: 100,
		},
		"Single product": {
			component:     &Product{100},
			expectedPrice: 100,
		},
		"Single box": {
			component:     &Box{},
			expectedPrice: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualPrice := test.component.GetPrice()

			assert.Equal(t, test.expectedPrice, actualPrice)
		})
	}
}
