package main

import (
	"fmt"
	"log"
	"net/http"
	"packages/greet"
	"packages/normalize"
	"strings"
)

func main() {

	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func sayHelloName(w http.ResponseWriter, r *http.Request) {

	fmt.Println("=================== Start Request ===================")

	fmt.Println("Get data with FormValue", r.FormValue("sort"))

	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("Key", k)
		fmt.Println("Val", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello astaxia!")

}

func myMap() {
	type IPAddr [4]byte

	var blackList map[string]IPAddr
	// blackList["loopback"] = IPAddr{1,0,0,1}

	blackList = make(map[string]IPAddr)
	blackList["loopback"] = IPAddr{1, 0, 0, 1}

	hosts := map[string]IPAddr{
		"googleDNS": {8, 8, 8, 8},
	}

	fmt.Println(blackList, hosts)
}

func usePackage() {
	normalize.Hello()
	greet.Hello()

	fmt.Println(greet.Shark)

	oct := greet.Octopus{"Jesse", "orange"}

	fmt.Println(oct.String())
}
