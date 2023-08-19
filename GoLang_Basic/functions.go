package main

import ("fmt"
"errors"
)
func get1(a1,a2 float64)(float64,float64){
	return a1*10,a2*10
}

func get2(p1,p2 float64)(float64,float64,error){
	if(p1==0){
		err:=errors.New("Previous price cannot be zero")
		
		return 0,0,err
	}

	return p1*10,p2*10,nil
}
func get3(p1,p2 float64)(c1,c2 float64){
	c1=p1*10
	c2=p2*10
	return c1,c2
}
func main(){
	// c1,_:=get1(12.7,43.8)
// fmt.Println(c1," ")


// c1,c2,err:=get2(0,2)
// fmt.Println(c1,c2,err)
// if(err!=nil){ fmt.Println("super") }
   
c1,c2:=get3(2,3)
fmt.Println(c1,"",c2)

}