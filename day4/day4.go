package day4

func AllIncreasing(number int) bool {
	for prev := number % 10; number > 0; {
		number /= 10
		remainder := number % 10
		if remainder > prev {
			return false
		}
		prev = remainder
	}
	return true
}

func HasPairedNumbers(number int) bool {
	for prev := number % 10; number > 0; {
		number /= 10
		remainder := number % 10
		if remainder == prev {
			return true
		}
		prev = remainder
	}
	return false
}

func HasAnEvenPair(number int) bool {
	nums := map[int]int{}
	for number > 0 {
		remainder := number % 10
		nums[remainder]++
		number /= 10
	}

	for _, v := range nums {
		if v == 2 {
			return true
		}
	}

	return false
}
