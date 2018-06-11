package main

import (
	"fmt"
)

func main() {
	var a, k, p, ans int
	fmt.Scan(&a)
	k = 0
	p = 0
	f := []int{}
	if a == 1 {
		fmt.Println("1")
	} else {
		for i := 1; i <= a+1; i++ {
			k = k + i
			p = p + k
			if p == a {
				ans = i
			}
			if p > a {
				f = append(f, i)
				break
			}
		}
		for _, j := range f {
			ans = j - 1
		}
		fmt.Println(ans)

	}
}

