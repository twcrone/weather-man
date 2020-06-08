package main

import (
	"os"

	//	"github.com/twcrone/wman/pkg/string"
	"github.com/twcrone/wman/pkg/cmd"
)

//type Salutation struct {
//	name     string
//	greeting string
//}

func main() {
	if err := cmd.NewWmanCmd().Execute(); err != nil {
		os.Exit(-1)
	}
	//args := os.Args
	//	var s = Salutation{"todd", "hello"}
	////message := "hello go world"
	////var greeting = &message
	////a, b, c := 1, 2, 3
	////foo(&a)
	////var name = args[1]
	//fmt.Println(name, *greeting, a, b, c)
	//fmt.Printf("hello %s, your name backward is %q", name, string.Reverse(name))
}
