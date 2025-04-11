package main

// arr = {1, 5, 3, 2, 7, 6, 8, 3, 4, 8}
// Size = 10

// Numbers range from 1 to 8 (since n = 10, and n - 2 = 8)

// Repeating numbers: 3 and 8

import(
	"fmt"
)

func main(){
	var n int
	fmt.Print("ENter the range of numbers")
	fmt.Scan(&n)
	ar:=make([]int,n)
	fmt.Println("Enter array elements")
	for i:=0;i<n;i++{
		fmt.Scan(&ar[i])
	}
	//xor of array then xor of elements--->3^8=11
	 xornum:=0
	 for i:=0;i<n;i++{
		xornum^=ar[i]
	 }
	 for i:=1;i<=n-2;i++{
		xornum^=i
	}
	//check for right most set bit-->11 & -11 wll give
	setbit:=xornum & -xornum
	a:=0
	b:=0
	//sort array and n elements into set and unset bucket , return eahts left in both ->ans
	for i:=0;i<n;i++{
		if ar[i] & setbit != 0{
			a^=ar[i]
		}else{
			b^=ar[i]
		}	
	}
	for i:=1;i<=n-2;i++{
		if i & setbit != 0{
			a^=i
		}else{
			b^=i
		}
	}
	fmt.Printf("The two repeating numbers are %v and %v in the range of %v",a,b,n)
}