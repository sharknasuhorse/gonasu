//P109
package main

import (
	"fmt"
)

const ONE = 1

func one() (int, int) {
	const TWO = 2
	return ONE, TWO
}

func main() {
	x, y := one()
	fmt.Println("x=%d, y=%d/n", x, y)
}
