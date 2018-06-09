package main
import(
	"fmt"
)
func main(){
	var i, j,r1,r2 int
	fmt.Scan(&i)
	fmt.Scan(&j)
	var nume int
	nume=i/j
	var deno int
	deno=i%j
	var result int
	var result2,totalx int
	totalx=0
	result=nume+deno
	s := []int{}
	if(i<j){
		fmt.Println(i)
	}else{
		if(result>=j){
			for(result>=j){
				r1=result/j
				r2=result%j
				result=r1+r2
				s=append(s, r1)
			}
		for _, valuex := range s {
			//valuex=int(valuex)
			totalx =totalx+ int(valuex)
		}
		
		result2 = totalx+nume+i
		fmt.Println(result2)
		}else{
			result2 = nume+i
			fmt.Println(result2)
			
		}
	}
	
}
