package problem2326

type ListNode struct {
	Val  int
	Next *ListNode
}

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	result:=make([][]int,m)
	for i:=0;i<m;i++{
		result[i]= make([]int, n)
		for j:=0;j<n;j++{
			result[i][j]=-2
		}
	}
	direction:=Right
	i,j:=0,0
	count:=0
	for count<m*n{
		if head!=nil{
			result[i][j]=head.Val
			head=head.Next
		}else{
			result[i][j]=-1
		}
		
		count++
		
		switch direction{
		case Right:
			if j<n-1&&result[i][j+1]==-2{
				j++
			}else{
				i++
				direction=Down
			}
		case Down:
			if i<m-1&&result[i+1][j]==-2{
				i++
			}else{
				j--
				direction=Left
			}
		case Left:
			if j>0&&result[i][j-1]==-2{
				j--
			}else{
				i--
				direction=Up
			}
		case Up:
			if i>0&&result[i-1][j]==-2{
				i--
			}else{
				j++
				direction=Right
			}
		}
	}
	return result
}
