package problem380

import (
	"fmt"
	"testing"

	"golang.org/x/exp/rand"
)

type RandomizedSet struct {
    valueMap map[int]int
	values []int
}

func Constructor() RandomizedSet {
    return RandomizedSet{valueMap: make(map[int]int),values:make([]int,0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _,e:=this.valueMap[val];e{
		return false
	}
	this.values=append(this.values, val)
	this.valueMap[val]=len(this.values)-1
    
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if v,e:=this.valueMap[val];e{
		last := len(this.values) - 1
		this.values[v] = this.values[last]
		this.valueMap[this.values[v]] = v
		this.values = this.values[:last]
		delete(this.valueMap,val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.values[rand.Intn(len(this.values))]
}

func Test380(t *testing.T) {
	randomizedSet:=Constructor()
	randomizedSet.Insert(1)
	randomizedSet.Remove(2)
	randomizedSet.Insert(2)
	fmt.Println(randomizedSet.GetRandom())
	randomizedSet.Remove(1)
	randomizedSet.Insert(2)
	fmt.Println(randomizedSet.GetRandom())
}