package actions

// Checks if a number is prime or not
func IsPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func CheckTheChange() string {
	return "there is an change"
}
func CheckTheChang1e() string {
	return "there is an change"
}
