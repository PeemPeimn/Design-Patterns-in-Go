package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	tests := map[string]struct {
		expectedCar Car
		mockBuilder func() Car
	}{
		"set all fields": {
			expectedCar: Car{
				seats:        4,
				engine:       "test engine",
				tripComputer: true,
				gps:          true,
			},
			mockBuilder: func() Car {
				b := CarBuilder{}
				b.SetSeats(4)
				b.SetEngine("test engine")
				b.SetTripComputer()
				b.SetGPS()
				return b.GetCar()
			},
		},
		"reset": {
			expectedCar: Car{},
			mockBuilder: func() Car {
				b := CarBuilder{}
				b.SetSeats(4)
				b.SetEngine("test engine")
				b.SetTripComputer()
				b.SetGPS()
				b.Reset()
				return b.GetCar()
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualCar := test.mockBuilder()

			assert.Equal(t, test.expectedCar, actualCar)
		})
	}
}

func TestDirector(t *testing.T) {
	tests := map[string]struct {
		expectedCar  Car
		mockDirector func(builder Builder) Car
	}{
		"SUV": {
			expectedCar: Car{
				seats:  4,
				engine: "SUV",
			},
			mockDirector: func(builder Builder) Car {
				d := Director{}
				return d.makeSUV(builder)
			},
		},
		"Sports Car": {
			expectedCar: Car{
				seats:        1,
				engine:       "Sports",
				tripComputer: true,
				gps:          true,
			},
			mockDirector: func(builder Builder) Car {
				d := Director{}
				return d.makeSportsCar(builder)
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			b := &CarBuilder{}
			actualCar := test.mockDirector(b)

			assert.Equal(t, test.expectedCar, actualCar)
		})
	}
}
