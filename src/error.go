package main

import (
	"fmt"
	"errors"
	. "packages/exception"
)

func main()  {
	defer ShowErr()

	s, err := Sqrt(-1)

	WriteLog(err)
	CheckErr(err)

	fmt.Println("pass", s)
}

func Sqrt (f float64) (float64, error){
	if f < 0 {
		return 0, errors.New("something went wrong")
	}
	return f * f, nil
}