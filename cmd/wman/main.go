package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

func main() {
	var s = Salutation{"todd", "hello"}
	message := "hello go world"
	var greeting = &message
	a, b, c := 1, 2, 3
	foo(&a)
	fmt.Println(s.greeting, s.name, *greeting, a, b, c)
}

func foo(count *int) {
	fmt.Println("count: ", *count)
	*count++
}
