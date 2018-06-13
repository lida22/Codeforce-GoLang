package main

import (
	"fmt"
)

func main() {
	var no, a, b, c, tmp, tmp1, tmp2 int
	fmt.Scan(&no)
	tmp = 0
	tmp1 = 0
	tmp2 = 0
	arr := []int{}
	arr1 := []int{}
	arr2 := []int{}
	for i := 0; i < no; i++ {
		fmt.Scan(&a, &b, &c)
		arr = append(arr, a)
		arr1 = append(arr1, b)
		arr2 = append(arr2, c)
	}
	for j := 0; j < len(arr); j++ {
		tmp = tmp + arr[j]
		tmp1 = tmp1 + arr1[j]
		tmp2 = tmp2 + arr2[j]
	}
	if tmp == 0 && tmp1 == 0 && tmp2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

