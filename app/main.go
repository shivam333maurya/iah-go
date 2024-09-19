package main

import (
	"fmt"
	"i-am-here/app/internal/server"
)

func main() {

	server := server.IAHServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
	// var name string = "II"
	// var name1 string = "IIIIII"
	// n, err := fmt.Println("I am here", name, name1)
	// fmt.Println(n, err)
}
