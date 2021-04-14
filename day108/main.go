package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

type People struct {
	name string `json:"name"`
}

func test1() {
	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}

var ip string
var port int

func init() {
	flag.StringVar(&ip,"ip","0.0.0.0","ip address")
	flag.IntVar(&port,"port",8000,"port number")
}

func test2() {
	flag.Parse()
	fmt.Printf("%s:%d", ip, port)
}

func main() {
	test1()
	test2()
}
