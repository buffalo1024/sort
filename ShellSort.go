package main

import (
	"fmt"
	"time"
)

func ShellSort(s []int) {
	t := time.Now()
	var a []int = make([]int, len(s))
	copy(a, s)
	for k := len(a) / 2; k >= 1; k /= 2 {
		for i := 0; i < len(a); i++ {
			for j := i - k; j >= 0 && a[j+k] < a[j]; j -= k {
				a[j], a[j+k] = a[j+k], a[j]
			}
		}
	}
	fmt.Print("\n","Figures ShellSorted:")
	printSlice(a)
	fmt.Print("\n",time.Since(t))
}
