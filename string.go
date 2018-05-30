package main

import (
	"fmt"	
	"strings"
)

func main(){
	var wrd,word string
	fmt.Scan(&wrd)
	word=strings.ToLower(wrd)
	
	s := []string{}
	for _,e1 := range word{
		if((string(e1)=="a") || (string(e1) =="e") || (string(e1)=="i") || (string(e1)=="o") || (string(e1)=="y") || (string(e1)=="u")){
			continue
		}else{
			s=append(s,".")
			s=append(s,string(e1))
		}			
	}
	result1 := strings.Join(s, "")
    	fmt.Println(result1)
}


