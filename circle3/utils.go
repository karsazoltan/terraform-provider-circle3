package provider

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsString(pattern string, array []string) bool {
	for _, i := range array {
		if i == pattern {
			return true
		}
	}
	return false
}

func CheckPowerOfTwo(n int) int {
	if n == 0 {
		return 1
	}
	return n & (n - 1)
}
