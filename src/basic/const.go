package main

import "fmt"

// const identifier [type] = value
// 显式类型定义
const a string = "abc"

// 隐式类型定义
const b = 1.0

/* 
wrong usage:
const c string
c = "abc"
*/

// 常量大写表示更直观
const WIDTH int = 10
const HEIGHT int = 69

// 常量用作枚举
const (
  unknown = 0
  Female = 1
  Male = 2
)

// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
const (
  d = iota
  e = iota
  f = iota
)

const (
  g = iota
  h
  i
)

const (
  j = "haha"
  k
  l = 'h'
  m
)

const (
  n = iota   // 0
  o          // 1
  p = "haha" // "haha", iota=2
  q          // "haha", iota=3
  r = iota   // 4
  s          // 5
  t = 100    // 100, iota=6
  u          // 100, iota=7
)

const (
  v = 1 << iota // 1 << 0 (= 1)
  w = 3 << iota // 3 << 1 (= 6)
  x             // 3 << 2 (= 12)
  y             // 3 << 3 (= 24)
)

var ii = 1
const (
  // xx = (ii + iota) //除了iota，不能出现别的变量
  xx = (1 + iota)
  yy
  zz
)

func main() {
  fmt.Println(a, b)
  var area = WIDTH * HEIGHT
  fmt.Printf("面积是:%d", area)
  println()

  fmt.Println(Female, Male, unknown)
  fmt.Println(d, e, f)
  fmt.Println(g, h, i)
  fmt.Println(j, k, l, m)
  fmt.Println(n, o, p, q, r, s, t, u) // 0, 1, "haha", "haha", 4, 5, 100, 100
  fmt.Println(v, w, x, y)
  fmt.Println(xx, yy, zz)
}
