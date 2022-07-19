package main

import "fmt"

func insertSort(arr []int, n int) []int {
    return arr
}

func printArray(arr []int, n int) {
    for i := 0; i < n; i++ {
        fmt.Println("arr: ", i, arr[i])
    }
}

/**
  * tasks:
  * declare an array to be sorted fill with some int
  * declare a method, passed the array as its input param
  * assert the sorted array by for loop iterator
  */
func main() {
    fmt.Println("Hello, World!")

    fmt.Println("Insert Sort begin")
    a := []int{45, 34, 78, 12, 34, 32, 29, 64}
    b := []int{12, 29, 32, 34, 34, 45, 64, 78}
    insertSort(a, len(a))
    printArray(a, len(a))
    printArray(b, len(b))
    fmt.Println("Insert Sort end")
}
