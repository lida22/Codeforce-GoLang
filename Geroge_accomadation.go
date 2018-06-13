/*
Codeforce - A. George and Accommodation
*/
package main
import (
	"fmt"
)
func main() {
	var no, a, b, avail, tmp, sum int
	fmt.Scan(&no)
	avail = 0
	tmp = 0
	arr := []int{}
	arr1 := []int{}
	//total := []int{}
	for i := 0; i < no; i++ {
		fmt.Scan(&a, &b)
		arr = append(arr, a)
		arr1 = append(arr1, b)
	}
	for j := 0; j < len(arr); j++ {
		if arr[j] == arr1[j] {
			tmp = 0
			//total = append(total, avail)
		}
		if arr[j] < arr1[j] {
			if arr1[j] - arr[j] >= 2 {
				avail = avail + 1
				//total = append(total, avail)	
			}
		}
	}
	sum = tmp + avail
	fmt.Println(sum)

}

