package main

import (
   "math/rand"
)

func generate(n int, a []int) []int {
   for i := 0; i < n; i++ {
      a = append(a, rand.Intn(100))
   }
   return a
}