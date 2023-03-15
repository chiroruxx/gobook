package practice2_4

func PopCount(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		value := x >> i
		if value%2 == 1 {
			count++
		}
	}

	return count
}
