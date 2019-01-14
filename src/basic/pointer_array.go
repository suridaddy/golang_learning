package main

import "fmt"

func main() {
  a := []int{10, 100, 200}

  /*
   指针数组
  */
  var ptr [3]*int

  for i := 0; i < 3; i ++ {
    ptr[i] = &a[i]
  }

  for i := 0; i < 3; i ++ {
    fmt.Printf("a[%d] = %d\n", i, *ptr[i])
  }
}
