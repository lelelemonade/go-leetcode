package problem134

func canCompleteCircuit(gas []int, cost []int) int {
	stationCount:=len(gas)
	start, end := 0, 0
	tank:=0

	if gas[0]>cost[0] {
		tank=tank+gas[0]-cost[0]
		end++
	}else {
		tank=tank+gas[0]-cost[0]
		start--
		start=(start+stationCount)%stationCount
	}

	for start!=end {
		
	}
}
