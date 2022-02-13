package actions

// Checks if a number is prime or not
func IsNotPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i != 0 {
			return false
		}
	}
	return true
}

func test(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i != 0 {
			return false
		}
	}
	return true
}
func test2(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i != 0 {
			return false
		}
	}
	return true
}
