package problem729

type MyCalendar struct {
	calendar [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{calendar: make([][]int, 0)}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, v := range this.calendar {
		if v[0] < end && start < v[1] {
			return false
		}
	}

	this.calendar = append(this.calendar, []int{start, end})
	return true
}
