package main
import "fmt"
func main() {
	var n int
	fmt.Println("Enter the size")
	fmt.Scanln(&n)
	arr:=make([][]int,n)
	for i:=0;i<n;i++{
		arr[i]=make([]int,n)
	}
	var s = n - 1
	var k = 1
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			if j == 0 {
				arr[i][j] = 1
			} else if j == k-1 {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j] + arr[i-1][j-1]
			}
		}
		k++
	}
	for ind, val := range arr {
		for i := 0; i < s; i++ {
			fmt.Print(" ")
		}
		s--
		for cind, _ := range val {
			if arr[ind][cind] == 0 {
				continue
			}
			fmt.Print(arr[ind][cind], " ")
		}
		fmt.Println()
	}
}
