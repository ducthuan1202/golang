package greet

import "fmt"

// export 1 biến
var Shark = "Duc Thuan"

// export 1 struct
type Octopus struct {
	Name string
	Color string 
}

func (o Octopus) String() string {
	return fmt.Sprintf("The octopus's name is %q and is the color %s.", o.Name, o.Color)
}

// Đây là 1 phương thức gọi được từ bên ngoài
func Hello() {
	fmt.Println("Hello world")
}

// Đây là 1 phương thức chỉ dùng nội bộ (chữ cái đầu tiên ko viết hoa)
func makeup(){
	fmt.Println("makeup")
}
