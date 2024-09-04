package problem874

type direction int

const (
	Up direction = iota
	Right
	Down
	Left
)

type position struct {
	x int
	y int
}

func robotSim(commands []int, obstacles [][]int) int {
	obstaclesMap := make(map[position]struct{})
	for _, v := range obstacles {
		obstaclesMap[position{x: v[0], y: v[1]}] = struct{}{}
	}
	currentPosition := position{0, 0}

	currentDirection := Up

	result:=0
	for _, v := range commands {
		if v == -1 {
			currentDirection = (currentDirection + 5) % 4
		} else if v == -2 {
			currentDirection = (currentDirection + 3) % 4
		}else{
			switch currentDirection {
			case Up:
				for i:=1;i<=v;i++{
					currentPosition.y++
					if _,e:=obstaclesMap[currentPosition];e{
						currentPosition.y--
						break
					}
				}
			case Right:
				for i:=1;i<=v;i++{
					currentPosition.x++
					if _,e:=obstaclesMap[currentPosition];e{
						currentPosition.x--
						break
					}
				}
			case Down:
				for i:=1;i<=v;i++{
					currentPosition.y--
					if _,e:=obstaclesMap[currentPosition];e{
						currentPosition.y++
						break
					}
				}
			case Left:
				for i:=1;i<=v;i++{
					currentPosition.x--
					if _,e:=obstaclesMap[currentPosition];e{
						currentPosition.x++
						break
					}
				}
			}

			result = max(result,currentPosition.x*currentPosition.x+currentPosition.y*currentPosition.y)

		}
	}

	return result
}
