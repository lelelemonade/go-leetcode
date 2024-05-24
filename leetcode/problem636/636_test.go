package problem636

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func Test636(t *testing.T) {
	assert.Equal(t, []int{3, 4}, exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))
	assert.Equal(t, []int{8}, exclusiveTime(1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}))
	assert.Equal(t, []int{8, 1}, exclusiveTime(2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"}))
}

func exclusiveTime(n int, logs []string) []int {
	callStack := make([]Function, 0)
	result := make([]int, n)
	for _, log := range logs {
		logArray := strings.Split(log, ":")
		functionId, _ := strconv.Atoi(logArray[0])
		timestamp, _ := strconv.Atoi(logArray[2])
		if logArray[1] == "start" {
			callStack = append(callStack, Function{
				functionId: functionId,
				timestamp:  timestamp,
			})
		} else {
			pop := callStack[len(callStack)-1]
			callStack = callStack[:len(callStack)-1]

			result[pop.functionId] += timestamp - pop.timestamp + 1 - pop.minus
			if len(callStack) > 0 {
				callStack[len(callStack)-1].minus += timestamp - pop.timestamp + 1
			}
		}
	}

	return result
}

type Function struct {
	functionId int
	timestamp  int
	minus      int
}
