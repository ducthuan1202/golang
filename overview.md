# Tổng quan về Go 
[Tài liệu gốc](https://golang.org/doc/faq#nil_error)

## Mục đích
- Là ngôn ngữ hướng đến các lập trình viên.
- Tối ưu việc sử dụng nhiều lõi trên CPU.
- Dễ dàng quản lý tài nguyên trong chương trình.
- Có khả năng thu gom rác để giải phóng bộ nhớ, quản lý bộ nhớ tự động.

## Lịch sử
- Tác giả: `Robert Griesemer`, `Rob Pike` và `Ken Thompson`.
- Phác thảo mục tiêu cho Go vào ngày 21 tháng 09 năm 2007.
- Tháng 01 năm 2008, `Ken` đã bắt đàu làm việc với một trình biên dịch để khám phá các ý tưởng; Tạo mã C làm đầu ra của Go.
- Tháng 05 năm 2008, `Lan Taylor` bắt đầu lập giao diện người dùng GCC cho Go. Cuối năm, `Russ` Cõ đã tham gia vào giúp chuyển 
ngôn ngữ và thư việc từ nguyên mẫu sang thực thể.
- Public vào ngày 10 tháng 11 năm 2009.

## Logo
- Logo (linh vật Go) được thiết kế bởi `Renée French` - người thiết kế ra Glenda, chú thỏ Plan 9.

## Tên chính thức 
- Tên chính thức là `Go`. 
- `Golang` được gọi là do trang web `golang.org` - trang chủ của Go.
- Là `Go`, không phải GO 

## Lý do chọn Go
- Go được sinh ra từ sự thất vọng với các ngôn ngữ và môi trường hiện tại cho công việc chúng tôi đang làm tại Google. Lập trình đã trở nên quá khó khăn và việc lựa chọn ngôn ngữ là một phần đáng trách. Người ta phải chọn biên dịch hiệu quả, thực thi hiệu quả hoặc dễ lập trình; cả ba đều không có sẵn trong cùng một ngôn ngữ chính. Các lập trình viên, những người có thể lựa chọn dễ dàng về sự an toàn và hiệu quả bằng cách chuyển sang các ngôn ngữ được gõ động như Python và JavaScript thay vì C ++ hoặc, ở mức độ thấp hơn, Java.

## Các kiểu dữ liệu
```go
bool // true | false

string // chuỗi

int  int8  int16  int32  int64 // số nguyên
uint uint8 uint16 uint32 uint64 uintptr // số nguyên dương

byte // alias for uint8

rune // alias for int32, represents a Unicode code point

float32 float64 // số thập phân

complex64 complex128
```

## Quản lý packages
```bash
go mod init github.com/ducthuan1202/bidgear
```

Dòng lệnh trên sẽ tạo ra 1 file `go.mod` tại thư mục hiện hành với nội dung như sau:
```mod
module github.com/ducthuan1202/bidgear

go 1.13
```

Với mọi thư mục có chứa file `go.mod` thì Go sẽ coi đó là 1 package, và file `go.mod` là file 
quản lý các packages phụ thuộc (dependencies) - giống như file `package.json` trong Nodejs

Nếu trong bất kỳ file .go nào ở trong thư mục có import thư viện (không phải package có sẵn trong Go) 
thì thông tin đều được thêm vào đây. Trong lần chạy file go tiếp hệ, Go command sẽ tự động download các gói 
về và install vào thư mục `pkg/mod` tại thư mục GOROOT
