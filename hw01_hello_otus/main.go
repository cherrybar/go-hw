package main

import (
	"fmt"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func main() {
	str := "Hello, OTUS!"
	reversedSWtr := stringUtils.Reverse(str)
	fmt.Println(reversedSWtr)
}
