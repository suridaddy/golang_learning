package main

import "fmt"

/*
Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下： 
func (variable_name variable_data_type) function_name() [return_type]{
   \/* 函数体*\/
}
*/
type Circle struct {
  radius float64
}

func (c Circle) area() float64 {
  //c.radius 即为 Circle 类型对象中的属性
  return 3.14 * c.radius * c.radius
}

func main() {
  var c1 Circle
  c1.radius = 10.00
  fmt.Printf("半径为 %f 的圆的面积 = %f \n", c1.radius, c1.area())
}
