package mylog

import (
	"fmt"
	"time"
	"os"
	"net/http"
)

const logPath = "logs"

func init(){	
	if _, err := os.Stat(logPath); os.IsNotExist(err) {		
		os.Mkdir(logPath, 0777)
	}
}

func Add(r *http.Request){

	dt := time.Now()
	dateString := dt.Format("2006-01-02")
	logFile := fmt.Sprintf("%s/%s.log", logPath, dateString)

	// localhost:3000 2020-04-20 17:57:46.455172 +0700 +07 m=+1.744359396
	// fmt.Println(r.Host, dt)

	fout, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)		
	if err != nil {
		return
	}

	defer fout.Close()

	content := fmt.Sprintf("[%s] %s %s \r\n", dt.Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
	
	fout.WriteString(content)
}