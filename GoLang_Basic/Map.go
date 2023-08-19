package main

import("fmt")
func main(){
	m:=make(map[string]int)
	m["salary"]=200
	m["age"]=20
	m["emloyeeId"]=1
    for i,j:=range m{
		fmt.Println(i," ",j)
	}
	m["sachin"]=360
//to find if elements already exist
	x,isExisting:=m["age"]
	fmt.Println(x)
	fmt.Println(isExisting)
	fmt.Println(m)
	delete(m,"age")
	fmt.Println(m)
}

