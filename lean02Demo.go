package main

import "fmt"

func main() {

	var name = []int{1, 1, 2, 2, 3, 3, 2, 4, 4, 5, 5, 67, 67, 98, 98, 45}

	var sum int16

	for _, v := range name {
		i := 0
		for _, e := range name {
			sum++
			if v-e == 0 {
				i++
			}
		}
		if i == 1 {
			fmt.Printf("只出现一次的数为:%d,共循环了%d次", v, sum)

			break
		}

	}

}
