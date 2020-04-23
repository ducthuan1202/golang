# Tổng quan về Go 
[Tài liệu gốc](https://golang.org/doc/faq#nil_error)

## Mục đích
- tối ưu việc sử dụng nhiều lõi trên CPU 
- dễ dàng quản lý tài nguyên trong chương trình
- có khả năng thu gom rác để giải phóng bộ nhớ, quản lý bộ nhớ tự động

## Lịch sử
- Tác giả: `Robert Griesemer`, `Rob Pike` và `Ken Thompson`.
- Phác thảo mục tiêu cho Go vào ngày 21 tháng 09 năm 2007.
- Tháng 01 năm 2008, `Ken` đã bắt đàu làm việc với một trình biên dịch để khám phá các ý tưởng; Nó tạo mã C làm đầu ra của nó.
- Tháng 05 năm 2008, `Lan Taylor` bắt đầu lập giao diện người dùng GCC cho Go. Cuối năm, `Russ` Cõ đã tham gia vào giúp chuyển 
ngôn ngữ và thư việcn từ nguyên mẫu sang thực thể.
- Public vào ngày 10 tháng 11 năm 2009.

## Logo
- Logo (linh vật Go) được thiết kế bởi `Renée French` - người thiết kế ra Glenda, chú thỏ Plan 9.

## Tên chính thức 
- Tên chính thức là `Go`. 
- `Golang` được gọi là do trang web `golang.org` - trang chủ của Go.
- Là `Go`, không phải GO 

## Lý do chọn Go
- Go được sinh ra từ sự thất vọng với các ngôn ngữ và môi trường hiện tại cho công việc chúng tôi đang làm tại Google. Lập trình đã trở nên quá khó khăn và việc lựa chọn ngôn ngữ là một phần đáng trách. Người ta phải chọn biên dịch hiệu quả, thực thi hiệu quả hoặc dễ lập trình; cả ba đều không có sẵn trong cùng một ngôn ngữ chính. Các lập trình viên, những người có thể lựa chọn dễ dàng về sự an toàn và hiệu quả bằng cách chuyển sang các ngôn ngữ được gõ động như Python và JavaScript thay vì C ++ hoặc, ở mức độ thấp hơn, Java.

