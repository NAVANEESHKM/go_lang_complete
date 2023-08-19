package main

import ("fmt")
func fun1(){
	defer fun2()
	fmt.Println("fun1")
}
func fun2(){

	defer fun3()
	fmt.Println("fun2")
}
func fun3(){
  fmt.Println("function 3")
}
func main(){
//    fmt.Println("I am first")
//    defer fmt.Println("PREDICT")
//    fmt.Println("Last")
//    defer fmt.Println("secpnd")

defer fun1()
fmt.Println("main fun")
}