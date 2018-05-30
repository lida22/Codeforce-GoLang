package main
import( 	
	"fmt"
)
func main(){
	var i int
	fmt.Scan(&i)
	var j string
	fmt.Scan(&j)
	scale:=j
	k:=0
	for i,_ := range scale{
		if(i==len(scale)-1){
				break}
		if(scale[i]==scale[i+1]){
			k=k+1
			
		}
		

	
}
fmt.Println(k)

}
