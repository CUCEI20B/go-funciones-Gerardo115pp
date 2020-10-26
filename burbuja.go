package main

func Burbuja(s []int64) {
	var n int = len(s)
	var temp int64
	for h := 0; h < n-1; h++ {
		for k := 0; k < (n - h - 1); k++ {
			if s[k] > s[k+1] {
				temp = s[k]
				s[k] = s[k+1]
				s[k+1] = temp
			}
		}

	}
}

func main() {
	// var nums []int64 = []int64{-100, -10, -2, 1, 14, 90}
	// Burbuja(nums)
	// fmt.Println(nums)
}
