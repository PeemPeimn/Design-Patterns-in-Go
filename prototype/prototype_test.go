package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShape(t *testing.T) {
	tests := map[string]struct {
		baseShape     Shape
		expectedPrint string
	}{
		"Rectangle": {
			baseShape: &Rectangle{
				name:   "Rectangle",
				width:  10.01,
				height: 10.01,
			},
			expectedPrint: "Rectangle_clone: 10.01 10.01",
		},
		"Circle": {
			baseShape: &Circle{
				name:   "Circle",
				radius: 10.01,
			},
			expectedPrint: "Circle_clone: 10.01",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			clone := test.baseShape.Clone()

			assert.Equal(t, test.expectedPrint, clone.Print())
		})
	}
}

func TestShapeList(t *testing.T) {
	tests := map[string]struct {
		baseShape     Shape
		expectedPrint string
	}{
		"test list": {
			baseShape: &ShapeList{
				name: "ShapeList",
				shapes: []Shape{
					&Rectangle{
						name:   "Rectangle",
						width:  10.01,
						height: 10.01,
					},
					&Circle{
						name:   "Circle",
						radius: 10.01,
					},
				},
			},
			expectedPrint: "ShapeList_clone\nRectangle_clone: 10.01 10.01\nCircle_clone: 10.01\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			clone := test.baseShape.Clone()

			assert.Equal(t, test.expectedPrint, clone.Print())
		})
	}
}
