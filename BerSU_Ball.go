package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x, y, g, no int

	fmt.Scan(&y)
	no = 0
	a := make([]int, y)
	for i := 0; i < y; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&x)
	b := make([]int, x)
	for j := 0; j < x; j++ {
		fmt.Scan(&b[j])
	}
	re := reflect.DeepEqual(a, b)
	if re == true {
		fmt.Println(len(a))
	} else {
		for k := 0; k < len(a); k++ {
			for g = 0; g < len(b); g++ {

				if a[k]-b[g] == 0 || a[k]-b[g] == 1 || a[k]-b[g] == -1 {
					g = 0
					//fmt.Println(a[k], " : ", b[g])
					b = RemoveIndex(b, g)
					k = k + 1

					no = no + 1
				}
			}
		}
		fmt.Println(no)
	}

}
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

