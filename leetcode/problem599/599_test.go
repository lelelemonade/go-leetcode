package problem599

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findRestaurant(list1 []string, list2 []string) []string {
    list1Map:=make(map[string]int)
    stringIdx:=make(map[string]int)
	for i,v:=range list1{
		list1Map[v]=i
	}
	for i,v:=range list2{
		if idx,e:=list1Map[v];e{
			stringIdx[v]=i+idx
		}
	}
	result:=make([]string,0)
	minIdxSum:=len(list1)+len(list2)
	for k,v:=range stringIdx{
		if v<minIdxSum{
			result=make([]string, 0)
			minIdxSum=v
		}
		if v==minIdxSum{
			result = append(result, k)
		}
	}
	return result
}

func Test599(t *testing.T) {
	assert.Equal(t,[]string{"Shogun"},findRestaurant([]string{"Shogun","Tapioca Express","Burger King","KFC"},[]string{"Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"}))
}