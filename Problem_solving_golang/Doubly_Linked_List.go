package main

import "fmt"

type don struct {
	prev *don
	val  int
	next *don
}

var head *don
var tail *don
var count=0

func addb(val int) {
	count++
	insert := &don{nil, val, nil}
	if tail == nil {
		head = insert
		tail = insert
	}
	insert.prev = tail
	tail.next = insert
	tail = insert
}
func addf(val int) {
	count++
	insert := &don{nil, val, nil}
	if tail == nil {
		head = insert
		tail = insert
	}
	insert.next = head
	head.prev = insert
	head = insert
}
func addm(val int,pos int){
	var temp*don=head
	count++
	insert := &don{nil, val, nil}
	if tail == nil {
		head = insert
		tail = insert
	}else{
	   for i:=0;i<pos-2;i++{
             temp=temp.next
	    }
		fmt.Println("Well")
		insert.next=temp.next
		temp.next.prev=insert
		temp.next=insert
		insert.prev=temp

}
}
func deletef() {
	count--
	head = head.next
	if head == nil {
		tail = nil
	}
}
func deleteb() {
	count--
	tail = tail.prev

}
func deletem(pos int){
	
	    
		if(pos==1){
			deletef()
		}else if(pos==count){
               deleteb()
		}else if(pos<=0){
			fmt.Println("Enter correct value")
		}else{
		pos=pos-1
       var temp*don=head
	   var supp*don
	   for true{
		   pos--
		  if(pos==0){
			break
		  }
		  temp=temp.next
		  
	   }
	   supp=temp
	   temp=temp.next
	   supp.next=temp.next
	   temp=temp.next
	   temp.prev=supp
	}
	count--

}
func display() {
	var temp*don=head
	for i:=0;i<count;i++{
		fmt.Println(temp.val)
        temp=temp.next
	}
}
func main() {
	for true{
		var n int
		var a=0
		fmt.Println("Enter a number\n1.add front\n2.add back\n3.add middle\n4.delete front\n5.Delete back\n6.Delete Middle\n7.Display\n8.stop")
		fmt.Scanln(&n)
		switch(n){
		case 1:{ var num int
				 fmt.Println("Enter the front value")
				 fmt.Scanln(&num)
				 addf(num)
			  }
			case 2:{ var num int
				fmt.Println("Enter the back value")
				fmt.Scanln(&num)
				addb(num)
			 }
			case 3:{
				var num int
				var pos int
				fmt.Println("Enter the middle value and its position")
				fmt.Scanln(&num)
				fmt.Scanln(&pos)
				addm(num,pos)
			}
		  case 4:{
			  if head != nil {
					  deletef()
			  }
		  }
		case 5:{
			if tail != nil {
				deleteb()
		}
	}
		case 6:{
			var pos int
			fmt.Println("Enter the position")
			fmt.Scanln(&pos)
			deletem(pos)
		}
		
		  case 7:display()
		  case 8: a=1
		}
		if(a==1) {break}
}
fmt.Println("Number of elements",count)

}
