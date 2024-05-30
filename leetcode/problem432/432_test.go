package problem432

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

type node struct {
	keys  map[string]struct{}
	count int
}

type AllOne struct {
	allOneList       *list.List
	stringCountNodes map[string]*list.Element
}

func Constructor() AllOne {
	return AllOne{
		allOneList:       list.New(),
		stringCountNodes: make(map[string]*list.Element),
	}
}

func (this *AllOne) Inc(key string) {
	if v, e := this.stringCountNodes[key]; e {
		currentNodeValue, _ := v.Value.(node)

		addNewNode := func() {
			newKeysMap := make(map[string]struct{})
			newKeysMap[key] = struct{}{}
			currentNodeCount := currentNodeValue.count
			currentNodeCount++
			newKeyNode := node{
				keys:  newKeysMap,
				count: currentNodeCount,
			}
			this.stringCountNodes[key] = this.allOneList.InsertAfter(newKeyNode, v)
			delete(currentNodeValue.keys, key)
			if len(currentNodeValue.keys) == 0 {
				this.allOneList.Remove(v)
			}
		}

		if v.Next() == nil {
			addNewNode()
		} else {
			nextNodeValue, _ := v.Next().Value.(node)
			if nextNodeValue.count > currentNodeValue.count+1 {
				addNewNode()
			} else {
				delete(currentNodeValue.keys, key)
				nextNodeValue.keys[key] = struct{}{}
			}
		}
	} else {
		firstNodeValue := this.allOneList.Front()
		addFrontNode := func() {
			newKeysMap := make(map[string]struct{})
			newKeysMap[key] = struct{}{}
			newKeyNode := node{
				keys:  newKeysMap,
				count: 1,
			}

			this.stringCountNodes[key] = this.allOneList.PushFront(newKeyNode)
		}
		if firstNodeValue == nil {
			addFrontNode()
			return
		}
		firstNode := firstNodeValue.Value.(node)

		if firstNode.count == 1 {
			firstNode.keys[key] = struct{}{}
			this.stringCountNodes[key] = firstNodeValue
		} else {
			addFrontNode()
		}
	}
}

func (this *AllOne) Dec(key string) {
	if v, e := this.stringCountNodes[key]; e {
		currentNodeValue, _ := v.Value.(node)

		if currentNodeValue.count <= 1 {
			delete(currentNodeValue.keys, key)
			if len(currentNodeValue.keys) == 0 {
				this.allOneList.Remove(v)
			}
			delete(this.stringCountNodes, key)
			return
		}

		addNewNode := func() {
			newKeysMap := make(map[string]struct{})
			newKeysMap[key] = struct{}{}
			currentNodeCount := currentNodeValue.count
			currentNodeCount--
			newKeyNode := node{
				keys:  newKeysMap,
				count: currentNodeCount,
			}
			this.stringCountNodes[key] = this.allOneList.InsertBefore(newKeyNode, v)
			delete(currentNodeValue.keys, key)
			if len(currentNodeValue.keys) == 0 {
				this.allOneList.Remove(v)
			}
		}

		if v.Prev() == nil {
			addNewNode()
		} else {
			prevNodeValue, _ := v.Prev().Value.(node)
			if prevNodeValue.count < currentNodeValue.count-1 {
				addNewNode()
			} else {
				delete(currentNodeValue.keys, key)
				prevNodeValue.keys[key] = struct{}{}
			}
		}
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.allOneList.Len() == 0 {
		return ""
	}
	maxValueNode, _ := this.allOneList.Back().Value.(node)
	for k := range maxValueNode.keys {
		return k
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.allOneList.Len() == 0 {
		return ""
	}
	minValueNode, _ := this.allOneList.Front().Value.(node)
	for k := range minValueNode.keys {
		return k
	}
	return ""
}

func Test432(t *testing.T) {
	//testCase1 := Constructor()
	//
	//testCase1.Inc("hello")
	//testCase1.Inc("hello")
	//assert.Equal(t, "hello", testCase1.GetMaxKey())
	//assert.Equal(t, "hello", testCase1.GetMinKey())
	//
	//testCase1.Inc("leet")
	//assert.Equal(t, "hello", testCase1.GetMaxKey())
	//assert.Equal(t, "leet", testCase1.GetMinKey())

	testCase2 := Constructor()
	testCase2.Inc("hello")
	testCase2.Inc("goodbye")
	testCase2.Inc("hello")
	testCase2.Inc("hello")
	assert.Equal(t, "hello", testCase2.GetMaxKey())

	testCase2.Inc("leet")
	testCase2.Inc("code")
	testCase2.Inc("leet")
	testCase2.Dec("hello")
	testCase2.Inc("leet")
	testCase2.Inc("code")
	testCase2.Inc("code")
	assert.Equal(t, "leet", testCase2.GetMaxKey())

}
