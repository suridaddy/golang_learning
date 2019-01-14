package main

import "fmt"

/*
 递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。
*/

// 阶乘
func Factorial(n uint64) uint64 {
  var result uint64
  if n > 0 {
    result = n * Factorial(n - 1)
    return result
  }

  return 1
}

// 斐波拉契数列
func fibonacci(n int) int {
  if n == 1 || n == 2 {
    return 1
  }

  return fibonacci(n - 1) + fibonacci(n - 2)
}

func main() {
  var i int = 15
  fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

  fmt.Print("打印前20的斐波拉契数列: ")
  for j := 1; j <= 20; j ++ {
    fmt.Printf("%d ", fibonacci(j))
  }
}
