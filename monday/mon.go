package monday

// 1652. Defuse the Bomb - 18/11/24
func decrypt(code []int, k int) []int {
	dec, cl := make([]int, len(code)), len(code)
	if k == 0 {
		return dec
	}
	end := k
	if k < 0 { // sliding window
		for i := cl + k; i < cl; i++ {
			dec[0] += code[i]
		}
		end = cl + k
	} else {
		for i := 1; i <= k; i++ {
			dec[0] += code[i]
		}
	}
	for i := 1; i < cl; i++ {
		if k < 0 {
			dec[i] = dec[i-1] + code[i-1] - code[end]
		}
		end++
		if end >= cl {
			end = 0
		}
		if k > 0 {
			dec[i] = dec[i-1] - code[i] + code[end]
		}
	}
	return dec
}
