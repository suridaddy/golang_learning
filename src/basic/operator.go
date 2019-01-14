package main

import "fmt"

/*
运算符优先级

有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：
优先级	运算符
7 	^ !
6 	* / % << >> & &^
5 	+ - | ^
4 	== != < <= >= >
3 	<-
2 	&&
1 	||

*/

func main() {
  var a int = 21
  var b int = 10
  var c int

  fmt.Println("a = 21, b = 10:")
  // 算术运算符
  fmt.Println("===============算术运算符==============")
  c = a + b
  fmt.Printf("a + b 的值为 %d\n", c )
  c = a - b
  fmt.Printf("a - b 的值为 %d\n", c )
  c = a * b
  fmt.Printf("a * b 的值为 %d\n", c )
  c = a / b
  fmt.Printf("a / b 的值为 %d\n", c )
  c = a % b
  fmt.Println("a % b 的值为", c )
  a ++
  fmt.Printf("a++ 之后 a 的值为 %d\n", a )
  a = 21   // 为了方便测试，a 这里重新赋值为 21
  a --
  fmt.Printf("a-- 之后 a 的值为 %d\n", a )

  a = 21
  // 关系运算符
  fmt.Println("===============关系运算符==============")
  fmt.Println("a == b:", a == b)
  fmt.Println("a != b:", a != b)
  fmt.Println("a > b:", a > b)
  fmt.Println("a < b:", a < b)
  fmt.Println("a >= b:", a >= b)
  fmt.Println("a <= b:", a <= b)

  // 位运算符
  fmt.Println("===============位运算符==============")
  fmt.Println("a & b:", a & b)
  fmt.Println("a | b:", a | b)
  fmt.Println("a ^ b:", a ^ b)
  fmt.Println("a << 1:", a << 1)
  fmt.Println("a >> 1:", a >> 1)

  // 赋值运算符
  fmt.Println("===============赋值运算符==============")
  a = b
  fmt.Println("a = b, a:", a)
  a = 21
  a += b
  fmt.Println("a += b, a:", a)
  a = 21
  a -= b
  fmt.Println("a -= b, a:", a)
  a = 21
  a *= b
  fmt.Println("a *= b, a:", a)
  a = 21
  a /= b
  fmt.Println("a /= b, a:", a)
  a = 21
  a %= b
  fmt.Println("a %= b, a:", a)
  a = 21
  a <<= 10
  fmt.Println("a <<= 10, a:", a)
  a = 21
  a >>= 10
  fmt.Println("a >>= 10, a:", a)
  a = 21
  a &= b
  fmt.Println("a &= b, a:", a)
  a = 21
  a |= b
  fmt.Println("a |= b, a:", a)
  a = 21
  a ^= b
  fmt.Println("a ^= b, a:", a)

  println()
  // 逻辑运算符
  fmt.Println("d = true, e = false:")
  fmt.Println("===============逻辑运算符==============")
  var d bool = true
  var e = false
  fmt.Println("d && e:", d && e)
  fmt.Println("d || e:", d || d)
  fmt.Println("!d:", !d)
  fmt.Println("!e:", !e)

  // & * 运算符
  a = 21
  var ptr * int
  ptr = &a
  fmt.Println("===============& * 运算符==============")
  fmt.Println("var ptr * int; ptr = &a, ptr:", ptr)
  fmt.Println("var ptr * int; ptr = &a, *ptr:", *ptr)
}
