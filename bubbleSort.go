package main

import (
   "fmt"
   "time"
)

func bubbleSort(s []int) {
   t := time.Now()
   var a = make([]int, len(s))
   copy(a, s)
   for i := 0; i < len(a); i++ {
      for j := 0; j < len(a)-i-1; j++ {
         if a[j] > a[j+1] {
            a[j], a[j+1] = a[j+1], a[j]
         }
      }
   }
   fmt.Println("\n")
   fmt.Print("Figures bubbleSorted:")
   printSlice(a)
   fmt.Print("\n", time.Since(t))
}