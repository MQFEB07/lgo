package wednesday

import "math"

// 2516. Take K of Each Character From Left and Right
func takeCharacters(s string, k int) int {
	freq := make([]int, 3)
	var n int = len(s)
	var ans int = math.MaxInt
	for _, ch := range s {
		freq[ch-'a']++
	}
	if freq[0] < k || freq[1] < k || freq[2] < k {
		return -1
	}
	for i, j := 0, 0; i < n; i++ {
		for j <= n && freq[0] >= k && freq[1] >= k && freq[2] >= k {
			if i+n-j < ans {
				ans = i + n - j
			}
			if j < n {
				freq[s[j]-'a']--
			}
			j++
		}
		freq[s[i]-'a']++
	}
	return ans
}
