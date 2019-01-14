package main

import (
  "fmt"
)

func main() {
  // nil slice
  var nilSlice []int
  fmt.Println("var nilSlice []int, (nilSlice == nil): ", (nilSlice == nil))
  /* panic: runtime error: index out of range */ //nilSlice[0] = 1
  // nil slice 赋值
  newSlice := append(nilSlice, 0)
  fmt.Println("nil slice, 只能通过 append(like, newSlice := append(nilSlice, 0)) 赋值, newSlice = ", newSlice)
  // 或者使用make函数创建一个非nil的slice, 仔赋值

  // nil map
  var nilMap map[string]string
  fmt.Println("var nilMap map[string]string, (nilMap == nil): ", (nilMap == nil))
  // nil map 赋值
  /* panic: assignment to entry in nil map */ // nilMap [ "key1" ] = "value1"
  fmt.Println("nil map, 不可以直接赋值，like: nilMap [ \"key1\" ] = \"value1\"")
  // 再使用make函数创建一个非nil的map，nil map不能赋值
  nilMap = make(map[string]string)
  // 最后给已声明的map赋值
  nilMap["a"] = "aa"

  // 没有nil array
  var arr [3]int
  arr[0] = 1
  /* cannot convert nil to type [3]int */ // fmt.Println("var arr [3]int, (arr == nil): ", (arr == nil))
}
