package main

import "fmt"

/**
  * 插入排序，移动插入版
  */
func insertSort(arr []int, n int) []int {
    for i := 1; i < n; i++ {
        tmp := arr[i]
        j := i - 1
        for ; j >= 0; j-- {
            if arr[j] > tmp {
                arr[j + 1] = arr[j]
            } else {
                break
            }
        }
        arr[j + 1] = tmp
    }
    return arr
}

func printArray(arr []int, n int) {
    for i := 0; i < n; i++ {
        fmt.Println("arr: ", i, arr[i])
    }
}

func assertArrayEquals(arr []int, expect []int, n int) bool {
    for i := 0; i < n; i++ {
        if expect[i] != arr[i] {
            return false
        }
    }
    return true
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
    //printArray(a, len(a))
    insertSort(a, len(a))
    //printArray(a, len(a))
    //printArray(b, len(b))
    if assertArrayEquals(a, b, len(a)) {
        fmt.Println("Equals!!!")
    } else {
        fmt.Println("NOT Equals!!!")
    }
    fmt.Println("Insert Sort end")
}
