package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println("Prime num:", primeNumbers(100))
}

func Swapbits(i, j uint, n uint64) uint64 {
	x := (n >> i) ^ (n>>j)&0x1
	return (n ^ (x<<i | x<<j))
}

const MaxUint = ^uint64(0)

//Implement function , which implement string into Int
func stringToInf(num string) (uint64, error) {
	var number uint64
	for _, s := range num {
		r := rune(s)
		if r < '0' || r > '9' {
			return 0, errors.New("given string Can't be converted to Int")
		}
		if MaxUint-uint64(r) < number {
			return 0, errors.New("Can't convert string , Number is going to overflow")
		}
		number = number*10 + uint64(r)
	}
	return number, nil
}

//Find GCD of two Number
func GCD(x, y int) int {

	if x == y {
		return x
	}
	if (x&0x1) == 0 && (y&0x1) == 0 { // if x and y both are Even then call its recursive and multiply them
		return (GCD(x>>1, y>>1) << 1)
	} else if (x&0x1) == 0 && (y&1) == 1 { //if x is even and y is odd
		return (GCD(x>>1, y))
	} else if (x&0x1) == 1 && (y&1) == 0 { //if x is even and y is odd
		return (GCD(x, y>>1))
	} else if x > y { //if both are odd and x > y
		return (GCD(x-y, x))
	}
	return (GCD(x, y-x))

}

//FInd all Prime number between 1 to given n number

func primeNumbers(n int) []int {
	var isPrime []bool = make([]bool, n+1)
	var list = make([]int, 0)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false
	for i := 2; i < n; i++ {
		if isPrime[i] {
			list = append(list, i)
			p := i
			for j := 2 * p; j < n; j += p {
				isPrime[j] = false
			}
		}
	}
	return list
}
