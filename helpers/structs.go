package helpers

type Instance struct {
	Name  string
	Count int
}

type Costum struct {
	What  string
	Count int
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}
