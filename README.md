# Go Overview
Sử dụng `go doc` để xem hướng dẫn sử dụng package, method trong package

```bash
# xem hướng dẫn sử dụng hàm Sprintf trong gói fmt
go doc fmt Sprintf
```
## Setup
- Sau khi cài đặt GO, cần set GOPATH cho thư mục dự án
```bash
# thực hiện gán đường dẫn cho GOPATH
export PATH=/usr/local/bin:$PATH

# ví dụ đường dẫn trên máy Mac
export GOPATH=${HOME}/Documents/mygo

# áp dụng thay đổi
source ~/.bash_profle
```

## Lưu ý trong Go
- Khi viết package, đối với biến, phương thức, fields, ... nếu viết hoa chữ cái đầu tiên thì sẽ được coi như public và có thể gọi từ bên ngoài. Nếu viết thường, thì chỉ có thể sử dụng trong nội bộ gói đó.

```bash
export PATH=/usr/local/bin:$PATH

source ~/.bash_profile
```

- `defer` yêu cầu thực thi 1 câu lệnh ở cuối cùng của function

```go
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

    fmt.Println("done")

    /*
    counting
    done
    9
    8
    7
    6
    5
    4
    3
    2
    1
    0
    */
```

- `pointer` biến con trỏ

```go
// khai báo biến nhưng chưa gán giá trị
var p *int

i := 42

// khai báo biến và gán giá trị
c := &i

p = c

i = 12

// gán giá trị thông qua biến con trỏ
*p = 1202

// lấy giá trị từ biến con trỏ
fmt.Println(i, *p, *c)
```

- `struct` 

- sử dụng kiểu any
```go
type any interface{}

func dump(input any) {
	Println(input)
}
```

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

```

Các int, uintvà uintptrcác loại thường rộng 32 bit trên các hệ thống 32 bit và rộng 64 bit trên các hệ thống 64 bit. 
Khi bạn cần một giá trị số nguyên, bạn nên sử dụng inttrừ khi bạn có lý do cụ thể để sử dụng loại số nguyên có kích thước hoặc không dấu.

## try catch 
```go
package main

import (
	"fmt"
)

type Block struct {
    Try     func()
    Catch   func(Exception)
    Finally func()
}
 
type Exception interface{}
 
func Throw(up Exception) {
    panic(up)
}
 
func (tcf Block) Do() {
    if tcf.Finally != nil {
 
        defer tcf.Finally()
    }
    if tcf.Catch != nil {
        defer func() {
            if r := recover(); r != nil {
                tcf.Catch(r)
            }
        }()
    }
    tcf.Try()
}


func main()  {

	fmt.Println("We started")
    Block{
        Try: func() {
            fmt.Println("I tried")
            Throw("Oh,...sh...")
        },
        Catch: func(e Exception) {
            fmt.Printf("Caught %v\n", e)
        },
        Finally: func() {
            fmt.Println("Finally...")
        },
    }.Do()

	fmt.Println("We went on")	
	
}

```