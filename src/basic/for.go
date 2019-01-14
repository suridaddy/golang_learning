package main

import "fmt"

/*
Go语言的For循环有3中形式，只有其中的一种使用分号。

和 C 语言的 for 一样：
for init; condition; post { }

和 C 的 while 一样：
for condition { }

和 C 的 for(;;) 一样：
for { }

Attention:
  - init： 一般为赋值表达式，给控制变量赋初值；
  - condition： 关系表达式或逻辑表达式，循环控制条件；
  - post： 一般为赋值表达式，给控制变量增量或减量。

循环嵌套:
for [condition |  ( init; condition; increment ) | Range]
{
   for [condition |  ( init; condition; increment ) | Range]
   {
      statement(s);
   }
   statement(s);
}
*/
func main() {
  numbers := [6] int{1, 2, 3, 4, 5, 6}

  // for 循环
  for i := 0; i < 10; i ++ {
    fmt.Printf("i 的值是: %d\n", i)
  }

  var b int = 15
  var a int = 10
  // while
  for a < b {
    a ++
    fmt.Printf("a 的值是: %d \n", a)
  }

  // range
  for idx, value := range numbers {
    fmt.Printf("第%d位的值是: %d\n", idx, value)
  }

  // 嵌套，打印100以内的素数
  var x, y int
  for x = 2; x < 100; x ++ {
    for y = 2; y <= (x / y); y ++ {
      if x % y == 0 {
        // 不是素数
        break
      }
    }

    if y > (x / y) {
      fmt.Printf("%d 是素数\n", x)
    }
  }
}
