package main

import "fmt"

// Printing from 0 to 4 and passing the values
// through a channel
func numbers(num chan int){
	for x:=0; x<5; x++ {
		num <- x
	}
	close(num)
}


// Printing from a to e and passing the letters
// through a channel
func letters(alpha chan string){
	for x:=0; x<5; x++ {
		alpha <- string(rune(97+x))
	}

	close(alpha)
}


func main(){
	// Setting up channels
	num := make(chan int)
	alpha := make(chan string)

	// Executing concurrent functions
	go numbers(num)
	go letters(alpha)

	// Printing received values in pairs as and when they arrive
	for {
		number, number_bool := <-num
		alphabet, alphabet_bool := <-alpha

		if alphabet_bool == false && number_bool == false {
			break
		}
	
		fmt.Println("Received", number, alphabet)
	}
}