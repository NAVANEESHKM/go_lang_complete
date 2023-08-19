package main

import "fmt"

func main(){
     x:=f1(2,3)
	 fmt.Println(x,"well")
}

func f1(a,b int)func()int{
	c:=34
	return func() int{
		return a+b+c
	}
}