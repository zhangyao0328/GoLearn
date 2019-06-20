package main

import (
	_ "fmt"
	"fmt"
)

var a  = "hello world"

func main() {
}

func init() {

	for i := 0; i < len(a); i++ {
		fmt.Println(i)
	}

}


