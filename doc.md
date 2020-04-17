## 1. HTTP/HTTPS là gì?

**HTTP**
- Là viết tắt của `HyperText Transfer Protocol` - Giao thức truyền tải siêu văn bản.

- Là một trong các giao thức chuẩn về mạng Internet, dùng để trao đổi thông tin giữa máy chủ (server) và máy khách (client).

- HTTP là một giao thức ứng dụng của bộ giao thức TCP/IP (giao thức nền tảng cho internet).

- HTTP là một giao thức không có trạng thái - `stateless protocol`, hay nói cách khác, request sau không biết gì về request trước đó, cho dù cả 2 request có thể cùng từ 1 client. Để giải quyết vấn đề này, các ứng dụng web sửu dụng cookie để duy trì trạng thái kết nối.

**HTTPS**
- Là viết tắt của `HyperText Transfer Protocol Secure`.

- Là HTTP + giao thức bảo mật SSL hay TSL cho phép trao đổi thông tin một các bảo mật trên internet.

- Giao thức HTTPS sử dụng port `443`, giúp bảo đảm các tính chât sau của thông tin:

    - `Confientiality` - bảo mật: phương thức này mã hóa gói tin, bảo đảm các thông điệp trao đổi giữa cliet và server không bị bên thứ 3 đọc được.

    - `Integrity` - chính trực: sử dụng phương thức hashing để cả client và server đều nhận được thông tin nguyên vẹn, không bị mất mát hay chỉnh sửa gì.

    - `Authenticity` - xác thực: sử dụng chứng chỉ số (digital certificate) để xác thực địa chỉ

## 2. TCP/IP là gì?
- `TCP`: Là viết tắt của Transmission Control Protocol - Giao thức điều khiển truyền dẫn

- `IP`: Là viết tắt của Internet Protocol - Địa chỉ đơn nhất mà những thiết bị điện tử hiện nay đang sử dụng để nhận diện và liên lạc với nhau trên mạng máy tính.

=> `TCP/IP`: Là việc điều khiển truyền dẫn dữ liệu giữa các máy tính với nhau ở trong mạng máy tính thông qua địa chỉ máy.

- `TCP/IP` là một chồng giao thức với 4 tầng (layer):

    + **Tầng 1**: Network Access Layer - Tầng truy cập: tầng này xác định chi tiết về cách thức dữ liệu được gửi qua mạng, bởi các thiết bị phần cứng trực tiếp giao tiếp 
    với môi trường mạng, chẳng hạn như cáp đồng trục, cap quang hay dây đồng xoắn đôi.

    - **Tầng 2**: Internet Layer: Tầng Internet: tầng này đóng gói dữ liệu và các gói dữ liệu được biết đến dưới dạng các gói tin thông qua giao thức Internet Protocol, 
    có chứa địa chỉ nguồn và địa chỉ đích được sử dụng để chuyển tiếp các gói tin giữa máy chủ và qua các mạng.

    - **Tầng 3**: Transport Layer: Tầng vận chuyển: tầng này thực hiện việc vận chuyển các gói tin từ địa chỉ nguồn tới địa chỉ đích và bảo đảm việc vận chuyển toàn vẹn dữ liệu, đúng thứ tự.

    - **Tầng 4**: Application Layer: Tầng ứng dụng: tầng này cung cấp các ứng dụng cho phép người dùng thực hiện việc trao đổi dữ liệu.

## 3. URL là gì?
- Là viết tắt của `Uniform Resource Locator`, hiểu nôm na là địa chỉ cụ thể để truy cập tới một tài nguyên trên mạng máy tính.

- `URL` gồm có 5 phần: (1)scheme, (2)host, (3)port, (4)uri (hoặc path) chính

    + `scheme`: là giao thức được sử dụng (http, https, ftp, ...)

    + `host`: là IP hoặc tên của DNS domain chứa tài nguyên.

    + `port`: là cổng TCP trên server dùng để lắng nghe request từ phía client.

    + `uri`: Uniform Resource Identifier - là một chuỗi ký tự dùng để xác định tài nguyên.

    + `query-string`: là phần dữ liệu kèm theo trong lần request đó.



```bash

scheme://host[:port#]/path/.../[?query-string][#anchor]

scheme          # Giao thức sử dụng để kết nối (có thể là: HTTP, HTTPS, FTP)
host            # IP hoặc domain name của server
port#           # Cổng kết nối, mặc định là 80        
path            # Đường dẫn cụ thể
query-string    # phần dữ liệu gửi lên server
anchor          # anchor - một vị trí trên trang (ví dụ: #content)
```

## 4. DNS - Domain Name System

- `DNS`: Domain Name System là một hệ thống quản lý tên miền, nó chuyển đổi domain sang địa chỉ IP thực để kết nối.

## 5. HTTP request package
Hay có thể nói là thông tin trình duyệt, gói tin được gửi lên trong 1 lần request. Gồm có 3 phần 

- `request line`: 

- `request header`

- `request body`

```bash
GET /domains/example/ HTTP/1.1      # request line: request method, URL, protocol and its version
Host：www.iana.org                  # domain name
User-Agent：Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.4 (KHTML, like Gecko) Chrome/22.0.1229.94 Safari/537.4   # browser information
Accept：text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8   # mime that clients can accept
Accept-Encoding：gzip,deflate,sdch      # stream compression
Accept-Charset：UTF-8,*;q=0.5           # character set in client side
// request header       # blank line          
// request body,        # request resource arguments (for example, arguments in POST)
```

## 6. HTTP Response status 
- 1xx Informational: 
- 2xx Success
- 3xx Redirection
- 4xx Client Error
- 5xx Server Error