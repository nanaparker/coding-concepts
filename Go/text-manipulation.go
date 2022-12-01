package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main(){
	text := "Text Manipulation"

	// Character Count, First Three Letters, Second Word
	fmt.Println("---Task 1: Substrings---")
	fmt.Println("Character Count:",len(text))
	fmt.Println("First 3 Letters:", text[:3])

	arrayz := strings.Split(text, " ")
	if len(arrayz) >= 2 {
		fmt.Println("Second Word:", arrayz[1])
	} else {
		fmt.Println("Second Word: No second word available")
	}



	// Converting a phrase/word into lower case, UPPER CASE, and MiXeD CaSe
	fmt.Println("\n---Task 2: Character Cases---")
	fmt.Println("Lower Case:", strings.ToLower(text))
	fmt.Println("Upper Case:", strings.ToUpper(text))

	var mixedText string
	for i:=0; i<=len(text)-1; i++ {
		if i % 2 == 0 {
			mixedText = mixedText+strings.ToUpper(string(text[i]))
		} else {
			mixedText = mixedText+strings.ToLower(string(text[i]))
		}
	}
	fmt.Println("Mixed Case:", mixedText)
	


	// Encrypting a phrase using the Caesar Cipher
	fmt.Println("\n---Task 3: Ciphers---")
	shiftCipher := 3
	var cipherText string

	for i:=0;i<=len(text)-1;i++ {

		letter := text[i]+byte(shiftCipher)

		// If the character is " ", add it without shifting
		if (text[i] == 32){
			cipherText = cipherText + string(text[i])
		} else {
			// Check if character is in the Lower case or Upper Case
			if unicode.IsLower( (rune(text[i])) ){
				// a = 97 & z = 122
				// If after shifting, the ASCII value is greater than 122
				// use the formula: (value mod 97 mod 26) + 97 to return to the beginning
				if (letter) > 122{
					cipherText = cipherText + string( (letter % 97 % 26) + 97 )
				} else{
					cipherText = cipherText + string(letter)
				}
			} else {
				// A = 65 & Z = 90
				// If after shifting, the ASCII value is greater than 90
				// use the formula: (value mod 65 mod 26) + 65 to return to the beginning
				if (letter) > 90{
					cipherText = cipherText + string( (letter % 65 % 26) + 65 )
				} else{
					cipherText = cipherText + string(letter)
				}
			}
		}
	}

	fmt.Println("Shift Cipher: ", shiftCipher, "\nCipher Text: ", cipherText)
}