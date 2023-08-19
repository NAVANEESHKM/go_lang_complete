package main

import ("fmt")

func main(){
	p:=[]int{11,22,33,44,55,66,77,88,99}
	for _,val:=range p{
		fmt.Println(val)
	}
}