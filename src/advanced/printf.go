package main

import "fmt"

/*
 https://768992698.iteye.com/blog/2326069
*/
func main() {
  var var1 string = "aaaa"
  // 以默认的方式打印变量的值
  fmt.Printf("var var1 string = \"aaa\", var1: %v \n", var1)

  // 打印变量的类型
  fmt.Printf("var var1 string = \"aaa\", type of var1: %T\n ", var1)

{
  // %+d 带符号的整型，fmt.Printf("%+d", 255)输出+255
  var i int = 10
  fmt.Printf("var i int = 10, i: %+d \n", i)

  // %5d, 表示该整型最大长度是5
  fmt.Printf("var i int = 10, i: %5d \n", i)

  // %-5d, 表示该整型最大长度是5, 左对齐
  fmt.Printf("var i int = 10, i: %-5d \n", i)

  // %05d, 在数字的前面补零
  fmt.Printf("var i int = 10, i: %05d \n", i)

  // %-05d, 在数字的前面补零
  fmt.Printf("var i int = 10, i: %-05d \n", i)

  i = 255
  // %o 不带零的八进制
  fmt.Printf("var i int = 255, i: %o \n", i)

  // %#o 带零的八进制
  fmt.Printf("var i int = 255, i: %#o \n", i)

  // %+#o 带符号的带零的八进制
  fmt.Printf("var i int = 255, i: %+#o \n", i)

  // %x 小写的16进制
  // %X 大写的16进制
  // %#x 带0x的16进制
  fmt.Printf("var i int = 255, i: %#X \n", i)

  // %b 打印整型的2进制
  fmt.Printf("var i int = 255, i: %b \n", i)
}

{
  var1 = "\"my `\""
  // %s 正常输出字符串
  fmt.Printf("var var1 string = \"my\", var1: %s \n", var1)

  // %q 字符串带双引号，字符串中的”带转义
  fmt.Printf("var var1 string = \"my\", var1: %q \n", var1)

  // %#q 字符串带反引号，如果字符串内有反引号，就用双引号代替
  fmt.Printf("var var1 string = \"my\", var1: %#q \n", var1)
  var1 = "\"my\""
  fmt.Printf("var var1 string = \"my\", var1: %#q \n", var1)
  // %x 将字符串转换为小写的16进制格式
  fmt.Printf("var var1 string = \"my\", var1: %x \n", var1)
  // %X, 大写的16进制；
  fmt.Printf("var var1 string = \"my\", var1: %X \n", var1)
  // % x, 带空格的16进制格式
  fmt.Printf("var var1 string = \"my\", var1: % x \n", var1)

  // %-5.7s, 最小宽度为5，最大宽度为7，左对齐
  var1 = "aa"
  fmt.Printf("var var1 string = \"my\", var1: %-5.7s \n", var1)
  var1 = "aaaaa"
  fmt.Printf("var var1 string = \"my\", var1: %-5.7s \n", var1)
  var1 = "aaaaaaa"
  fmt.Printf("var var1 string = \"my\", var1: %-5.7s \n", var1)
  var1 = "aaaaaaaaa"
  fmt.Printf("var var1 string = \"my\", var1: %-5.7s \n", var1)
}

{
  // %f， %.6f 6位小数点
  var f float64 = 3.2
  fmt.Printf("var f float = 3.2, f: %f\n", f)
  fmt.Printf("var f float = 3.2, f: %.2f \n", f)

  // %e, %.6e 6位小数点，科学计数法
  fmt.Printf("var f float = 3.2, f: %e\n", f)

  // %g 用最少的数字来表示
  fmt.Printf("var f float = 3.2, f: %g\n", f)

  // %.3g 最多三位数字来表示
  f = 3.3333
  fmt.Printf("var f float = 3.2, f: %.3g\n", f)

  // %.3f 最多三位小数来表示
  f = 3.3333
  fmt.Printf("var f float = 3.2, f: %.3f\n", f)
}

{
  type AS struct {
    name string
    addr string
  }

  as := AS {
    name: "yxb",
    addr: "xicheng",
  }

  // %v 打印struct
  fmt.Printf("struct AS: %v \n", as)
  //%+v 带字段名称。比如：{name:sam phone:{mobile:12345 office:67890}
  fmt.Printf("struct AS: %+v \n", as)
  // %#v  用Go的语法打印
  fmt.Printf("struct AS: %#v \n", as)
}

{
  var b bool = true
  // %t 打印true或false
  fmt.Printf("var b bool = true, b: %t \n", b)
}

{
  var p *int
  var i int = 10
  p = &i
  fmt.Printf("pointer p: %v \n", p)
  // 不带0x的指针
  fmt.Printf("pointer p: %#p \n", p)
  // 带0x的指针
  fmt.Printf("pointer p: %p \n", p)
}
}
