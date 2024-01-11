package main

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T){
	table := []struct{
		normalArray []int
		reversedArray []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1}},
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{67, 98, 29, 10, 11}, []int {11, 10, 29, 98, 67}},
		{[]int{11, 10, 29, 98, 67}, []int {67, 98, 29, 10, 11}},
		{[]int{-1, 55, -14}, []int{-14, 55, -1}},
	}

	for _, prop := range table{
		reversed := arrayReverser(prop.normalArray)

		if !reflect.DeepEqual(reversed, prop.reversedArray){
			t.Errorf(("wrong result!"))
		}
	}
}
