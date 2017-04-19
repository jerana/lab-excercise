package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ipAddr string
	fmt.Println("vim-go")
	fmt.Printf("Enter IP number which you want:\n")
	fmt.Scanln(&ipAddr)
	fmt.Println("Scan IP add:", ipAddr)
	fmt.Println("Correct Ip addr:", convertValidIpAddress([]byte(ipAddr)))
}

func convertStringToInt(str string) int {
	var isSign bool = false
	var s []byte
	s = []byte(str)
	if s[0] == '-' {
		isSign = true
		s = s[1:]
	}
	var sum int
	for _, c := range s {
		num := c - '0'
		sum = 10*sum + int(num)
	}
	if isSign {
		sum = -1 * sum
	}
	return sum
}
func converIntToString(num int) {
	var isSign = false
	if num < 0 {
		isSign = true
	}
	num = -1 * num
	var s []byte
	for num > 0 {
		c := '0' + num%10
		s = append(s, byte(c))
		num /= 10
	}
	if isSign {
		s = append(s, '-')
	}
}

//Look and say Problem

func lookandSay(num int) []byte {
	var s []byte
	s = []byte("1")
	for i := 1; i < num; i++ {
		s = convert(s)
	}
	return s
}
func convert(str []byte) []byte {
	count := 1
	var b []byte
	for i := 0; i < len(str)-1; i++ {
		for i+1 < len(str)-1 && str[i] == str[i+1] {
			count++
			i++
		}
		b = append(b, str[i])
		b = append(b, []byte(strconv.Itoa(count))...)
		//[]byte(strconv.Itoa(count)), str[i])
	}
	return b
}

//Make Valid IP Address from mingled string
func convertValidIpAddress(IPaddr []byte) []string {
	var result []string
	//We can put period at 3 places in string to make it Valid Ip Address.
	//Will Put Period and check put other period other places make it Valid Ip address
	//For Every substring, Make String check routine to validate if Substring can make valid Ip address
	// Substring valid region is 1-255

	for i := 1; i < 4 && i < len(IPaddr); i++ {
		a := IPaddr[0:i]
		if isValidIp(a) {
			for j := 1; i+j < len(IPaddr) && j < 4; j++ {
				b := IPaddr[i:j]
				if isValidIp(b) {
					for k := 1; (i+j+k) < len(IPaddr) && k < 4; k++ {
						c := IPaddr[i+j : k]
						d := IPaddr[i+j+k:]
						if isValidIp(c) && isValidIp(d) {
							var finalStr []byte
							finalStr = append(finalStr, a...)
							finalStr = append(finalStr[len(a):], b...)
							finalStr = append(finalStr[len(a)+len(b):], c...)
							finalStr = append(finalStr[len(a)+len(b)+len(c):], d...)
							result = append(result, string(finalStr))

						}
					}
				}
			}
		}

	}
	return (result)
}

func isValidIp(s []byte) bool {
	if len(s) > 3 {
		return false
	}
	//"00","01","000" are also not valid string
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	val, _ := strconv.Atoi(string(s))
	if val > 255 || val < 0 {
		return false
	}
	return true

}

func Encoding(s string) string {
	var rslt []byte
	count := 1
	for i := 1; i < len(s); i++ {
		if i == len(s)-1 || s[i] != s[i-1] {
			tmp := strconv.Itoa(count) + s[i-1]
			count = 1
		} else {
			count++
		}
	}
	return string(rslt)
}
func Decoding(s string) string {
	var rslt []byte
	count := 0
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			count = 10*count + s[i] - '0'
		} else {
			for count > 0 {
				rslt = append(rslt, s[i])
				count--
			}

		}
	}
	return rslt
}

func rollingHashFactor(s []byte) int {
	var power_s = 1
	for i := 1; i < len(s); i++ {
		power_s = power_s * kBase
	}
	return power_s

}

func Hash(s []byte) int {
	for i := 0; i < len(s); i++ {
		t_hash = kBase*t_hash + t[i]
	}
	return t_hash
}
func rollingHash(oldHash, s string, oldIdx, newIdx int) int {
	oldHash -= s[oldIdx] * power_s
	oldHash += kBase*oldHash + s[newIdx]
	return oldHash
}

//Robin-Karp Algo for String searching at
func stringSearch(s, t string) int {
	var t_hash, s_hash int
	var kBase = 26
	var power_s = 0
	t_hash = Hash(t[0:len(s)])
	s_hash = Hash(s)
	for i := len(s); i < len(t); i++ {
		if t_hash == s_hash && compare(t[i-len(s):i], s) {
			return i - len(s)
		}
		t_hash = rollingHash(t_hash, t, i-len(s), i+1)
	}
	if t_hash == s_hash && compare(t[len(t)-len(s):], s) {
		return len(t) - len(s)
	}
	return -1
}

//Two Sum Problem
//Given Array, Find Two array index which Sum equal to given target
//Use Hash Map : it Take o(n) space
//Use Value to index map

func findTargetSum(arr []int, target int) []int {
	var ret []int
	m := make(map[int]int)
	for i := 0; i < len(arr)-1; i++ {
		idx, ok := m[target-arr[i]]
		if ok {
			ret = append(ret, idx+1, i+1)
		}
		m[arr[i]] = i
	}
}
//Brute force approach
//Condition 1: return 0 if target string is null ; return -1 if source is is null , 
//Condition 2: if target and soruce are equal  make sure it -1 on match
func strstr(s, t string) int {
	var i , j int 
	for i =0; i++ {
		for j =0; j++ {
			if j == len(t) {
				return i
			}
			if i + j == len(s) { //String can't be found at source
				return -1
			} 
			if s[i] != s[j] {
				break
			}
		}
	}
	return -1

}
