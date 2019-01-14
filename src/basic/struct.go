package main

import "fmt"

/*
结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体有中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

type struct_variable_type struct {
   member definition;
   member definition;
   ...
   member definition;
}

一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：

variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
*/
type Human struct {
  Name string
  id string
  address string
}

func modify(human Human) {
  human.id = human.id + "-modified"
}

func modifyId( human *Human) {
  (*human).id = (*human).id + "-modified"
}

func main() {
  // 初始化方法1
  var human = Human{
    Name : "zhangsan",
    id : "#0002",
    address : "haidian"}
  // 初始化方法2
  var human1 = Human { "xbyu", "#001", "xicheng"}

  fmt.Printf("%v\n", human1)
  fmt.Printf("%v\n", human)

  // 初始化方法3
  var person = Human{}
  person.Name = "lisi"
  person.id = "#0003"
  person.address = "changping"
  fmt.Printf("%v\n", person)

  // 初始化方法4
  var person1 Human
  person1.Name = "wangwu"
  person1.id = "#0004"
  person1.address = "shunyi"
  fmt.Printf("%v\n", person1)

  // 作为函数参数，by value 传值，copy
  fmt.Println("modify #0001 human's id")
  modify(human1)
  fmt.Printf("After modification(copy by value): %v\n", human1)

  // 结构体指针做函数参数，by reference
  fmt.Println("modify #0001 human's id - pointer")
  modifyId(&human1)
  fmt.Printf("After modification(copy by value): %v\n", human1)
}
