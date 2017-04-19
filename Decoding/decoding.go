package main

import "fmt"

func main() {
	var val = 0x7
	var digit []byte
	val = flipBits(val)
	fmt.Println("flip val:%d\n", val)

	fmt.Println("Enter digit:")
	fmt.Scanln(&digit)
	fmt.Printf("type %t and %v\n", digit, digit)
	fmt.Println("deconding cnt:", decoding(digit, len(digit)))
	fmt.Println("deconding cnt:", decondingDp(digit))
}
func flipBits(num int) int {
	var mask = 0x11
	var fRslt int
	for i := 0; i < 31; i = i + 2 {
		bits := num & 0x11
		num = num >> 2
		if bits != 0x11 && bits != 0x0 {
			bits &^= mask
		}
		fRslt |= bits << uint(i)
	}
	return fRslt
}

func decoding(digist []byte, n int) int {
	var count int

	if n == 0 || n == 1 {
		return 1
	}

	if digist[n-1] > '0' {
		count = decoding(digist, n-1)
	}
	if digist[n-2] < '2' || (digist[n-2] == '2' && digist[n-1] < '7') {
		count += decoding(digist, n-2)
	}
	return count
}
func decondingDp(digit []byte) int {
	cnt := make([]int, len(digit)+1)
	fmt.Println(cnt, len(cnt))
	cnt[0] = 1
	cnt[1] = 1
	for i := 2; i <= len(digit); i++ {
		if digit[i-1] > '0' {
			cnt[i] = cnt[i-1]
		}
		if digit[i-2] < '2' || (digit[i-2] == '2' && digit[i-1] < '7') {
			cnt[i] += cnt[i-2]
		}

	}
	return cnt[len(digit)]
}
