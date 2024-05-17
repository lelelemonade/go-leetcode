package problem841

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test841(t *testing.T) {
	//testCase1 := [][]int{{1}, {2}, {3}, {}}
	//
	//assert.Equal(t, true, canVisitAllRooms(testCase1))

	testCase2 := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}

	assert.Equal(t, false, canVisitAllRooms(testCase2))
}

func canVisitAllRooms(rooms [][]int) bool {
	unvisitedRoom := make(map[int]struct{})
	for i := 0; i < len(rooms); i++ {
		unvisitedRoom[i] = struct{}{}
	}

	//bfs
	keyStack := make([]int, 0)
	keyStack = append(keyStack, 0)

	for len(keyStack) != 0 {
		key := keyStack[0]
		if _, e := unvisitedRoom[key]; !e {
			keyStack = keyStack[1:]
			continue
		}
		delete(unvisitedRoom, key)
		for i := 0; i < len(rooms[key]); i++ {
			if rooms[key][i] != 0 {
				keyStack = append(keyStack, rooms[key][i])
			}
		}
		keyStack = keyStack[1:]
	}

	return len(unvisitedRoom) == 0
}
