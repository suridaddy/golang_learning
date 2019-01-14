package main

import (
  "fmt"
  "time"
)

/*
 Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
 goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

 goroutine 语法格式：
   go 函数名( 参数列表 )

 例如：
   go f(x, y, z) // 开启一个新的 goroutine:  f(x, y, z)

 Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。 同一个程序中的所有 goroutine 共享同一个地址空间。
*/

func say(s string) {
  for i := 0; i < 5; i ++ {
    var sleepTime = 100 * time.Millisecond
    fmt.Println(i, sleepTime, s)
    time.Sleep(sleepTime)
  }
}

func main() {
  go say("world1")
  go say("world2")
  say("hello")
}
