package problem50


func myPow(x float64, n int) float64 {
	var positivePow func(x float64,n int)(result float64)
	positivePow=func(x float64,n int)(result float64){
		if n==1{
			return x
		}
		next:=positivePow(x,n>>1)
		if n&1==0{
			return next*next
		}else{
			return next*next*x
		}
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