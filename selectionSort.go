package main

import (
   "fmt"
   "time"
)

func selectionSort(s []int) {
   t := time.Now()
   var a = make([]int, len(s))
   copy(a, s)
   i, j, k := 0, 0, 0
   for ; i < len(a); i++ {
      for j, k = i, i; j < len(a); j++ {
         if a[j] < a[k] {
            k = j
         }
      }
      a[i], a[k] = a[k], a[i]
   }
   fmt.Print("\n","Figures selectionSorted:")
   printSlice(a)
   fmt.Print("\n", time.Since(t))
}