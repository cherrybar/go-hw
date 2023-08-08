package main

import (
	"github.com/agrison/go-commons-lang/stringUtils"
)

func main() {
	str := "Hello, OTUS!"
	reversedSWtr := stringUtils.Reverse(str)
	print(reversedSWtr)
}
