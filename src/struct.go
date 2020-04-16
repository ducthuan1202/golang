package main

import . "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r *Rectangle) acreage() float64 {
	return (r.height + r.width) * 2
}

func (r Rectangle) doubleWidth() Rectangle {
	r.width = r.width * 2
	return r
}

func (r *Rectangle) doubleWidthWithPointer() {
	r.width = r.width * 2
}

func main() {

	r1 := Rectangle{width: 3, height: 4}

	r1 = r1.doubleWidth()

	r2 := r1
	r2.doubleWidthWithPointer()

	// r1.doubleWidthWithPointer()

	Println(r1, r1.area(), r1.acreage(), r2)

}
