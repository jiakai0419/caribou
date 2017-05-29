package arithmetic

const (
	UPPER_BOUND_I64 = 9223372036854775807  // 2^63 - 1
	LOWER_BOUND_I64 = -9223372036854775808 // - 2^63

	UPPER_BOUND_I = 2147483647  // 2^31 - 1
	LOWER_BOUND_I = -2147483648 // 2^31 - 1
)

func MaxI64(nums ...int64) int64 {
	var mx int64 = LOWER_BOUND_I64
	for _, num := range nums {
		if num > mx {
			mx = num
		}
	}
	return mx
}

func MinI64(nums ...int64) int64 {
	var mn int64 = UPPER_BOUND_I64
	for _, num := range nums {
		if num < mn {
			mn = num
		}
	}
	return mn
}

func MaxI(nums ...int) int {
	var mx int = LOWER_BOUND_I
	for _, num := range nums {
		if num > mx {
			mx = num
		}
	}
	return mx
}

func MinI(nums ...int) int {
	var mn int = UPPER_BOUND_I
	for _, num := range nums {
		if num < mn {
			mn = num
		}
	}
	return mn
}
