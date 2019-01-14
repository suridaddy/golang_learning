package main

import "fmt"

func main() {
  a := 1.69
  b := 1.7
  c := a * b      // 结果应该是2.873
  fmt.Println(c)  // 输出的是2.8729999999999998

  // 浮点数计算输出有一定的偏差，你也可以转整型来设置精度。
  x := 1690           // 表示1.69
  y := 1700           // 表示1.70
  z := x * y          // 结果应该是2873000表示 2.873
  fmt.Println(z)      // 内部编码
  fmt.Println(float64(z) / 1000000) // 显示
}
