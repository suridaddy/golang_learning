package main

import "fmt"

func main() {
  // 值传递
  var a, b int = 100, 200

  fmt.Printf("交换前，a 的值 : %d\n", a )
  fmt.Printf("交换前，b 的值 : %d\n", b )

  /* 调用 swap() 函数 */
  swap(a, b)

  fmt.Printf("交换后，a 的值 : %d\n", a )
  fmt.Printf("交换后，b 的值 : %d\n", b )
}

/* 定义相互交换值的函数 */
func swap(x, y int) {
  x, y = y, x    /* 将 y 值赋给 x, x 值赋给 y */
}
