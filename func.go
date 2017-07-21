package main


import "fmt"


func print_(){
	fmt.Println("Inside print function")

}
func add(a int, b int) {


	fmt.Println("sum is:", a+b)

}
func main(){

	fmt.Println("Inside the main function")
	print_()
	add(1,2)
	fmt.Println("End of the program")

}
