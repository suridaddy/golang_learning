package main

import (
  "fmt"
  "time"
)
/*
 Channel类型的定义格式如下：

 ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .

 它包括三种类型的定义。可选的<-代表channel的方向。如果没有指定方向，那么Channel就是双向的，既可以接收数据，也可以发送数据。

 chan T          // 可以接收和发送类型为 T 的数据
 chan<- float64  // 只可以用来发送 float64 类型的数据
 <-chan int      // 只可以用来接收 int 类型的数据
*/

/*
 通道（channel）是用来传递数据的一个数据结构。
 通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
 ch <- v    // 把 v 发送到通道 ch
 v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v

 声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：
 ch := make(chan int)

 注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须又接收端相应的接收数据。
*/

// 以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：
func sum(s []int, c chan int) {
  sum := 0
  for _, v := range s {
    sum += v
  }

  c <- sum
}

func main() {
  s := []int {1, 2, 4, 5, 2, 3}

  c := make(chan int)
  go sum(s[:len(s)/2], c)
  go sum(s[len(s)/2: ], c)

  x, y := <-c, <-c //从通道c中读取
  fmt.Println("s := []int {1, 2, 4, 5, 2, 3}, 各项之和: ", (x + y))

  // 关闭通道
  close(c)

  /* 通道缓冲区
   通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
   ch := make(chan int, 100)

   带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。

   不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。

   注意：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。
  */
  c = make(chan int, 10)
  go sum(s[:len(s)/2], c)
  go sum(s[len(s)/2: ], c)

  time.Sleep(2 * time.Second)

  sum := 0
  /*
   range 函数遍历每个从通道接收到的数据; range之前通道需要关闭，否则会出现： fatal error: all goroutines are asleep - deadlock!
   如果 c 通道不关闭，那么 range 函数就不会结束，从而发生阻塞。
  */
  close(c)
  for i := range c {
    sum += i
  }
  fmt.Println("通过设置缓冲区计算: s := []int {1, 2, 4, 5, 2, 3}, 各项之和: ", sum)

  ch := make(chan string, 4)
  go waitToSend(2, "2222222", ch)
  go waitToSend(0, "0000000", ch)
  go waitToSend(1, "1111111", ch)
  go waitToSend(3, "3333333", ch)

  for msg := range ch {
    fmt.Println(msg)
  }

  // 可以通过 len(c), cap(c)来获得的当前channel中未被消费的消息数及最大数
  cc := make(chan int, 4)
  cc <- 1
  cc <- 2
  cc <- 3
  cc <- 4
  // cc <- 5 // 如果发送消息大于缓冲区cap，或者没设置缓冲区的channel收到超过1个消息, 则会出错: fatal error: all goroutines are asleep - deadlock!
  fmt.Println("length of the channel, expected: 4, actual:", len(cc))
  fmt.Println("cap of the channel, expected: 4, actual:", cap(cc))
  close(cc)
  //for v := range c {
  //  fmt.Println(v)
  //}
}

func waitToSend(sleep time.Duration, msg string, c chan string) {
  time.Sleep(sleep * time.Second)
  c <- msg
  if sleep == 3 {
    fmt.Println("close the channel during 'range'")
    close(c)
  }
}
