package main

import (
	"fmt"
	"unicode"
)

var word string = ""

func main(){
	cipher("Text Manipulation", 3)
}


func cipher(message string, shift int){
	index := 0
	length := len(message)

	fmt.Println(caesar(message, index, length, shift))
}

func caesar(text string, index int, length int, shift int) string{
	
	if index == length{
		return string("Cipher Text: " + word)
	} else {
		letter := text[index] + byte(shift)

		if (text[index] == 32){
			word = word + string(text[index])
		} else {
			if unicode.IsLower(rune(text[index])){
				if (letter) > 122{
					word = word + string((letter % 97 % 26) + 97)
				} else{
					word = word + string(letter)
				}
			} else {
				if (letter) > 90{
					word = word + string((letter % 65 % 26) + 65 )
				} else{
					word = word + string(letter)
				}
			}
		}
	
		index++
		return caesar(text, index, length, shift)
	}
}