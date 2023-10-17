package builder

type Car struct {
	seats        int
	engine       string
	tripComputer bool
	gps          bool
}

type Builder interface {
	Reset()
	SetSeats(seats int)
	SetEngine(engine string)
	SetTripComputer()
	SetGPS()
	GetCar() Car
}

type CarBuilder struct {
	car Car
}

func (c *CarBuilder) Reset() {
	c.car = Car{}
}

func (c *CarBuilder) SetSeats(seats int) {
	c.car.seats = seats
}

func (c *CarBuilder) SetEngine(engine string) {
	c.car.engine = engine
}

func (c *CarBuilder) SetTripComputer() {
	c.car.tripComputer = true
}

func (c *CarBuilder) SetGPS() {
	c.car.gps = true
}

func (c *CarBuilder) GetCar() Car {
	return c.car
}
