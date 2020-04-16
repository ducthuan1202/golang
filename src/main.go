package main

import (
	"fmt"
	"runtime"
)

// import "strings"

func main() {
	os := runtime.GOOS

	switch os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}

	fmt.Println(os, runtime.NumGoroutine(), runtime.NumCPU(), runtime.GOARCH)

}
