package main

import "fmt"

/*
如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：

指向指针的指针变量声明格式如下：
var ptr **int;
*/
func main() {
  var a float32 = 39.29
  var ptr *float32
  var pptr **float32

  ptr = &a
  pptr = &ptr

  fmt.Printf("变量 a = %f\n", a)
  fmt.Printf("指针变量 *ptr = %f\n", *ptr)
  fmt.Printf("指向指针的指针变量 **pptr = %f\n", **pptr)
}
