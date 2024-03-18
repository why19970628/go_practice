package main

func hammingWeight(num uint32) int {
	resp := uint32(0)

	for num != 0 {
		resp += num % uint32(2)
		num = num / 2
	}

	return int(resp)
}
