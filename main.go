package main

import "fmt"

func main() {

	//1.reverse a string
	str := "hello"
	reversed := reverseString(str)
	fmt.Println("reverse string:", reversed) //op:olleh

	//2.check if 2 strings are anagrams
	str1 := "listen"
	str2 := "silent"
	isAnagram := areAnagrams(str1, str2)
	fmt.Println("are anagrams:", isAnagram) //true

	//3.find the first non-repeating character in a string
	str3 := "programming"
	nonRepeating := firstNonRepeating(str3)
	fmt.Println("first non-repeating character:", string(nonRepeating)) //op:p

}

func firstNonRepeating(s string) rune {
	count := make(map[rune]int)

	for _, ch := range s {
		count[ch]++
	}

	for _, ch := range s {
		if count[ch] == 1 {
			return ch
		}
	}
	return 0
}

func areAnagrams(str1, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	count := make(map[rune]int)

	for _, ch := range str1 {
		count[ch]++
	}

	for _, ch := range str2 {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}

func reverseString(str string) string {

	//b:=[]byte(str)  wrong approach - breaks for emojis and multi-byte characters

	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
