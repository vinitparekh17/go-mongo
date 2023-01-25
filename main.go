package main

import "fmt"

func main() {
	fmt.Println("MongoDB")
}

func fetal(err error) {
	if err != nil {
		panic(err)
	}
}
