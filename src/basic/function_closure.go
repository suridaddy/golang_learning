package main

import "fmt"

/*
 Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

以下实例中，我们创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量，代码如下：
*/

func getSequence() func() int {
  i := 0
  return func() int {
    i += 1
    return i
  }
}

// 带多个返回值的闭包函数调用:
func add(x, y int) func() (int, int) {
  i := 0
  return func() (int, int) {
    i ++
    return i, x + y
  }
}

// 带参数闭包函数，返回多值
func add_add(x, y int) func(i, j int)(int, int, int) {
  a := 0
  return func(i, j int) (int, int, int) {
    a ++
    return a, x + y, i + j
  }
}

func main() {
  /* nextNumber 为一个函数，函数 i 为 0 */
  nextNumber := getSequence()

  /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
  fmt.Println(nextNumber())
  fmt.Println(nextNumber())
  fmt.Println(nextNumber())

  /* 创建新的函数 nextNumber1，并查看结果 */
  nextNumber1 := getSequence()
  fmt.Println(nextNumber1())
  fmt.Println(nextNumber1())

  // 调用带参数的闭包函数
  add_func := add(1, 2)
  fmt.Println(add_func())
  fmt.Println(add_func())

  add_func1 := add_add(1, 2)
  fmt.Println(add_func1(1, 2))
  fmt.Println(add_func1(2, 3))
}
