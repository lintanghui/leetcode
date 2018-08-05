package divideTwoIntegers

func divide(dividend int, divisor int) int {
	var flag bool
	if dividend < 0 && divisor < 0 || (dividend > 0 && divisor > 0) {
		flag = true
	}
	if dividend < 0 {
		dividend = 0 - dividend
	}
	if divisor < 0 {
		divisor = 0 - divisor

	}
	var result int
	for dividend >= divisor {
		dividend -= divisor
		result++
	}
	if !flag {
		result = 0 - result
	}
	if result > 2147483647 {
		result = 2147483647
	}
	return result
}
