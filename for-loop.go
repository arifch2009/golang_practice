package  main

import "fmt"

func main() {


	fmt.Println("For loop starts here")

	i := 1

        for i<=10 {
	fmt.Println("i=", i)
	i=i+1
	}


	for i:=20; i<30; i++ {

	fmt.Println("i=", i)
	}


	fmt.Println("Only odd numbers will be printed")

	for i:=0;i<20;i++ {
	
	if i%2==0 {
	fmt.Println("i=", i)
	}
	}

}
