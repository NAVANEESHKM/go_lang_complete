package main

import "fmt"
//credit card validation
func help(num int) int {
	var sum = 0
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return sum
}
func main() {
	fmt.Println("Enter the card number")
	var a int
	fmt.Scanln(&a)
	var dig = 10
	var oddp = 0
	var even = 0
	var rul = 1
	for a > 0 {
		if rul%2 != 0 {
			oddp = oddp + a%dig

		} else {
			val := a % dig
			if val+val > 9 {
				even = even + help(val+val)

			} else {
				even = even + (val + val)

			}

		}
		a = a / 10
		rul++
	}

	if (oddp+even)%10 == 0 {
		fmt.Println("valid")
	} else {
		fmt.Println("Invalid")
	}
}
