package main

import (
	"fmt"
	"os"

	"github.com/twcrone/wman/pkg/string"
)

//type Salutation struct {
//	name     string
//	greeting string
//}

func main() {
	args := os.Args
	//	var s = Salutation{"todd", "hello"}
	message := "hello go world"
	var greeting = &message
	a, b, c := 1, 2, 3
	foo(&a)
	var name = args[1]
	fmt.Println(name, *greeting, a, b, c)
	fmt.Printf("hello %s, your name backward is %q", name, string.Reverse(name))
}

func foo(count *int) {
	fmt.Println("count: ", *count)
	*count++
}
