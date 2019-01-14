package main

import (
  "fmt"
  "time"
)

// Attention: 多运行几次，就会发现3个receiver，都有可能第一个接收消息，接收到的消息也呈随机状

func receive(index int, c chan int) {
  fmt.Printf("接收(queue%d) ready...\n", index)
  // 阻塞，等待数据
  fmt.Printf("接收(queue%d)接收数据: %d\n", index, <-c)
}

func main() {

  c := make(chan int)
  go receive(1, c)
  go receive(2, c)
  go receive(3, c) // 接收可以多于发送，不报错! 
  // 并行读取时，receiver随机接收消息!

  time.Sleep(2 * time.Second)
  println("必须读取和接收都存在, 否则: fatal error: all goroutines are asleep - deadlock!")
  c <- 1
  c <- 2
  // 发送必须有接收! 发送数不能小于接受数，否则: fatal error: all goroutines are asleep - deadlock!
  // 运行时检测deadlock，实时检测 
  c <- 3

  // 阻塞主程序结束，否则等不到并发线程打印结果
  time.Sleep(2 * time.Second)
}

