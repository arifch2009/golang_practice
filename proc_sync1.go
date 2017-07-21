package main



import ("fmt"
	"time"
//	"math/rand"

)

func download(pings chan<- string){

	time.Sleep(5 * time.Second)
	pings <- "This is done"
}

func extract(ping chan string){
     msg:= <-ping
     fmt.Println("this is inside extract", msg)

}
func main(){

	fmt.Println("This is a sync and wait program")
	pings := make(chan string )
	go func() {

		download(pings)
	}()
	defer extract(pings)
	
}
