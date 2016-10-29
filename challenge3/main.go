package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Yea")
}

func reverse(str string) (reversed string) {
	split := strings.Split(str, " ")
	for i := len(split) - 1; i >= 0; i-- {
		reversed += split[i] + " "
	}

	reversed = reversed[:len(reversed)-1]

	return
}

func reverseSlice(array []string) {
	for i, slice := range array {
		array[len(array)-i-1] = slice
	}
}
