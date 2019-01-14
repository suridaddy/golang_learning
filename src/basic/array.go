package main

import "fmt"

/*
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。

Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：
  var variable_name [SIZE] variable_type

Go 语言支持多维数组，以下为常用的多维数组声明方式：
  var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

数组初始化:
[5] int {1,2,3,4,5}
长度为5的数组，其元素值依次为：1，2，3，4，5

[5] int {1,2}
长度为5的数组，其元素值依次为：1，2，0，0，0 。在初始化时没有指定初值的元素将会赋值为其元素类型int的默认值0,string的默认值是"" 

[...] int {1,2,3,4,5}
长度为5的数组，其长度是根据初始化时指定的元素个数决定的 

[5] int { 2:1,3:2,4:3}
长度为5的数组，key:value,其元素值依次为：0，0，1，2，3。在初始化时指定了2，3，4索引中对应的值：1，2，3 

[...] int {2:1,4:3}
长度为5的数组，起元素值依次为：0，0，1，0，3。由于指定了最大索引4对应的值3，根据初始化的元素个数确定其长度为5赋值与使用
*/

/*
Go 语言向函数传递数组
方式一
形参设定数组大小：

void myFunction(param [10]int)
{
.
.
.
}

方式二
形参未设定数组大小：

void myFunction(param []int)
{
.
.
.
}
*/
//func getAverage(arr [...] int, size int) float32 { // ./array.go:95:21: use of [...] array outside of array literal
func getAverage(arr [] int, size int) float32 {
  var sum int
  for i := 0; i < size; i ++ {
    sum += arr[i]
  }

  return float32(sum) / float32(size)
}

func main() {
  var balance [10] float32
  for i := 0; i < 10; i ++ {
    fmt.Print(float32(i))
    // 逐个赋值, 不能超过声明个数, 否则 panic: runtime error: index out of range
    balance[i] = 1.1 * float32(i)
    fmt.Println("\t", balance[i])
  }
  fmt.Println(balance)

  // 整体赋值, 初始化数组中 {} 中的元素个数不能大于 [] 中的数字。否则， ./array.go:24:78: array index 10 out of bounds [0:10]
  balance = [10] float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9}
  for idx :=range balance {
    fmt.Println(balance[idx])
  }

  // 初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
  var newBalance = [10] float32{1.0, 2.1, 3.2, 4.3, 5.4, 6.5, 7.6, 8.7, 9.8}
  fmt.Println(newBalance)

  // 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小
  var noSizeSetBalance = [...] float32{1.1, 2.2}
  // noSizeSetBalance[2] = 3.3 //出错: ./array.go:35:16: invalid array index 2 (out of bounds for 2-element array)
  fmt.Println(noSizeSetBalance)

  // 访问数组
  var balance2 float32 = balance[1]
  fmt.Println(balance2)

  // 声明了三维的整型数组，并直接赋值
  var threedim = [5][10][3] int{
  //var threedim [5][10][3]int = [5][10][3] int{
                                 {{1, 2, 3}, {3, 4, 5}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}, {16, 17, 18}, {19, 20, 21}, {22, 23, 24}, {25, 26, 27}, {28, 29, 30}},
                                 {{1, 2, 3}, {3, 4, 5}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}, {16, 17, 18}, {19, 20, 21}, {22, 23, 24}, {25, 26, 27}, {28, 29, 30}},
                                 {{1, 2, 3}, {3, 4, 5}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}, {16, 17, 18}, {19, 20, 21}, {22, 23, 24}, {25, 26, 27}, {28, 29, 30}},
                                 {{1, 2, 3}, {3, 4, 5}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}, {16, 17, 18}, {19, 20, 21}, {22, 23, 24}, {25, 26, 27}, {28, 29, 30}},
                                 {{1, 2, 3}, {3, 4, 5}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}, {16, 17, 18}, {19, 20, 21}, {22, 23, 24}, {25, 26, 27}, {28, 29, 30}}, // 最后的, 一定要存在，否则: ./array.go:50:165: syntax error: unexpected newline, expecting comma or }
                               }
  fmt.Println(threedim)
  fmt.Println("len(threedim) = ", len(threedim))
  fmt.Println("threedim[1][3][2]应该是12(第二行，第4列，第三个元素):", threedim[1][3][2])
}
