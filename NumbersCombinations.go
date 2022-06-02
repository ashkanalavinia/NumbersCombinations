package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	cs := combin.Combinations(len(data), 3)
	for _, c := range cs {
		fmt.Printf("%d%d%d\n", data[c[0]], data[c[1]], data[c[2]])
	}
}

