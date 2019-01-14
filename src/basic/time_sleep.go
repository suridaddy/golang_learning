package main

import (
  "fmt"
  "time"
)

/*
 golang的休眠可以使用time包中的sleep。
 函数原型为：
    func Sleep(d Duration)

 其中的Duration定义为：
    type Duration int64

 Duration的单位为 nanosecond。
 为了便于使用，time中定义了时间常量：
    const (
      Nanosecond Duration = 1
      Microsecond = 1000 * Nanosecond
      Millisecond = 1000 * Microsecond
      Second = 1000 * Millisecond
      Minute = 60 * Second
      Hour = 60 * Minute
    )
*/
func main() {
  // 下面实现休眠2秒功能。
  fmt.Println("休眠2秒, time.Sleep(2 * time.Second)")
  time.Sleep(2 * time.Second)
  fmt.Println("休眠结束")

  fmt.Println("休眠2秒, time.Sleep(2 * 1000 * time.Millisecond)")
  time.Sleep(2 * 1000 * time.Millisecond)
  fmt.Println("休眠结束")

  fmt.Println("休眠2秒, time.Sleep(2 * 1000 * 1000 * time.Microsecond)")
  time.Sleep(2 * 1000 * 1000 * time.Microsecond)
  fmt.Println("休眠结束")

  fmt.Println("休眠2秒, time.Sleep(2 * 1000 * 1000 * 1000 * time.Nanosecond)")
  time.Sleep(2 * 1000 * 1000 * 1000 * time.Nanosecond)
  fmt.Println("休眠结束")
}
