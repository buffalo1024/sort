package main

import "fmt"

func printSlice(a []int) {
   fmt.Print("\n")
   for _, v := range a {
      fmt.Print(v, " ")
   }
}