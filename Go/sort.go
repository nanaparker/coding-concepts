package main

import "fmt"

// --- Variables ---
var array1 = []int{6,1,4,2,9,10}
var array2 = []int{6,4,1,2,9,10}


// --- Main Function ---
func main(){

	bubble(array1)

	fmt.Println("\n---Merge Sort---")
	fmt.Println("Sorted Array:",sort(array2))

}



// --- Sorting Algorithms ---
// Bubble Sort
func swap(array []int, a int, b int){
	temp := array[a]
	array[a] = array[b]
	array[b] = temp
}

func bubble(array []int){
	for i:=0; i < len(array)-1; i++ {
		for j:=0; j < len(array)-1; j++{
			if array[j] > array[j+1]{
				swap(array, j, j+1)
			} 
		}
	}
	fmt.Println("---Bubble Sort---")
	fmt.Println("Sorted Array:", array)
}



// Merge Sort
func sort(arr []int) []int{
	if len(arr) == 1{
		return arr
	}

	length := len(arr)/2
	ArrayOne := arr[:length]
	ArrayTwo := arr[length:]

	ArrayOne = sort(ArrayOne)
	ArrayTwo = sort(ArrayTwo)

	return merge(ArrayOne, ArrayTwo)
}

func merge(arrA []int, arrB []int) []int{
	var mixedArr []int

	for len(arrA) > 0 && len(arrB) > 0 {
		if arrA[0] > arrB[0]{
			mixedArr = append(mixedArr, arrB[0])
			arrB = arrB[1:]
		} else {
			mixedArr = append(mixedArr, arrA[0])
			arrA = arrA[1:]
		}
	}

	for (len(arrA) != 0){
		mixedArr = append(mixedArr, arrA[0])
		arrA = arrA[1:]
	}
	for (len(arrB) != 0){
		mixedArr = append(mixedArr, arrB[0])
		arrB = arrB[1:]
	}
	return mixedArr
}