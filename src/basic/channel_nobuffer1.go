package main

import (
  "fmt"
  "time"
)

// 以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：
func send(i int, c chan int) {
  fmt.Printf("发送(%d)...\n", i)
  c <- i
}

func main() {
  c := make(chan int)
  go send(1, c)
  go send(2, c)

  time.Sleep(2 * time.Second) // runtime判断是否deadlock, 此时还未执行完，不能判定是否deadlock
  println("发送必须有相应接收，否则: fatal error: all goroutines are asleep - deadlock!")
  x, y := <-c, <-c // 从通道c中读取
  fmt.Println("读取到的先后顺序:", x, y) // 先发送到channel里的消息，被最先消费/读取
}
