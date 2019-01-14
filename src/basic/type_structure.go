package main

import . "fmt"
import "fmt"


// 常量定义
const PI = 3.14

// 全局变量的声明和赋值
var name = "gopher"

// 一般类型声明
type newType int
var i newType = 5

// 结构的声明
type gopher struct{}
var gs gopher

// 接口的声明
type golang interface{}
var gl golang

// 由main函数作为程序入口点启动
func main() {
    Println(PI, name, i, gs, gl)
    fmt.Print("a", "b", 1, 2, 3, "c", "d", "\n")
    fmt.Println("a", "b", 1, 2, 3, "c", "d")
    fmt.Printf("ab %d %d %d cd\n", 1, 2, 3)
    //println(PI, name, i, gs, gl)
    print("a", "b", 1, 2, 3, "c", "d", "\n")
    println("a", "b", 1, 2, 3, "c", "d")
    print("ab %d %d %d cd\n", 1, 2, 3)
}

