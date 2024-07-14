package random

import (
	"fmt"
	"os"
)

func varibles() {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("The PATH is: %s\n", path)
	var a string = "meow"
	var b int = 5
	fmt.Println(a)
	fmt.Println(b)
}
