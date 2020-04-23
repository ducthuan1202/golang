package main

import "fmt"

func greet(c chan string){
	val := <- c
	fmt.Println("Hello " + val + "!")
}

func react(c chan string) {
	fmt.Printf("Xin chào %s \n", <-c)
}

func main() {
	fmt.Println("main() stared")	
	c := make(chan string)
	
	username := []string{"Duc Thuan", "Huyen Trang"}
	
	for _, name := range username {
		go greet(c)				
		go react(c)
		c <- name
		c <- name
	}

	fmt.Println("main() stopped")
}
