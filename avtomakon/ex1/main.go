package main

import (
	"fmt"
	"math/rand"
)

func randCustom(n int) []int {
	m := make(map[int]struct{})

	for {
		x := rand.Int()

		m[x] = struct{}{}

		if len(m) == n {
			break
		}
	}

	res := make([]int, 0, n)

	for i, _ := range m {
		res = append(res, i)
	}

	return res
}

func main() {
	fmt.Println(randCustom(5))
}
