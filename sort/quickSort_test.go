package main

import (
	"testing"
)

func TestQuickSort(t *testing.T) {

	b := []int{}
	for i := 0; i <100000; i ++ {
		b = append(b,i)
	}
	QuickSort(b)
	//fmt.Println(c)

	for i := 0; i <len(b); i ++ {
		for j := i +1; j <len(b); j ++ {
			if b[i] > b[j] {
				t.Errorf("wrone")
			}
		}
	}
}

