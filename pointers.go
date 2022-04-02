package main

import "fmt"

type Car struct {
	speed int
}

func main() {
	t_pointer()
	t_pointer_syntacticSugar()
}

func t_pointer() {
	a := 42
	b := &a
	*b = 22
	c := &b
	fmt.Printf("Var value : %v, var type %T\n", a, a)
	fmt.Printf("Var value : %v, var type %T\n", b, b)
	fmt.Printf("Var value : %v, var type %T", **c, c)
}

func t_pointer_syntacticSugar() {

	// Without syntactic sugar
	var c0 *Car
	c0 = &Car{}
	(*c0).speed = 1
	// (*c0).speed syntax means : the '.' (dot) operator has the higher precedence on '*'
	// And if you write *c0.speed, Go will try to dereference c0.speed which is not a pointer.
	// Thats why we enclose with the bracket (*c0) so Go first dereference *co and then access value speed on this object.
	fmt.Printf("%v\n", (*c0).speed)

	// With syntactic sugar
	c1 := &Car{speed: 20}
	// You can omit the syntax (*c1) and go will dereference it for you i.e. (*c1.speed) == c1.speed
	c1.speed = 123
	fmt.Printf("%v\n", c1.speed)

}
