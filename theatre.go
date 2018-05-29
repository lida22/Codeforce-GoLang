package main

import "fmt"

func main() {
	var i,j,k int
  	fmt.Scan(&i)
	fmt.Scan(&j)
	fmt.Scan(&k)
        var b,c,d int
	b=j/k
	c=i/k
	if j%k !=0{
		b = b+1
	}
	if i%k !=0{
		c = c+1
	}
	d = b*c
	fmt.Println(d)
}
 

	
