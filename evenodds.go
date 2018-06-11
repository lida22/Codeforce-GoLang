package main

import (
	"fmt"
)

func main() {
	var inp, inp1, re, ans, k int
	fmt.Scan(&inp, &inp1)
	if inp%2 == 0 {
		re = inp / 2
	} else {
		re = (inp / 2) + 1
	}
	if re >= inp1 {
		ans = (inp1 * 2) - 1
	} else {
		k = inp1 - re
		ans = k * 2
	}
	fmt.Println(ans)
}

