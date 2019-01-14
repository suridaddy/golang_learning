package main

import "fmt"

/*
你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。

函数声明告诉了编译器函数的名称，返回类型，和参数。

Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。

Go 语言函数定义格式如下：
func function_name( [parameter list] ) [return_types] {
   函数体
}

函数定义解析：
    func：函数由 func 开始声明
    function_name：函数名称，函数名和参数列表一起构成了函数签名。
    parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
    return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
    函数体：函数定义的代码集合。
*/
func main() {
  fmt.Println("max(100, 20): ", max(100, 20))
  fmt.Println("max(10, 100): ", max(10, 100))
  fmt.Println("max(-100, 20): ", max(-100, 20))
  fmt.Println("max(100, -20): ", max(100, -20))
  fmt.Println("max(-100, -20): ", max(-100, -20))
  fmt.Println("max(-20, -100): ", max(-20, -100))
  fmt.Println("max(-20, -20): ", max(-20, -100))

  fmt.Println(swap("aaaa", "bbbb"))
}

func max(num1, num2 int) int {
  if (num1 > num2) {
    return num1
  } else {
    return num2
  }
}

func swap(x, y string) (string, string) {
  return y, x
}
