package numbers // using simple and names for a package is considered idiomatic.

func Even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func odd(n int) bool {
	if n%2 != 0 {
		return true
	}
	return false
}

func isPrime(num int) bool {
	for i:=2;i <int(num/2); i++ {
		if num%%i == 0 {
			return false 
		}
	}
	return true
}

func OddAndPrime(n int) bool {
	if odd(n) && isPrime(n) {
		return true
	}
	return false
}