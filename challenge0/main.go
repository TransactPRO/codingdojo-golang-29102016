package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func solve(str string) (res string, err error) {
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		var curRes string
		curRes, err = computeResult(numbers)
		if err != nil {
			return
		}
		res += curRes + " "
	}

	res = res[:len(res)-1]
	return
}
func main() {
	fmt.Println("hello fom golang rally!")
}

func computeResult(array []string) (res string, err error) {
	if len(array) != 3 {
		err = errors.New("Array len not 3")
		return
	}
	var a, b, c int
	a, err = strconv.Atoi(array[0])
	if err != nil {
		err = errors.New("String can not be converted to integer")
		return
	}
	b, err = strconv.Atoi(array[1])
	if err != nil {
		err = errors.New("String can not be converted to integer")
		return
	}
	c, err = strconv.Atoi(array[2])
	if err != nil {
		err = errors.New("String can not be converted to integer")
	}
	calculationResult := strconv.Itoa(a*b + c)

	var sumOfDigits int
	for _, charrune := range calculationResult {
		var a int
		a, err = strconv.Atoi(string(charrune))

		if err != nil {
			err = errors.New("String can not be converted to integer")
		}

		sumOfDigits += a
	}
	res = strconv.Itoa(sumOfDigits)
	return

}
