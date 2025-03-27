package main

import (
	"fmt"
)

type Number int

func (n *Number) Double() {
	*n *= 2
}

func main() {
	n := Number(4)
	fmt.Println("Первоначальное значение номера", n)
	n.Double()
	fmt.Println("Новое значение номера", n)
}
