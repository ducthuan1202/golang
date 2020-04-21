# Go Overview

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