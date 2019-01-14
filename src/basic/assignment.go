package main

import "fmt"

type Human struct {
  Name string
  Address string
}

func main() {
  var i int = 10
  var j = i

  fmt.Println(&i)
  fmt.Println(&j)

  i = 100
  fmt.Println(i, j)


  var human1 = Human {Name: "xbyu", Address: "xicheng"}
  var human2 = human1

  human1.Address = "xicheng-modified"
  fmt.Println(human1)
  fmt.Println(human2)
  fmt.Printf("%v, %v\n", human1, human2)

  //Golang 只有slice(切片)、map(字典)、channel(管道)三种引用类型
  a := [5]int{2, 3, 4, 5, 6}
  b := a
  fmt.Println(a,b)
  b[2] = 77
  fmt.Println(a,b)

  c := []int{2, 3, 4, 5, 6}
  d := c
  fmt.Println(c,d)
  b[2] = 77
  fmt.Println(c,d)

  // 并行/同时 赋值
  // 并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)
  e, f, g, h := 2, 1, 3, 4

  // 交换变量的值
  e, f = f, e
  fmt.Println(e, f)

  e, f, g, h = h, g, f, e
  fmt.Println(e, f, g, h)

  // _ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值
  _, k := 1, 2
  fmt.Println(k)
}
