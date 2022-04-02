package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start main func")
	t_panicker()
	fmt.Println("End main func")
}

func t_defer() {
	fmt.Println("Start")

	// Middle will be executed after the function, but before it will return, if there is couple defer
	// it will be executed bu LIFO order (reverse)
	msg := "Middle"
	defer fmt.Println(msg)
	msg = "NotMiddle"
	// Eventghough msg was rewriter, arguments evaluated at the time defer is executed, not at the time of called function executes
	fmt.Println("Finish")
}

func t_panic() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	http.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello web!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func t_panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("ERROR:", err)
			log.Println("error handeled by recover")
			// panic(err)
		}
	}()
	panic("Some thing went wrong")
	fmt.Println("End panicker")
}
