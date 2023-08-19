package main

import (
	"fmt"
)

// func main(){
// 	var input int
// 	var res=1
// 	  fmt.Scanln(&input)
// 	  for i:=input;i>0;i--{
// 		   res=res*i
// 	  }
// 	  fmt.Println(res)
// }

var fins=1
func fact(num int){
	if(num==1){
		return 0
	}
     fins=fins*num
	return fact(num-1)
}
func main(){
      var value=5
	  res:=fact(value)
	  fmt.Println(res)
}