package builder

type Director struct {}

func (d *Director) makeSUV(builder Builder) Car {
  builder.SetSeats(4)
  builder.SetEngine("SUV")
  return builder.GetCar()
}

func (d *Director) makeSportsCar(builder Builder) Car {
  builder.SetSeats(1)
  builder.SetEngine("Sports")
  builder.SetTripComputer()
  builder.SetGPS()
  return builder.GetCar()
}
