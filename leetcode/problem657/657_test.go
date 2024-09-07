package problem657

func judgeCircle(moves string) bool {
	directionMap:=make(map[rune]int)
	for _,move:=range moves{
		directionMap[move]++
	}
	return directionMap['U']-directionMap['D']==0&&directionMap['L']-directionMap['R']==0
}