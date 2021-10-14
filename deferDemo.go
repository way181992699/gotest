package main

import (
	"fmt"
	"strings"
)

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {

	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
	fmt.Println()
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y)) //A 10  20
	// calc("AA", x, calc("A", x, y)) //A 1 2 3 \n AA 1 3 4
	x = 10
	defer calc("BB", x, calc("B", x, y))
	// calc("BB", x, calc("B", x, y)) // B 10 2 12  \n BB 10 12 22
	y = 20

}