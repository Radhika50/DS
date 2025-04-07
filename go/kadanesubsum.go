package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Print("enter size of array: ")
	fmt.Scan(&n)

	ar:=make([]int,n)
	fmt.Println("Enter the elements")
	for i:=0; i <n; i++{
		fmt.Scan(&ar[i])
	}
	//main code
	var maxi= -1<<31
	curr:=0
	for i:=0; i<n; i++{
		curr=curr+ar[i]
		maxi=max(maxi,curr)
		if curr<0{
			curr=0
		}
	}
	fmt.Print("maximum sub array sum: ",maxi)



}