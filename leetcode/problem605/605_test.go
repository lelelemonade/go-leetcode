package problem605

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n==0{
		return true
	}
	if len(flowerbed)==1{
		return flowerbed[0]==0&&n==1
	}
	for i := 0; i < len(flowerbed); i++ {
		if i == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		} else if i == len(flowerbed)-1 && flowerbed[i] == 0 && flowerbed[i-1] == 0 {
			flowerbed[i] = 1
			n--
		} else if i != 0 && 
		i != len(flowerbed)-1 &&
			flowerbed[i-1] == 0 &&
			 flowerbed[i] == 0 &&
			flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
		}
	}
	return n<=0
}
