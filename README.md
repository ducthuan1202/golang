# Go Overview

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
