package main 

import (
	"fmt"
	"time"
)

/**
- Tại sao hàm sum() lại bắt buộc phải là 1 goroutines?
- Tại sao [x, y := <-c, <-c] thứ tự trả về lại bị đảo ngược nhau?
*/
func main () {

	s := []int {7, 2, 8, -9, 4, 0}

	c := make(chan int)

	s1 := s[:len(s)/2]
	s2 := s[len(s)/2:]


	go sum(s1, c) // goroutines 1
	go sum(s2, c) // goroutines 2

	x, y := <-c, <-c // nhận dữ liệu từ c	

	// bị đảo ngược giá trị trả về
	fmt.Println(s1, x) // [7 2 8] -5
	fmt.Println(s2, y) // [-9 4 0] 17
}

func sum(s []int, c chan int){
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // gửi giá trị sum tới kênh c
}