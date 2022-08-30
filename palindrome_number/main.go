package main

import (
	"fmt"
	"math/big"
)

func main() {
	is := IsPalindrome(121)
	fmt.Println(is)
	is = IsPalindrome(-121)
	fmt.Println(is)
	is = IsPalindrome(10)
	fmt.Println(is)
}

func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	num := 0
	orig := x

	for i := x - 1; i < x; i-- {
		if i < 0 {
			break
		}
		x2, n := new(big.Int).DivMod(big.NewInt(int64(x)), big.NewInt(10), new(big.Int))
		x = int(x2.Int64())
		i = x
		num = num*10 + int(n.Int64())
	}
	return num == orig
}
