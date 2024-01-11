package main

import "fmt"

func main()  {
	
	arr := [5] int{12, 34, 14, 56, 99}
	reversedArray := arrayReverser(arr[:])

	fmt.Println(reversedArray)

}

func arrayReverser(array[] int) []int{

	length := len(array)
	reversed := make([]int, 0, length)

	for i := length-1; i >=0; i--{
		reversed = append(reversed, array[i])
	}

	return reversed[:]
}
