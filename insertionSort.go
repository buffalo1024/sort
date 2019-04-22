package main

import (
   "fmt"
   "time"
)

func insertionSort(s []int) {
   t := time.Now()
   var a = make([]int, len(s))
   copy(a, s)
   for i := 1; i < len(a); i++ {
      m := a[i]
      for j := i; j > 0; j-- {
         if a[j-1] > a[i] && (j-2 < 0 || a[j-2] <= a[i]) { 
            copy(a[j:i+1], a[j-1:i])
            a[j-1] = m
         }
      }
   }
   fmt.Print("\n", "Figures insertionSorted:")
   printSlice(a)
   fmt.Print("\n",time.Since(t))
}