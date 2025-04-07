package main

import(
	"fmt"
)
func helper(ar []int) []int{
	var sum int64=0
	var sqsum int64=0
	for i:=0; i<len(ar); i++{
		temp:=int64(ar[i])
		sum+=temp
		sum-=int64(i+1)
		sqsum+=temp*temp
		sqsum-=int64((i+1)*(i+1))
	}
	sumab:=sqsum/sum
	a:=(sum+sumab)/2
	b:=sumab-a
	return []int{int(a),int(b)}

}
func main(){
	fmt.Println("finding repeating and missing number given an array from 1 to N")
	var n int
	fmt.Print("Enter array size")
	fmt.Scan(&n)

	ar:=make([]int,n)
	fmt.Println("Enter elements")
	for i:=0; i<n; i++{
		fmt.Scan(&ar[i])
	}
	result:=helper(ar)
	fmt.Println("Duplicate",result[0])
	fmt.Println("Missing",result[1])
}

/*
	vector<int> solve(vector<int> &ar){
		long long sum=0;//a+b
		long long sqsum=0;
		long long trmp;
		for(int i=0;i<ar.size();i++){
			temp=ar[i];
			sum+=temp; //cal sum
			sum=sum-(i+1) //expected
			sqsum+=(tmp*temp);
			sqsum=sqsum-(long(i*i)*long(i*i))
		}
		//for a+b
		long long sumab=(sqsum)/sum
		int a=(sum+sumab)/2
		int b=(sumab)-a
		vector<int> ans;
		ans.push_back(a)
		ans.push_back(b)
		return ans
	}
*/