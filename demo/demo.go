package main

import (
	"fmt"
	"log"
	"os"
)

func main(){

	file, err := os.Open("command.txt")
	if err != nil {
		log.Fatal(err)
	}

	var login string
	var password string

	fmt.Fscan(file, login)

	fmt.Fscan(file, password)
}
