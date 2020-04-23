package main

import (
	"fmt"
	"crypto/rand"
)

func main()  {

	buf := make([]byte, 16)
	a, err := rand.Read(buf)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf), a)
	
	b, err := rand.Read(buf)

	fmt.Println(string(buf), b)
}

// byte alias of int8
// rune alias of int32

func test(){
	var i8 int8 = 10

	// khi chuyển từ string => byte, cần sử dụng dấu ngoặc đơn
	var b byte = byte('a')

	// Nếu là 1 slice thì sử dụng dấu nháy kép
	var c []byte = []byte("A")

	fmt.Println(i8, b, c)

	// khai báo kiểu int32
	var d int32 = 'A'

	// khai báo kiểu rune
	var e rune = 'A'
	fmt.Printf("%T %v \n",e, d)

	// int32 == rune
	fmt.Println("int32 == rune", e == d)	

	// int8 != byte
	var f int8 = 10
	var g byte = 10
	fmt.Printf("%T %v \n", f, f)
	fmt.Printf("%T %v \n", g, g)
}

func seem()  {
	myText := "Nguyễn Đức Thuận"

	// Chuyển đổi string sang kiểu byte
	myByte := []byte(myText)
	for _, vl := range(myByte) {
		fmt.Printf("%d : %s \n",vl, string(vl))	
	}
	fmt.Println(myByte, string(myByte), len(myByte))
	
	// Chuyển đổi string sang kiểu rune (int32)
	myRunes := []rune(myText)
	for _, vl := range(myRunes) {
		fmt.Printf("%d : %s \n",vl, string(vl))	
	}
	fmt.Println(myRunes, string(myRunes), len(myRunes))
}
