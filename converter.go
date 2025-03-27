package main

import (
	"fmt"
)

type Rub float64
type Sum float64
type Euro float64

func (r Rub) toSum() Sum {
	return Sum(r * 153.98)
}

func (e Euro) toSum() Sum {
	return Sum(e * 14025)
}

func (s Sum) toRub() Rub {
	return Rub(s / 153.97)
}

func (e Euro) toRub() Rub {
	return Rub(e * 90.91)
}

func main() {
	rusMoney := Rub(56)
	fmt.Printf("%0.2f рублей равняется %0.2f сум\n", rusMoney, rusMoney.toSum())
	euroMoney := Euro(321)
	fmt.Printf("%0.2f евро равняется %0.2f сум\n", euroMoney, euroMoney.toSum())
	sumMoney := Sum(535210)
	fmt.Printf("%0.2f сум равняется %0.2f рублей\n", sumMoney, sumMoney.toRub())
	fmt.Printf("%0.2f евро равняется %0.2f рублей\n", euroMoney, euroMoney.toRub())
}
