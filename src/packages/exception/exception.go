package exception

import (
	"log"
	"os"
)

/** Kiểm tra và kích hoạt panic nến có lỗi */
func CheckErr (err error){
	if err != nil {
		panic(err)
	}
}

/** Log lỗi và dùng chương trình */
func ShowErr(){
	if r := recover(); r != nil {
		// log.Fatal có kèm theo câu lệnh os.Exit(). Dừng chương trình và bỏ qua cả lệnh defer
		log.Fatal(r)
	}
}

/** Ghi lại nội dung lỗi vào file */
func WriteLog(err error){

	// mở file, nếu chưa tồn tại thì tạo mới
	f, _ := os.OpenFile("logs/error.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)		
	defer f.Close()	

	// Khởi tạo logger
	logger := log.New(f, "", log.LstdFlags)	// 	LstdFlags: định dạng thời gian (2020/04/22 14:06:08)

	// Ghi dữ liệu vào file
	logger.Println(err)
}