package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	set(n)
}

func set(n int) {
	var sum, total, tmp, tmp1, k, p int
	sum = 0
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	for _, k = range a {
		sum = sum + k
	}

	total = (sum / 2) + 1

	tmp = 0
	tmp1 = 0
	for _, p = range a {
		tmp = tmp + p
		tmp1 = tmp1 + 1
		if tmp >= total {
			break
		}
	}
	fmt.Println(tmp1)
}

