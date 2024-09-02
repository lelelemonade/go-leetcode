package problem427

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	if len(grid) < 2 {
		return &Node{Val: grid[0][0] == 1, IsLeaf: true}
	}

	allEqual:=func(topLeftx, topLefty, bottomRightx, bottomRighty int)bool{
		count:=0
		for i:=topLeftx;i<=bottomRightx;i++{
			for j:=topLefty;j<=bottomRighty;j++{
				count+=grid[i][j]
			}
		}
		return count==0||count==(bottomRightx-topLeftx+1)*(bottomRighty-topLefty+1)
	}
	var daq func(topLeftx, topLefty, bottomRightx, bottomRighty int) *Node

	daq = func(topLeftx, topLefty, bottomRightx, bottomRighty int) *Node {
		if allEqual(topLeftx, topLefty, bottomRightx, bottomRighty){
			return &Node{Val: grid[topLeftx][topLefty] == 1, IsLeaf: true}
		}else if bottomRighty-topLefty == 1 {
			return &Node{
				Val:         grid[topLeftx][topLefty] == 1,
				IsLeaf:      false,
				TopLeft:     &Node{Val: grid[topLeftx][topLefty] == 1, IsLeaf: true},
				TopRight:    &Node{Val: grid[topLeftx][bottomRighty] == 1, IsLeaf: true},
				BottomLeft:  &Node{Val: grid[bottomRightx][topLefty] == 1, IsLeaf: true},
				BottomRight: &Node{Val: grid[bottomRightx][bottomRighty] == 1, IsLeaf: true},
			}
		}else {
			return &Node{
				Val:         grid[0][0] == 1,
				IsLeaf:      false,
				TopLeft:     daq(topLeftx, topLefty, (topLeftx+bottomRightx)/2, (topLefty+bottomRighty)/2),
				TopRight:    daq(topLeftx, (topLefty+bottomRighty)/2+1, (topLeftx+bottomRightx)/2, bottomRighty),
				BottomLeft:  daq((topLeftx+bottomRightx)/2+1, topLefty, bottomRightx, (topLefty+bottomRighty)/2),
				BottomRight: daq((topLeftx+bottomRightx)/2+1, (topLefty+bottomRighty)/2+1, bottomRightx, bottomRighty),
			}
		}
	}

	result := daq(0, 0, len(grid)-1, len(grid[0])-1)

	return result
}

func Test427(t *testing.T) {
	assert.Equal(t, nil, construct([][]int{
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
	}))

	// assert.Equal(t, nil, construct([][]int{
	// 	{0, 1},
	// 	{1, 0},
	// }))
}
