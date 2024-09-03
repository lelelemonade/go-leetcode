package problem50

import "math"

func myPow(x float64, n int) float64 {
	math.Pow()
	positivePow:=func(x float64,n int)(result float64){
		result=1
		for i:=0;i<n;i++{
			result*=x
		}
		return
	}
    if n==0||x==1||(x==-1&& n%2!=1){
		return 1
	}else if x==-1&& n%2==1{
		return -1
	}else if n<0{
		if -n>20000000{
			return 0
		}
		return 1/positivePow(x,-n)
	}else{
		return positivePow(x,n)
	}
}