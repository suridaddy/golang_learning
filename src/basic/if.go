package main

import "fmt"

/*
if 语句	if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
if 布尔表达式 {
   \/* 在布尔表达式为 true 时执行 *\/
}

if...else 语句	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
if 布尔表达式 {
   \/* 在布尔表达式为 true 时执行 *\/
} else {
  \/* 在布尔表达式为 false 时执行 *\/
}

if 嵌套语句	你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
if 布尔表达式 1 {
   \/* 在布尔表达式 1 为 true 时执行 *\/
   if 布尔表达式 2 {
      \/* 在布尔表达式 2 为 true 时执行 *\/
   }
}
*/

func main() {
  if (1 < 10) {
    fmt.Println("if (1 < 10) {...}")
    fmt.Println("  match ...")
  }
  println()

  if 1 < 10 {
    fmt.Println("if 1 < 10 {...}")
    fmt.Println("  match ...")
  }
  println()

  fmt.Println("if 1 > 10 {...} else {...}")
  if 1 > 10 {
    fmt.Println("  match if block ...")
  } else {
    fmt.Println("  match else block ...")
  }
  println()

  fmt.Println("if 1 > 10 {...} else if 2 > 10 {....} else {...}")
  if 1 > 10 {
    fmt.Println("  match if block ...")
  } else if 2 > 10 {
    fmt.Println("  match else if block ...")
  } else {
    fmt.Println("  match else block ...")
  }
}
