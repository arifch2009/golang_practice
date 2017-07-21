

package main

import "fmt"

func main() {

	fmt.Println("Before the thread")
	defer func() {
		fmt.Println("Inside the thread 1")
	}()
	if true {
		func() {
			fmt.Println("This is inside the thread 1")
		}()
	}


	if true {
		func() {
			fmt.Println("This is inside the thread 2")
		}()
	}
	fmt.Println("After all the thread")
}
