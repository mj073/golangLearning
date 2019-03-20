package main

func main(){
	n := 8
	if isPowerOfTwo(n) {
		println("power of two")
	}
}
func isPowerOfTwo(n int) bool{
	for n != 1{
		if n % 2 != 0{
			return false
		}
		n = n / 2
	}
	return true
}