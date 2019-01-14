package main

import "fmt"
/*
import (
  "fmt"
)
*/

// 函数外定义的变量称为全局变量
var a int
var b uint
var c bool
var d float32
var e rune
var f byte
var ptr uintptr

var g int = 10
var h uint = 10
var i bool = true
var j float32 = 1.0
var k rune = 10
var l byte = 'a'
//var ptr2 uintptr = &l


// 因式分解关键字的写法一般用于声明全局变量
var (
  x int = 100
  y int
)

// 函数定义中的变量称为形式参数, 如下: i int
func test(i int) {
  println(i)
}

//这种不带声明格式的只能在函数体中出现
//hi := 123

func main() {
  println(a, b, c, d, e, f, ptr)
  fmt.Println(a, b, c, d, e, f, ptr)

  a = 10
  b = 10
  c = true
  d = 2.0
  e = 10
  f = 'b'
  println(a, b, c, d, e, f, ptr)
  fmt.Println(a, b, c, d, e, f, ptr)

  println(g, h, i, j, k, l)
  fmt.Println(g, h, i, j, k, l)

  var utype = 3.0
  fmt.Printf("%T", utype)

  // 函数内定义的变量称为局部变量
  // 多变量声明
  var name, id, address string = "xbyu", "#0001", "xicheng"
  var student_name, student_id, student_address = "zs", "#0002", "haidian"
  test_name, test_id, test_address := "ls", "#0003", "changping"

  fmt.Println(name, id, address)
  fmt.Println(student_name, student_id, student_address)
  fmt.Println(test_name, test_id, test_address)

  // 因式分解定义全局变量
  fmt.Println(x, y)
}
