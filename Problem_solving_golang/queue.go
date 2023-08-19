package main

import "fmt"

var front *don
var back *don
type don struct {
	val  int
	next *don
}



func add(val int) {
	include := &don{val, nil}
	if back == nil {
		front = include
		back = include
	}
	back.next = include
	back = include

}
func delete() {
	front = front.next
	if front == nil {
		back = nil
	}
}

func display() {
	var temp *don = front
	for temp.next != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
	fmt.Println(temp.val)

}
func main() {
      for true{
		      var n int
			  var a=0
			  fmt.Println("Enter a number\n1.add\n2.delete\n3.Display\n4.stop")
			  fmt.Scanln(&n)
			  switch(n){
			  case 1:{ var num int
				       fmt.Println("Enter the value")
				       fmt.Scanln(&num)
				       add(num)
					}
				case 2:{
					if front != nil {
					        delete()
					}
				}
				case 3: display()
				case 4: a=1
			  }
			  if(a==1) {break}
	  }
	
}




