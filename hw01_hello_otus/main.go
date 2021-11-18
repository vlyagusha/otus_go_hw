package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	reverseHelloString := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(reverseHelloString)
}
