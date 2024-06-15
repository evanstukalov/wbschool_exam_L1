package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1048577", 10)
	b.SetString("2097153", 10)

	sum := new(big.Int)
	diff := new(big.Int)
	prod := new(big.Int)
	quot := new(big.Int)
	rem := new(big.Int)

	sum.Add(a, b)
	diff.Sub(a, b)
	prod.Mul(a, b)
	quot.DivMod(a, b, rem)

	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())
	fmt.Printf("Сумма: %s\n", sum.String())
	fmt.Printf("Разность: %s\n", diff.String())
	fmt.Printf("Произведение: %s\n", prod.String())
	fmt.Printf("Частное: %s, Остаток: %s\n", quot.String(), rem.String())
}
