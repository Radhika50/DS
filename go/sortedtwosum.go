package main

//given sorted array so use of two pointer to find sum equals to target
import(
	"fmt"
)

func main(){
	//n,target,ar,res
	var n int
	fmt.Println("Enter the array size")
	fmt.Scan(&n)
	ar:=make([]int,n)
	fmt.Println("Enter sorted array")
	for i:=0;i<n;i++{
		fmt.Scan(&ar[i])
	}
	var target int
	fmt.Println("Enter target")
	fmt.Scan(&target)
	res:=[]int{}
	l:=0
	e:=n-1
	for l < e{
		if ar[l]+ar[e]==target{
			res=append(res,l)
			res=append(res,e)
			break
		}else if ar[l]+ar[e] > target{
			e--
		}else{
			l++
		}
	}
	for _,ele:=range res{
		fmt.Print(ele," ")
	}
	fmt.Println()
}