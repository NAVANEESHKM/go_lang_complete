package main

import "fmt"
type Point struct{
	X int
	Y int
	
}
// func (c Customer)isabove()string{
// 	c.Id=4574
// 	return "nandri"
// }
func(p Point)translate(dx,dy int){
	p.X=p.X+dx
    p.Y=p.Y+dy}
func main(){

	// s:=person{11,"jack"}
	// p:=&s
	// fmt.Println(p)
	// fmt.Println((*p).f1)
	// fmt.Println(p.f1)
	// (*p).f1=2333
	// fmt.Println(p.f1)
	// var p Customer
	// fmt.Println(p.isabove())
	p:=Point{3,4}
	fmt.Println("point p=",p)

	p.translate(7,9)
	fmt.Println("Aftrt translate p=",p)
}