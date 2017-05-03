package arithmetic

const (
	UPPER_BOUND_I64 = 9223372036854775807  // 2^63 - 1
	LOWER_BOUND_I64 = -9223372036854775808 // - 2^63
)

// ----------  int64  ---------- //
func MaxI64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MinI64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MaxI64S(nums ...int64) int64 {
	var mx int64 = LOWER_BOUND_I64
	for _, num := range nums {
		if num > mx {
			mx = num
		}
	}
	return mx
}

func MinI64S(nums ...int64) int64 {
	var mn int64 = UPPER_BOUND_I64
	for _, num := range nums {
		if num < mn {
			mn = num
		}
	}
	return mn
}

// ----------  int  ---------- //
func MaxI(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinI(a, b int) int {
	if a < b {
		return a
	}
	return b
}
