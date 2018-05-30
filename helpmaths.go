package main
import(
	"fmt"
	"sort"
	"strings"
)
func main() {
	var i string
	fmt.Scan(&i)
	result := strings.Split(i, "+")
	sort.Sort(sort.StringSlice(result))//sort the array
	s := []string{}
	for _,element := range result{
			s=append(s, element)
			s=append(s,"+")
		}
	result1 := strings.Join(s, "")
    	fmt.Println(result1[:len(s)-1])
}

