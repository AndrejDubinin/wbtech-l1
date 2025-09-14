package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b int64

	a = 1 << 21
	b = 1 << 22

	c := a + b

	fmt.Printf("%d + %d = %d\n", a, b, c)

	c = b - a
	fmt.Printf("%d - %d = %d\n", b, a, c)

	c = b / a
	fmt.Printf("%d / %d = %d\n", b, a, c)

	var d big.Int
	d.Mul(big.NewInt(a), big.NewInt(b))
	fmt.Printf("%d * %d = %s\n", a, b, d.Text(10))
}
