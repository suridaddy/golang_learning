package main

import "fmt"

func main() {
  /*
   在for多层嵌套时，有时候需要直接跳出所有嵌套循环， 这时候就可以用到go的label breaks特征了。
   label要写在for循环的开始而不是结束的地方。和goto是不一样的。虽然它是直接break退出到指定的位置。
  */
breakLabel1:
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      if i+j > 15 {
        fmt.Println("exit")
        // 跳到外面去啦，但是不会再进来这个for循环了
        break breakLabel1
      }
    }
  }

  // break 标签除了可以跳出 for 循环，还可以跳出 select switch 循环， 参考下面代码：
  count := 1
  var ch chan int
breakLabel2:
  for ; count < 8192; count++ {
    select {
      case e := <-ch:
        fmt.Println("e := \"aaaa\"", e)
      default:
        break breakLabel2 // 跳出 select 和 for 循环
    }
  }

  // break 标签除了可以跳出 for 循环，还可以跳出 select switch 循环， 参考下面代码：
  var grade string = "B"
  count = 1
breakLabel3:
  for ; count < 8; count++ {
    switch count {
      case 5: grade = "A"; println(grade); break breakLabel3
      case 4: grade = "B"; println(grade)
      case 3, 2: grade = "C"; println(grade)
      default: grade = "D"; println(grade)
    }
  }

  /*
   continue label跳出当前该次的循环圈，在Go编程语言中的continue语句有点像break语句。不是强制终止，只是继续循环下一个迭代发生，在两者之间跳过任何代码。
  */
  count = 1
continueLabel:
  for count < 10 {
    if count == 5 {
      count ++
      continue continueLabel
    }
    println(count)
    count ++
  }

  /*
   跳转语句 goto语句可以跳转到本函数内的某个标签
  */
  gotoCount := 0
GotoLabel:
  gotoCount++
  if gotoCount < 10 {
    goto GotoLabel //如果小于10的话就跳转到GotoLabel
  }
}
