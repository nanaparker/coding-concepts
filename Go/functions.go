package main

import (
	"fmt"
	"unicode"
)

func hello(){
	fmt.Println("Hello World")
}

func percent(a float32){
	var val float32 = a/100
	fmt.Println("Percentage: ", val,"%")
}

func caesar(shift int, text string) string {
	var cipherText, shiftedText string
	
	for i:=0; i<=len(text)-1; i++ {
		if text[i] == 32 {
			cipherText = cipherText + string(text[i])
		} else {
			if unicode.IsLower(rune(text[i])){
				if ((text[i] + byte(shift)) > 122){
					shiftedText = string((text[i] + byte(shift)) % 97 % 26 + 97)
					cipherText = cipherText + shiftedText
				} else{
					cipherText = cipherText + string(text[i] + byte(shift))
				}				
			} else {
				if ((text[i] + byte(shift)) > 90){
					shiftedText = string((text[i] + byte(shift)) % 65 % 26 + 65)
					cipherText = cipherText + shiftedText
				} else{
					cipherText = cipherText + string(text[i] + byte(shift))
				}	
			}
		}
	}
	return cipherText
}

func main(){
	fmt.Println("---Task 1: Hello World---")
	hello()
	fmt.Println("\n---Task 2: Percentages---")
	percent(4)
	fmt.Println("\n---Task 3: Caesar Cipher---")
	cipher := caesar(3, "Text Manipulation")
	fmt.Println("Cipher Text: ", cipher)
}