package main

import (
   "fmt"
)

var n int

var a []int

func main() {
   fmt.Print("Please imput the number of figures that you want to generate:")
   fmt.Scanln(&n)
   a = generate(n, a)
   fmt.Print("The figures generated that have not been sorted:")
   printSlice(a)
   bubbleSort(a)
   selectionSort(a)
   insertionSort(a)
   ShellSort(a)
}