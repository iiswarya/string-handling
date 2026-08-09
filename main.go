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

	//	4.find duplicate characters in a string
	duplicate := findDuplicate(str3)
	fmt.Println("duplicate characters:", duplicate)

	//5.check if a string is a palindrome
	str4 := "madam"
	isPalindrome := isPalindrome(str4)
	if isPalindrome {
		fmt.Println(str4, "is a palindrome")
	} else {
		fmt.Println(str4, "is not a palindrome")
	}

	//6. character frequency in a string
	str5 := "hello"
	freq := charFrequency(str5)
	fmt.Println("character frequency:", freq)
}

func charFrequency(s string) (freq map[string]int) {

	freq = make(map[string]int)
	for _, ch := range s {
		freq[string(ch)]++
	}
	return freq

}

func isPalindrome(s string) bool {

	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

func findDuplicate(s string) []string {

	duplicates := []string{}
	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}
	for ch, freq := range count {
		if freq > 1 {
			duplicates = append(duplicates, string(ch))
		}
	}
	return duplicates

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
