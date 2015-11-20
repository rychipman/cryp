package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	if len(os.Args) > 2 {
		fmt.Println("Too many command-line arguments.")
		os.Exit(1)
	}
	if len(input) != 20 {
		fmt.Println("Input needs to be 20 digits")
		os.Exit(1)
	}
	fmt.Printf("Input: %s\n", input)
	fmt.Println("Generating key...")
	firstTwenty := input
	lastTwenty := generateFrom(firstTwenty)
	fmt.Printf("Genrated Digits: %s\n", lastTwenty)
	key := firstTwenty + lastTwenty
	fmt.Printf("Key: %s\n", key)
}

func generateFrom(firstTwenty string) string {
	cipher := "last save file" + firstTwenty
	unicodeBytes := []byte(cipher) // TODO: make sure this is unicode
	digest := md5.Sum(unicodeBytes)
	lastTwenty := ""
	for k := 6; k < 16; k++ {
		uInt := digest[k] & 0xFF
		s := fmt.Sprintf("%x", uInt)
		if len(s) == 1 {
			s = "0" + s
		}
		lastTwenty += s
	}
	return lastTwenty
}
