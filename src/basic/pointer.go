package main

import "fmt"

func main() {
  /*
   我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。
   Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
   以下实例演示了变量在内存中地址
  */
  var a int = 10

  fmt.Printf("a 变量的值: %d\n", a)
  fmt.Printf("a 变量的地址: %x\n", &a)

  /*
    一个指针变量指向一个值的内存地址。
    类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：
    var var_name *var-type
  */
  var ip *int /* 指向整型 */
  ip = &a

  fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
  fmt.Printf("*ip 变量的值: %d\n", *ip)

  /*
   Go 空指针
   当一个指针被定义后没有分配到任何变量时，它的值为 nil。
   nil 指针也称为空指针。
   nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
   一个指针变量通常缩写为 ptr。
  */
  var ptr *int
  fmt.Printf("ptr 的值为: %x\n", ptr)
  if (ptr != nil) {
    fmt.Println("var ptr * int, ptr不是空指针")
  } else {
    fmt.Println("var ptr *int, ptr是空指针")
  }

  //指针作为函数参数
  var x, y int = 100, 200
  fmt.Printf("交换前 x 的值 : %d\n", x )
  fmt.Printf("交换前 y 的值 : %d\n", y )

  swap(&x, &y)

  fmt.Printf("交换后 x 的值 : %d\n", x )
  fmt.Printf("交换后 y 的值 : %d\n", y )
}

func swap (x, y *int) {
  *x, *y = *y, *x
}
