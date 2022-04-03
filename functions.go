package main

import "fmt"

func main() {
	// a := []int{1, 2, 3}
	// sum, nsum := sayM(a)
	// fmt.Printf("%v, %v", sum, nsum)

	// a, b := t_multuple_returns(5, 0)
	// fmt.Printf("%v %v\n", a, b)
	// if b != nil {
	// 	fmt.Printf("b is nil %v\n", b.Error())
	// }

	// Anonymous function
	ccc := func(x, y int) (int, int) {
		return x, y
	}
	fmt.Println(ccc(1, 2))

	var c func(x, y int) int = func(x, y int) int {
		return x + y
	}
	fmt.Println(c(1, 2))

	// Methods

	var sc sportCar = sportCar{
		name:  "FirstCar",
		speed: 255.00,
	}
	a, b := sc.tester()
	fmt.Printf("%v %v\n", a, b)
}

func sayM(msg []int) (i int, j int) {
	for _, v := range msg {
		i += v
		j = j - 1
	}
	return
}

func t_multuple_returns(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero %v", b)
	}
	return a / b, nil
}

type sportCar struct {
	name  string
	speed float32
}

func (x sportCar) tester() (string, float32) {
	fmt.Printf("Car name is %v, speed is %v\n", x.name, x.speed)
	return x.name, x.speed
}
