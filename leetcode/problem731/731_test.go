package problem731

type pair struct{ start, end int }
type MyCalendarTwo struct{ booked, overlaps []pair }

func Constructor() MyCalendarTwo {
    return MyCalendarTwo{}
}

func (c *MyCalendarTwo) Book(start, end int) bool {
    for _, p := range c.overlaps {
        if p.start < end && start < p.end {
            return false
        }
    }
    for _, p := range c.booked {
        if p.start < end && start < p.end {
            c.overlaps = append(c.overlaps, pair{max(p.start, start), min(p.end, end)})
        }
    }
    c.booked = append(c.booked, pair{start, end})
    return true
}
