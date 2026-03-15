package problem1622

type Fancy struct {
	operations []Operation
	vals       []Val
}

type Operation struct {
	addVal  int
	multVal int
}

type Val struct {
	currentVal   int
	operationIdx int
}

func Constructor() Fancy {
	return Fancy{
		operations: make([]Operation, 0),
		vals:       make([]Val, 0),
	}
}

func (this *Fancy) Append(val int) {
	this.vals = append(this.vals, Val{
		currentVal:   val,
		operationIdx: len(this.operations),
	})
}

func (this *Fancy) AddAll(inc int) {
	this.operations = append(this.operations,
		Operation{
			addVal: inc,
		})
}

func (this *Fancy) MultAll(m int) {
	this.operations = append(this.operations,
		Operation{
			multVal: m,
		})

}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= len(this.vals) || idx < 0 {
		return -1
	}
	for this.vals[idx].operationIdx < len(this.operations) {
		operation := this.operations[this.vals[idx].operationIdx]

		if operation.addVal != 0 {
			this.vals[idx].currentVal = (this.vals[idx].currentVal + operation.addVal) % 1000000007
		}

		if operation.multVal != 0 {
			this.vals[idx].currentVal = (this.vals[idx].currentVal * operation.multVal) % 1000000007
		}

		this.vals[idx].operationIdx++
	}

	return this.vals[idx].currentVal
}

/**
 * Your Fancy object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Append(val);
 * obj.AddAll(inc);
 * obj.MultAll(m);
 * param_4 := obj.GetIndex(idx);
 */
