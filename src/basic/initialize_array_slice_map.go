package main

import "fmt"

func main() {
  var arr [3]int = [3]int {1, 2, 3}
  fmt.Println(arr)

  arr = [3]int {
            3,
            2,
            1}
  fmt.Println(arr)

  arr = [3]int {
            1,
            2,
            3,
  }
  fmt.Println(arr)


  var slice []int = []int {1, 3, 5, 7}
  fmt.Println(slice)

  slice = []int {
            1,
            2,
            3,
            4,
            }
  fmt.Println(slice)

  slice = []int {
            1,
            2,
            3}
  fmt.Println(slice)

  var mp map[int]int = map[int]int {1: 1, 2: 2, 3: 3}
  fmt.Println(mp)

  mp = map[int]int {
            3: 3,
            2: 2,
            1: 1}
  fmt.Println(mp)

  mp = map[int]int {
            1: 1,
            2: 2,
            3: 3,
  }
  fmt.Println(mp)
}
