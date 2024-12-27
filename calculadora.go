package main

import "fmt"

func main() {
	x := soma(1, 2, 3)
	y := multiplicacao(10, 10)
	w := subtracao(5, 10)
	z := divisao(20)
	fmt.Println(x, y, w, z)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func subtracao(i ...int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}

func multiplicacao(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}

	return total
}

func divisao(i ...int) int {
	total := 1
	for _, v := range i {
		total = v / total
	}
	return total
}
