package problem2300

import "slices"

func successfulPairs(spells []int, potions []int, success int64) []int {
    slices.Sort(potions)

	for _,v:=range spells{
		start:=0
		end:=len(potions)-1
		iter:=(start+end)/2
		for {
			if potions[iter]*v>=start && potions[iter]*v<=end{

			}
		}
	}
}