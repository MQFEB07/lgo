package lgo

import (
	"fmt"
	"testing"
)

func areListsEqual(list1, list2 *ListNode) bool {
	for list1 != nil && list2 != nil {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}

	return list1 == nil && list2 == nil
}
func TestMerge2SortList(t *testing.T) {
	listExpect := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}
	// list1 := ListNode{}
	list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	// list2 := ListNode{
	// 	Val: 1,
	// }
	list2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	listMerge := mergeTwoLists(&list1, &list2)
	if !areListsEqual(&listExpect, listMerge) {
		t.Fail()
	}
}

func TestReverseList(t *testing.T) {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	listExpect := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}
	result := reverseList(&head)
	if !areListsEqual(&listExpect, result) {
		t.Fatalf("Expected: %v, but got: %v", &listExpect, result)
	}
}

func TestSnakeMatrix(t *testing.T) {
	command := []string{"DOWN", "RIGHT", "UP"}
	n := 3
	expect := 1
	result := finalPositionOfSnake(n, command)
	if result != expect {
		t.Fail()
	}
}

func TestPalindrome(t *testing.T) {
	list := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	var expect bool = true
	result := isPalindrome(&list)
	if result != expect {
		t.Fail()
	}
}

func TestMinSteps(t *testing.T) {
	var input int = 3
	var output int = 3
	result := minSteps(input)
	if result != output {
		t.Fail()
	}
}

func TestIsomorphic(t *testing.T) {
	s1 := "badc"
	s2 := "baba"
	var expect bool = false
	result := isIsomorphic(s1, s2)
	if expect != result {
		t.Fail()
	}
}

func TestLemonadeChange(t *testing.T) {
	bills := []int{5, 5, 5, 5, 20, 20, 5, 5, 5, 5}
	expect := false
	result := lemonadeChange(bills)
	if expect != result {
		t.Fail()
	}
}

func TestLargestNumber(t *testing.T) {
	nums := []int{3, 30, 34, 5, 9}
	expect := "9534330"
	result := largestNumber(nums)
	if expect != result {
		t.Fail()
	}
}

func TestIntersect(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	expect := []int{4, 9}
	result := intersect(nums1, nums2)
	for i := 0; i < len(expect); i++ {
		has := false
		for j := 0; j < len(result); j++ {
			if expect[i] == result[j] {
				has = true
				break
			}
		}
		if has {
			continue
		} else {
			t.Fail()
			break
		}
	}
}

func TestFindPositions(t *testing.T) {
	list := []int{2, 5, 6, 8, 10, 11, 15}
	x := 16
	pos1, pos2 := 1, 5
	res1, res2 := FindPositions(list, x)
	if res1 != pos1 || res2 != pos2 {
		t.Fail()
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	arr1 := []int{13, 27, 45}
	arr2 := []int{21, 27, 48}
	expect := 2
	result := longestCommonPrefix(arr1, arr2)
	if expect != result {
		t.Fail()
	}
}

func TestMinBitFlips(t *testing.T) {
	start := 10
	goal := 7
	expect := 3
	result := minBitFlips(start, goal)
	if expect != result {
		t.Fail()
	}
}

func TestAddTwoNumbers(t *testing.T) {
	listExpect := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
			},
		},
	}
	list1 := ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 7,
		},
	}
	list2 := ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 2,
		},
	}
	listMerge := addTwoNumbers(&list1, &list2)
	if !areListsEqual(&listExpect, listMerge) {
		t.Fail()
	}
}

func TestCanArrange(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
	k := 5
	expect := true
	result := canArrange(arr, k)
	if result != expect {
		t.Fail()
	}
}

func TestArrayRankTransform(t *testing.T) {
	arr1 := []int{2, -11, 24, 15, 26, -31, -48, -49, 22, 37, -36}
	expect := []int{6, 5, 9, 7, 10, 4, 2, 1, 8, 11, 3}
	result := arrayRankTransform(arr1)
	arrLen := len(arr1)
	for i := 0; i < arrLen; i++ {
		if result[i] != expect[i] {
			t.Fail()
			break
		}
	}

}

func TestInorderTraversal(t *testing.T) {
	tree := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	expect := []int{1, 3, 2}
	result := inorderTraversal(&tree)
	for i := 0; i < len(result); i++ {
		if result[i] != expect[i] {
			t.Fail()
			break
		}
	}
}

func TestIsSameTree(t *testing.T) {
	tree := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	tree1 := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	expect := false
	result := isSameTree(&tree, &tree1)
	if result != expect {
		t.Fail()
	}
}

func TestMinSubarray(t *testing.T) {
	nums := []int{6, 3, 5, 2}
	p := 9
	expect := 2
	result := minSubarray(nums, p)
	if result != expect {
		t.Fail()
	}
}

func TestLongestPalindrome(t *testing.T) {
	s := "abcccccdd"
	expect := 7
	result := longestPalindrome(s)
	if result != expect {
		t.Fail()
	}
}

func TestDividePlayers(t *testing.T) {
	sk := []int{2, 3, 4, 2, 5, 5}
	var expect int64 = 32
	result := dividePlayers2(sk)
	if result != expect {
		t.Fail()
	}
}

func TestMinLength(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"Basic case", "ABFCACDB", 2},
		{"No reduction", "ACBD", 4},
		{"All reducible", "ABCD", 0},
		{"Empty string", "", 0},
		{"Single reduction", "ABD", 1},
		{"Multiple reductions", "ABABCACD", 2},
		{"Alternating pairs", "ABCDABCD", 0},
		{"Long string", "ABCDBCABCDBCABCDBCABCDBC", 8},
		{"Error case", "AATQCABDCBE", 7},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := minLength(tc.input)
			if result != tc.expected {
				t.Errorf("minLength(%q) = %d; want %d", tc.input, result, tc.expected)
			}
		})
	}
}

func TestMinLengthMultipleReductions(t *testing.T) {
	input := "ABCDBCABCDBCABCDBCABCDBC"
	expected := 8

	fmt.Println("Testing input:", input)

	result := minLength(input)

	if result != expected {
		t.Errorf("minLength(%q) = %d; want %d", input, result, expected)
	}
}

func TestMinSwaps(t *testing.T) {
	s := "]]][[["
	expected := 2
	result := minSwaps(s)
	if expected != result {
		t.Errorf("minLength(%q) = %d; want %d", s, result, expected)
	}
}

func TestMinAddToMakeValid(t *testing.T) {
	s := "())"
	expected := 1
	result := minAddToMakeValid(s)
	if expected != result {
		t.Errorf("minLength(%q) = %d; want %d", s, result, expected)
	}
}
func TestMaxWidthRamp(t *testing.T) {
	nums := []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}
	expected := 7
	result := maxWidthRamp(nums)
	if expected != result {
		t.Errorf("minLength(%q) = %d; want %d", nums, result, expected)
	}
}

func TestSmallestChair(t *testing.T) {
	times := [][]int{{3, 10}, {1, 5}, {2, 6}}
	targetFriend := 0
	expected := 2
	result := smallestChair(times, targetFriend)
	if expected != result {
		t.Errorf("result = %d; want %d", result, expected)
	}
}

func TestLongestDiverseString(t *testing.T) {
	a, b, c := 1, 1, 7
	expected := "ccaccbcc"
	result := longestDiverseString(a, b, c)
	if expected != result {
		t.Errorf("result = %v; want %v", result, expected)
	}
}

func TestKthLargestLevelSum(t *testing.T) {
	// Tạo cây nhị phân để test
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	tests := []struct {
		k        int
		expected int64
	}{
		{k: 1, expected: 22},
		{k: 2, expected: 5},
		{k: 3, expected: 1},
		{k: 4, expected: -1},
	}

	for _, test := range tests {
		result := kthLargestLevelSum(root, test.k)
		if result != test.expected {
			t.Errorf("For k=%d, expected %d but got %d", test.k, test.expected, result)
		}
	}
}

func TestMaxUniqueSplit(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"ababccc", 5},
		{"aba", 2},
		{"aa", 1},
		{"abc", 3},
		{"", 0},
		{"abcdef", 6},
	}

	for _, test := range tests {
		result := maxUniqueSplit(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}
