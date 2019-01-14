package main

import "fmt"

/*
 Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
 // 定义接口
 type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
 }

 // 定义结构体
 type struct_name struct {
   // variables
 }

 // 实现接口方法
 func (struct_name_variable struct_name) method_name1() [return_type] {
   // 方法实现
 }
 ...
 func (struct_name_variable struct_name) method_namen() [return_type] {
   // 方法实现
 }
*/

type phone_interface interface {
  //call() 不能出现同名不同参函数
  call(pnum string)
}

type MatePhone struct {
  phoneNum string
}

type IPhone struct {
  phoneNum string
}

type NokiaPhone struct {
  phoneNum string
}

func (matePhone MatePhone) call(pnum string) {
  fmt.Println("This call is from Mate -", matePhone.phoneNum, ", and call to", pnum)
}

func (iphone IPhone) call(pnum string) {
  fmt.Println("This call is from IPhone -", iphone.phoneNum, ", and call to", pnum)
}

func (nokiaPhone NokiaPhone) call(pnum string) {
  fmt.Println("This call is from Nokia -", nokiaPhone.phoneNum, ", and call to", pnum)
}

func main() {
  var matePhone = MatePhone{"18915201207"}
  matePhone.call("13401128371")
  var iphone = MatePhone{"18915201222"}
  iphone.call("13401128371")
  var nokiaPhone = MatePhone{"18915200827"}
  nokiaPhone.call("13401128371")
}
