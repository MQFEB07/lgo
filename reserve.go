package lgo

// 432. All O`one Data Structure
// type AllOne struct {
// 	Keys map[string]*A1Node
// 	Min  *A1Node
// 	Max  *A1Node
// }

// func ConstructorAo() AllOne {
// 	left := &A1Node{Value: 0}
// 	right := &A1Node{Value: int(math.MaxInt64)}
// 	left.Next = right
// 	right.Prev = left
// 	return AllOne{
// 		Keys: make(map[string]*A1Node),
// 		Min:  left,
// 		Max:  right,
// 	}
// }

// func (this *AllOne) Inc(key string) {
// 	v, ok := this.Keys[key]
// 	if !ok {
// 		newNode := A1Node{
// 			Key:   key,
// 			Value: 1,
// 		}
// 		this.Keys[key] = &newNode
// 		this.Min = &newNode

// 		left := this.Min
// 		right := this.Min.Next
// 		left.Next = v
// 		right.Prev = v
// 		v.Prev = left
// 		v.Next = right
// 		return
// 	}

// 	v.Value++
// 	v.Next.Prev = v.Prev
// 	v.Prev.Next = v.Next
// 	cur := v.Next
// 	for cur.Value < v.Value {
// 		cur = cur.Next
// 	}

// 	prev := cur.Prev
// 	nex := cur
// 	prev.Next = v
// 	nex.Prev = v
// 	v.Prev = prev
// 	v.Next = nex
// }

// func (this *AllOne) Dec(key string) {
// 	node, ok := this.Keys[key]
// 	if !ok {
// 		panic("Key does not exist")
// 	}

// 	node.Value--
// 	if node.Value == 0 {
// 		delete(this.Keys, key)

// 		left := node.Prev
// 		right := node.Next
// 		left.Next = right
// 		right.Prev = left
// 		node.Prev = nil
// 		node.Next = nil
// 		return
// 	}

// 	node.Prev.Next = node.Next
// 	node.Next.Prev = node.Prev

// 	cur := node.Prev
// 	for cur.Value > node.Value {
// 		cur = cur.Prev
// 	}

// 	left := cur
// 	right := cur.Next
// 	left.Next = node
// 	right.Prev = node
// 	node.Prev = left
// 	node.Next = right
// }

// func (this *AllOne) GetMaxKey() string {
// 	return this.Max.Key
// }

// func (this *AllOne) GetMinKey() string {
// 	return this.Min.Key
// }

// type A1Node struct {
// 	Key   string
// 	Value int
// 	Next  *A1Node
// 	Prev  *A1Node
// }

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

// 1593. Split a String Into the Max Number of Unique Substrings

func maxUniqueSplit(s string) int {
	m := map[string]bool{}
	ans := 0

	var backtrack func(int)
	backtrack = func(start int) {
		if start == len(s) {
			ans = max(ans, len(m))
			return
		}

		for i := start; i < len(s); i++ {
			sub := s[start : i+1]
			if !m[sub] {
				m[sub] = true
				backtrack(i + 1)
				delete(m, sub)
			}
		}
	}
	backtrack(0)
	return ans
}
