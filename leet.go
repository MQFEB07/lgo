package lgo

import (
	"container/heap"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	var result *ListNode

	if list1.Val > list2.Val {
		result = list2
		result.Next = mergeTwoLists(list1, list2.Next)
	} else {
		result = list1
		result.Next = mergeTwoLists(list1.Next, list2)
	}

	return result
}

func reverseList(head *ListNode) *ListNode {
	// 	newHead := &ListNode{} //Val is 0
	var newHead *ListNode // nil
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}

func finalPositionOfSnake(n int, commands []string) int {
	var result int = 0
	for _, command := range commands {
		if command == "UP" {
			result -= n
		} else if command == "DOWN" {
			result += n
		} else if command == "RIGHT" {
			result++
		} else if command == "LEFT" {
			result--
		}
	}
	return result
}

func isPalindrome(head *ListNode) bool {
	f, s := head, head
	for f.Next != nil && f.Next.Next != nil {
		f = f.Next.Next
		s = s.Next
	}
	s.Next = reverseList(s.Next)
	s = s.Next
	for s != nil {
		if head.Val != s.Val {
			return false
		}
		head, s = head.Next, s.Next
	}

	return true
}
func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	steps := 0
	// if n < 2 {
	// 	return 0
	// }
	// if n < 4 || n == 5 || n == 7 {
	// 	return n
	// }
	// primeNum := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	// for i := 0; i < 12; i++ {
	// 	if n%primeNum[i] == 0 {
	// 		steps += primeNum[i]
	// 		steps += minSteps(n / primeNum[i])
	// 		break
	// 	}
	// }

	// for i := 2; i <= n; i++ {
	// 	for n%i == 0 {
	// 		steps += i
	// 		n /= i
	// 	}
	// }
	i := 2
	for n%i == 0 {
		steps += i
		n /= i
	}
	if i <= n {
		i++
	}
	for {
		for n%i == 0 {
			steps += i
			n /= i
		}
		i += 2
		if i*i > n {
			break
		}
	}
	if n > 1 {
		steps += n
	}
	return steps
}

func isIsomorphic(s string, t string) bool {
	s1 := []byte(s)
	t1 := []byte(t)
	map1 := make(map[byte]byte, 0)
	map2 := make(map[byte]byte, 0)
	l := len(s)
	for i := 0; i < l; i++ {
		val, ok := map1[s1[i]]
		_, ok1 := map2[t1[i]]
		if ok && val != t1[i] {
			return false
		}
		if !ok && ok1 {
			return false
		}
		map1[s1[i]] = t1[i]
		map2[t1[i]] = t1[i]
	}
	return true
}

func lemonadeChange(bills []int) bool {
	var x, y uint16
	for _, i := range bills {
		switch i {
		case 5:
			x++
		case 10:
			if x < 1 {
				return false
			}
			x--
			y++
		case 20:
			if y > 0 {
				if x < 1 {
					return false
				}
				x--
				y--
			} else {
				if x < 3 {
					return false
				}
				x -= 3
			}
		}
	}
	return true
}

type Node struct {
	Val      int
	Children []*Node
}

// func postorder(root *Node) []int {

// }

func maxDepth(s string) int {
	count := len(s)
	var (
		depth, temp, check int
	)
	for i := 0; i < count; i++ {
		if s[i] == '(' && check == 0 {
			depth++
			check = 0
		} else {
			if s[i] == ')' {
				check++
			}
		}
		if check > 1 && check > temp {
			temp = check
			if temp >= depth-1 {
				depth = temp + 1
			}
		}
	}
	return depth
}

func largestNumber(nums []int) string {
	// Chuyển các số thành chuỗi để dễ dàng thao tác
	numStrs := make([]string, len(nums))
	for i, num := range nums {
		numStrs[i] = strconv.Itoa(num)
	}

	// Sắp xếp các chuỗi theo tiêu chí ghép hai số lại và so sánh
	sort.Slice(numStrs, func(i, j int) bool {
		return numStrs[i]+numStrs[j] > numStrs[j]+numStrs[i]
	})

	// Nếu số lớn nhất là 0 thì kết quả là "0"
	if numStrs[0] == "0" {
		return "0"
	}

	// Nối các chuỗi lại thành kết quả
	return strings.Join(numStrs, "")
}

func intersect(nums1 []int, nums2 []int) []int {
	l1, l2 := len(nums1), len(nums2)
	var result []int
	short := l1
	if l1 > l2 {
		short = l2
	}
	map1 := make(map[int]uint16, short)
	if l1 < l2 {
		for i := 0; i < l1; i++ {
			map1[nums1[i]]++
		}
		for j := 0; j < l2; j++ {
			if map1[nums2[j]] > 0 {
				result = append(result, nums2[j])
				map1[nums2[j]]--
			}
		}
	} else {
		for i := 0; i < l2; i++ {
			map1[nums2[i]]++
		}
		for j := 0; j < l1; j++ {
			if map1[nums1[j]] > 0 {
				result = append(result, nums1[j])
				map1[nums1[j]]--
			}
		}
	}
	return result
}

func FindPositions(list []int, x int) (int, int) {
	var i, j int = 0, len(list) - 1
	for i < j {
		if list[i]+list[j] == x {
			return i, j
		}
		if list[i]+list[j] > x {
			j--
		} else if list[i]+list[j] < x {
			i++
		}
	}
	return i, j
}

// func longestCommonPrefix(arr1 []int, arr2 []int) int {
// 	prefixMap := make(map[string]int)

// 	// Step 1: Build the prefix map for arr1
// 	for _, num := range arr1 {
// 		strNum := strconv.Itoa(num)
// 		prefix := ""
// 		for _, ch := range strNum {
// 			prefix += string(ch)
// 			prefixMap[prefix]++
// 		}
// 	}

// 	maxLength := 0

// 	// Step 2: Check for common prefixes in arr2
// 	for _, num := range arr2 {
// 		strNum := strconv.Itoa(num)
// 		prefix := ""
// 		for _, ch := range strNum {
// 			prefix += string(ch)
// 			if _, found := prefixMap[prefix]; found {
// 				if len(prefix) > maxLength {
// 					maxLength = len(prefix)
// 				}
// 			}
// 		}
// 	}

// 	return maxLength
// }

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixMap := make(map[string]int)
	max := 0
	for i := 0; i < len(arr1); i++ {
		str := strconv.Itoa(arr1[i])
		pre := ""
		for j := 0; j < len(str); j++ {
			pre = pre + string(str[j])
			if _, ok := prefixMap[pre]; !ok {
				prefixMap[pre] = j
			}
		}
	}
	for i := 0; i < len(arr2); i++ {
		str := strconv.Itoa(arr2[i])
		pre := ""
		for j := 0; j < len(str); j++ {
			pre = pre + string(str[j])
			if _, ok := prefixMap[pre]; ok {
				if max < j+1 {
					max = j + 1
				}
			}
		}
	}
	return max
}

func minBitFlips(start int, goal int) int {
	xor := start ^ goal
	cn := 0
	for xor > 0 {
		cn += xor & 1
		xor >>= 1
	}
	return cn
}

type MyCalendar struct {
	root *bNode
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start int, end int) bool {
	return tryInsert(&c.root, start, end)
}

func tryInsert(head **bNode, start, end int) bool {
	switch {
	case *head == nil:
		*head = &bNode{Start: start, End: end}
		return true
	case end <= (*head).Start:
		return tryInsert(&((*head).Left), start, end)
	case start >= (*head).End:
		return tryInsert(&((*head).Right), start, end)
	default:
		return false
	}
}

type bNode struct {
	Start int
	End   int
	Left  *bNode
	Right *bNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first := l1
	second := l2

	dummy := ListNode{}
	head := &dummy

	carry := 0

	for first != nil && second != nil {

		current := first.Val + second.Val + carry

		if current >= 10 {
			current = current - 10
			carry = 1
		} else {
			carry = 0
		}

		head.Next = &ListNode{
			Val: current,
		}

		head = head.Next

		first = first.Next
		second = second.Next
	}

	if first != nil {
		for first != nil {
			current := first.Val + carry

			if current >= 10 {
				current = current - 10
				carry = 1
			} else {
				carry = 0
			}

			head.Next = &ListNode{
				Val: current,
			}

			head = head.Next
			first = first.Next
		}
	}

	if second != nil {
		for second != nil {
			current := second.Val + carry

			if current >= 10 {
				current = current - 10
				carry = 1
			} else {
				carry = 0
			}

			head.Next = &ListNode{
				Val: current,
			}

			head = head.Next
			second = second.Next
		}
	}

	if carry == 1 {
		head.Next = &ListNode{
			Val: 1,
		}
	}

	return dummy.Next
}

func canArrange(arr []int, k int) bool {
	freq := make(map[int]uint16)
	lenA := len(arr)
	for i := 0; i < lenA; i++ {
		freq[(((arr[i]%k)+k)%k)]++
	}
	for key, v := range freq {
		if key == 0 {
			if v%2 != 0 {
				return false
			}
			continue
		}
		if v != freq[k-key] {
			return false
		}
	}
	return true
}

// func arrayRankTransform(arr []int) []int {
// 	arrL := len(arr)
// 	m := make(map[int]bool, arrL)
// 	for i := 0; i < arrL; i++ {
// 		m[arr[i]] = true
// 	}
// 	mL := len(m)
// 	sortArr := make([]int, mL)
// 	count := 0
// 	for key := range m {
// 		sortArr[count] = key
// 		count++
// 	}
// 	slices.Sort(sortArr)
// 	for i := 0; i < arrL; i++ {
// 		for j := 0; j < mL; j++ {
// 			if arr[i] == sortArr[j] {
// 				arr[i] = j + 1
// 				break
// 			}
// 		}
// 	}
// 	return arr
// }

func arrayRankTransform(arr []int) []int {
	arr2 := append([]int{}, arr...)
	hashMap := make(map[int]int)
	for _, val := range arr {
		hashMap[val] = 0
	}
	slices.Sort(arr)
	i := 0
	for ind, val := range arr {
		if ind > 0 && val == arr[ind-1] {
		} else {
			i = i + 1
		}
		hashMap[val] = i
	}
	i = 0
	for ind, val := range arr2 {
		arr2[ind] = hashMap[val]
	}
	return arr2
}

func frequencySort(nums []int) []int {
	hashMap := make(map[int]int)
	for _, val := range nums {
		hashMap[val]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if hashMap[nums[i]] == hashMap[nums[j]] {
			return nums[i] > nums[j]
		}
		return hashMap[nums[i]] < hashMap[nums[j]]
	})
	return nums
}

func minSubarray(nums []int, p int) int {
	totalValue := 0
	nLen := len(nums)
	for i := 0; i < nLen; i++ {
		totalValue += nums[i]
	}

	remain := totalValue % p
	if remain == 0 {
		return 0
	}

	hashMap := make(map[int]int, nLen)
	hashMap[0] = -1
	minLength := nLen
	temp := 0
	for i := 0; i < nLen; i++ {
		temp = (temp + nums[i]) % p
		neededRemainder := (temp - remain + p) % p
		if v, ok := hashMap[neededRemainder]; ok {
			if i-v < minLength {
				minLength = i - v
			}
		}
		hashMap[temp] = i
	}

	if minLength == nLen {
		return -1
	}

	return minLength
}

func longestPalindrome(s string) int {
	sLen := len(s)
	result := 0
	hashMap := make(map[byte]int, sLen)
	for i := 0; i < sLen; i++ {
		hashMap[s[i]]++
	}
	for _, v := range hashMap {
		result += (v / 2) * 2
	}
	if result > 0 {
		result++
	}
	return result
}

func dividePlayers(skill []int) int64 {
	skpt := 0
	nLen := len(skill)
	if nLen == 2 {
		return int64(skill[0] * skill[1])
	}
	hashMap := make(map[int]int, nLen)
	for i := 0; i < nLen; i++ {
		skpt += skill[i]
		hashMap[skill[i]]++
	}
	var sOfChem int64
	skpt = int((float64(skpt) / float64(nLen)) * 2)
	for key, v := range hashMap {
		if key*2 == skpt {
			if v%2 != 0 {
				return -1
			}
			sOfChem += int64((key * key) * (v / 2))
			delete(hashMap, key)
			continue
		}
		remain := skpt - key
		val, ok := hashMap[remain]
		if !ok {
			return -1
		}
		if val != v {
			return -1
		}
		sOfChem += int64((key * remain) * v)
		delete(hashMap, remain)
		delete(hashMap, key)
	}
	return sOfChem
}

func dividePlayers2(skill []int) int64 {
	slices.Sort(skill)
	skpt := skill[0] + skill[len(skill)-1]
	var sOfChem int
	for i, j := 0, len(skill)-1; i < j; i, j = i+1, j-1 {
		if skill[i]+skill[j] == skpt {
			sOfChem += (skill[i] * skill[j])
		} else {
			return -1
		}
	}
	return int64(sOfChem)
}

func minLength(s string) int {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 &&
			((s[i] == 'B' && stack[len(stack)-1] == 'A') ||
				(s[i] == 'D' && stack[len(stack)-1] == 'C')) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack)
}

func minSwaps(s string) int {
	if len(s) < 2 {
		return 0
	}

	stack := make([]byte, 0, len(s))
	swaps := 0
	unmatched := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			stack = append(stack, '[')
		} else if s[i] == ']' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				unmatched++
			}
		}
	}

	unmatchedPairs := (len(stack) + unmatched) / 2

	swaps = (unmatchedPairs + 1) / 2

	return swaps
}

func minAddToMakeValid(s string) int {
	open, close := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			open++
		} else if s[i] == ')' {
			if open > 0 {
				open--
			} else {
				close++
			}
		}
	}
	return open + close
}

func maxWidthRamp(nums []int) int {
	stack := []int{}

	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}
	}

	answer := 0

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			answer = max(answer, i-stack[len(stack)-1])
			stack = stack[:len(stack)-1]

			if answer == i {
				return answer
			}
		}
	}

	return answer
}

// 1942. The Number of the Smallest Unoccupied Chair - med
func smallestChair(times [][]int, targetFriend int) int {
	target := times[targetFriend]
	chairs := make([]int, 0, len(times))
	sort.Slice(times, func(i, j int) bool {
		return times[i][0] < times[j][0]
	})
	if len(times) == 0 {
		return 0
	}
	chairs = append(chairs, times[0][1])
	for i := 1; i < len(times); i++ {
		cl := len(chairs)
		change := false
		for j := 0; j < cl; j++ {
			if chairs[j] <= times[i][0] {
				if times[i][0] == target[0] {
					return j
				}
				chairs[j] = times[i][1]
				change = true
				break
			}
		}
		if !change {
			chairs = append(chairs, times[i][1])
		}
		if times[i][0] == target[0] {
			return len(chairs) - 1
		}
	}

	return 0
}

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	if text1[len(text1)-1] == text2[len(text2)-1] {
		return 1 + longestCommonSubsequence(text1[:len(text1)-1], text2[:len(text2)-1])
	} else {
		return int(math.Max(float64(longestCommonSubsequence(text1[:len(text1)-1], text2)), float64(longestCommonSubsequence(text1, text2[:len(text2)-1]))))
	}
}

func longestCommonSubsequenceDp(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)

	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}
	return dp[m][n]
}

// 2275. Largest Combination With Bitwise AND Greater Than Zero
func largestCombination(candidates []int) int {
	bitCount := make([]int, 24)
	for i := 0; i < len(candidates); i++ {
		for j := 0; (1 << j) <= candidates[i]; j++ {
			if candidates[i]&(1<<j) != 0 {
				bitCount[j]++
			}
		}
	}
	max := bitCount[0]
	for i := 1; i < 24; i++ {
		if bitCount[i] > max {
			max = bitCount[i]
		}
	}
	return max
}

// 1829. Maximum XOR for Each Query

func getMaximumXor(nums []int, maximumBit int) []int {
	cur, nl := 0, len(nums)
	result := make([]int, nl)
	max := (1 << maximumBit) - 1 //left shift operator equivalent to math.Pow(2, maximumBit)
	for i := 0; i < nl; i++ {
		cur ^= nums[i] // remove condition (0 <= nums[i] < 2maximumBit) then need limit with cur ^= (nums[i] & max)
		result[nl-i-1] = max ^ cur
	}
	return result
}

// 2070. Most Beautiful Item for Each Query
func maximumBeauty(items [][]int, queries []int) []int {
	result := make([]int, len(queries))
	type IndexedValue struct {
		Value int
		Index int
	}
	indexedValues := make([]IndexedValue, len(queries))
	for i, value := range queries { // put data into struct array and sort
		indexedValues[i] = IndexedValue{Value: value, Index: i}
	}

	sort.Slice(indexedValues, func(i, j int) bool { // sort ascending
		return indexedValues[i].Value < indexedValues[j].Value
	})

	sort.Slice(items, func(i, j int) bool { // sort ascending items
		return items[i][0] < items[j][0]
	})
	for i, j := 0, 0; i < len(queries); i++ {
		maxb := 0
		if i > 0 {
			maxb = result[indexedValues[i-1].Index]
		}
		for ; j < len(items); j++ {
			if indexedValues[i].Value >= items[j][0] {
				maxb = max(maxb, items[j][1])
			} else {
				break
			}
		}
		result[indexedValues[i].Index] = maxb
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Định nghĩa MaxHeap
type MaxHeapBt []int

func (h MaxHeapBt) Len() int           { return len(h) }
func (h MaxHeapBt) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeapBt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeapBt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeapBt) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func maximumBeautyHeap(items [][]int, queries []int) []int {
	// Sắp xếp items theo price tăng dần
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	// Tạo một slice để lưu kết quả
	result := make([]int, len(queries))

	// Tạo một slice các query với index
	indexedQueries := make([][2]int, len(queries))
	for i, q := range queries {
		indexedQueries[i] = [2]int{q, i}
	}

	// Sắp xếp queries theo giá trị tăng dần
	sort.Slice(indexedQueries, func(i, j int) bool {
		return indexedQueries[i][0] < indexedQueries[j][0]
	})

	// Khởi tạo max heap
	h := &MaxHeapBt{}
	heap.Init(h)

	itemIndex := 0

	// Xử lý từng query
	for _, query := range indexedQueries {
		// Thêm tất cả items có giá <= query[0] vào heap
		for itemIndex < len(items) && items[itemIndex][0] <= query[0] {
			heap.Push(h, items[itemIndex][1]) // Thêm beauty value
			itemIndex++
		}
		// Heap tự sắp xếp theo giá trị khởi tạo (maxHeap) => phần tử đầu tiên là lớn nhất

		// Nếu heap không rỗng, lấy giá trị beauty lớn nhất
		if h.Len() > 0 {
			result[query[1]] = (*h)[0] // Lấy giá trị lớn nhất từ MaxHeapBt
		} else {
			result[query[1]] = 0
		}
	}

	return result
}

// 1574. Shortest Subarray to be Removed to Make Array Sorted
func findLengthOfShortestSubarray(arr []int) int {
	i, j, n := 0, len(arr)-1, len(arr)

	for i+1 < n && arr[i] <= arr[i+1] {
		i++
	}
	if i == n-1 {
		return 0
	}
	for j-1 >= i && arr[j] >= arr[j-1] {
		j--
	}
	result := min(n-i-1, j)

	left, right := 0, j
	for left <= i {
		for right < n && arr[right] < arr[left] {
			right++
		}
		result = min(result, right-left-1)
		left++
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
