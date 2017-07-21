package main

import "fmt"


// function_retun() return another function which is inside the own body
func func_return() (func()){

	return func() {
	fmt.Println("This is inside the function")
	}

}

func main() {
	

	func_return()
	/* function inside the main
	x:= func (x int) {
	fmt.Println("This is inside the function within main", x+1)
	}
	x(1)
       */
	

}
