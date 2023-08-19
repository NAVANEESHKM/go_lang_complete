package main

import ("fmt" 
       "time"
	)

func main(){
	start:=time.Now()
	a,b:=0,1
	var n int
	fmt.Scanln(&n)
	fmt.Print(a," ",b)
	var c int
	n=n-2
	for true{
		fmt.Print(" ",a+b," ")
		c=a+b
		a=b
		b=c
      n--
	  if(n==0){
		break
	  }
	}
	end:=time.Now().Sub(start)
	fmt.Println(end)	
}