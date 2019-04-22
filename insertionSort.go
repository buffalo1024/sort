package main

import (
   "fmt"
)

func insertionSort(s []int) {
   var a = make([]int, len(s))
   copy(a, s)
   for i := 1; i < len(a); i++ {
      m := a[i]
      for j := i; j > 0; j-- {
         if a[j-1] > m {
            a[j] = a[j-1]
         }
         if a[j-1] <= m {
            a[j] = m
            break
         }
      }
   }
   fmt.Print("\n", "Figures insertionSorted:")
   printSlice(a)
}