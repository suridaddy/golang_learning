package main

import "fmt"

/*
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

切片中有两个概念：一是len长度，二是cap容量，长度是指已经被赋过值的最大下标+1，可通过内置函数len()获得。容量是指切片目前可容纳的最多元素个数，可通过内置函数cap()获得。切片是引用类型，因此在当传递切片时将引用同一指针，修改值将会影响其他的对象。 
*/

func main() {
  /*
   你可以声明一个未指定大小的数组来定义切片：
   var identifier []type
   切片不需要说明长度。

   或使用make()函数来创建切片:
   var slice1 []type = make([]type, len)

   也可以简写为
   slice1 := make([]type, len)

   也可以指定容量，其中capacity为可选参数。
   make([]T, length, capacity)
   这里 len 是数组的长度并且也是切片的初始长度。
  */
  var slice1 []int
  fmt.Printf("var slice1 []int, slice1 = %v, (slice1 == nil): %v\n", slice1, slice1 == nil)

  var slice2 = make([]int, 0)
  fmt.Printf("var slice2 = make([]int, 0), slice2 = %v, (slice2 == nil): %v\n", slice2, slice2 == nil)
  println()

  // 尝试给nil slice赋值
  // slice1[0] = 1 // error: panic: runtime error: index out of range
  newSlice := append(slice1, 1)
  fmt.Println("nil slice只能通过 append 添加item, append(slice1, 1), slice1 = ", slice1)
  fmt.Println("nil slice只能通过 append 添加item, newSlice := append(slice1, 1), newSlice = ", newSlice)
  println()

  // 切片初始化
  s := []int {1, 2, 3, 4}
  fmt.Println("s := []int {1, 2, 3, 4}")
  fmt.Println("s = ", s)
  fmt.Println("len(s) = ", len(s))
  fmt.Println("cap(s) = ", cap(s))
  println()

  // 用数组s来初始化ns
  ns := s[0:2]
  fmt.Println("ns = s[0:2]")
  fmt.Println("ns = ", ns)
  fmt.Println("len(ns) = ", len(ns))
  fmt.Println("cap(ns) = ", cap(ns))
  println()

  // 用默认值初始化
  s = make([]int, 3)
  fmt.Println("s := make([]int, 3)")
  fmt.Println("s = ", s)
  fmt.Println("len(s) = ", len(s))
  fmt.Println("cap(s) = ", cap(s))
  println()

  // 用默认值初始化
  s = make([]int, 3, 6)
  fmt.Println("s := make([]int, 3, 6)")
  fmt.Println("s = ", s)
  fmt.Println("len(s) = ", len(s))
  fmt.Println("cap(s) = ", cap(s))
  println()

  var nilSlice []int
  fmt.Println("var nilSlice []int")
  fmt.Println("nilSlice = ", nilSlice)
  fmt.Println("len(nilSlice) = ", len(nilSlice))
  fmt.Println("cap(nilSlice) = ", cap(nilSlice))
  fmt.Println("(nilSlice == nil):", (nilSlice == nil))
  println()


  fmt.Println("切片截取:")
  // 切片截取
  /* 创建切片 */
  numbers := []int{0,1,2,3,4,5,6,7,8}
  printSlice(numbers)

  /* 打印原始切片 */
  fmt.Println("numbers ==", numbers)

  /* 打印子切片从索引1(包含) 到索引4(不包含)*/
  fmt.Println("numbers[1:4] ==", numbers[1:4])

  /* 默认下限为 0*/
  fmt.Println("numbers[:3] ==", numbers[:3])

  /* 默认上限为 len(s)*/
  fmt.Println("numbers[4:] ==", numbers[4:])

  numbers1 := make([]int,0,5)
  printSlice(numbers1)

  /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
  number2 := numbers[:2]
  printSlice(number2)

  /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
  number3 := numbers[2:5]
  number3[1] = 2
  printSlice(number3)
  printSlice(numbers)
  fmt.Println("调整子切片后，父切片也随之发生变化，子切片是对父切片的一个引用！！可以通过copy来实现复制!!")

  var parentArray [9]int = [9]int {0, 1, 2, 3, 4, 5, 6, 7, 8}
  number4 := parentArray[2:5]
  number4[1] = 2
  printSlice(number4)
  printSlice(numbers)
  fmt.Println("调整子切片后，父数组也随之发生变化，子切片是对父数组的一个引用！！可以通过copy来实现复制!!")
}

func printSlice(x []int){
  fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
