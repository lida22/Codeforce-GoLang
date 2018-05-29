package main

import "fmt"

func main() {
	 var i int
  	 fmt.Scan(&i)
	if i<=100 && i>2{
		if i%2 == 0 {
			fmt.Println("YES")
		}else {
			fmt.Println("NO")
		}
	}else {
		fmt.Println("NO")
	}
}
