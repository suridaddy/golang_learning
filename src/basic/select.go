package main

import "fmt"
import "time"
/*
 select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。
 select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。

 语法:
 select {
    case communication clause  :
       statement(s);      
    case communication clause  :
       statement(s); 
    \/* 你可以定义任意数量的 case *\/
    default : \/* 可选 *\/
       statement(s);
 }

 Attentions:
  - 每个case都必须是一个通信
  - 所有channel表达式都会被求值
  - 所有被发送的表达式都会被求值
  - 如果任意某个通信可以进行，它就执行；其他被忽略。
  - 如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
  - 否则：
     - 如果有default子句，则执行该语句。
     - 如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
*/

func send(i int, c chan int) {
  c <- i
}

func receive(c chan int, qName string) {
  fmt.Println("go: received", <-c, "from", qName)
}

/*
 下面的select充满各种可能性，例如多次执行会有不同的结果:

 beefeater@beefeater-ThinkPad-W541:~/BPM_DEV/go/golang_learning/src/basic$ go run select.go 
 sent 10 to c2
 go: received 10 from c2

 beefeater@beefeater-ThinkPad-W541:~/BPM_DEV/go/golang_learning/src/basic$ go run select.go 
 received 1 from c1

 beefeater@beefeater-ThinkPad-W541:~/BPM_DEV/go/golang_learning/src/basic$ go run select.go 
 received 3 from c3
*/
func main() {
  var c1, c2, c3 = make(chan int), make(chan int), make(chan int)
  var i1, i2 int
  i2 = 10
  go send(1, c1)
  go send(3, c3)

  go receive(c2, "c2")

  time.Sleep(2 * time.Second)
  select {
    case i1 = <-c1: // Attention: 如果没有之前的 go send(1, c1), 则该case不会被执行!!!!
      fmt.Println("received", i1, "from c1")
    case c2 <- i2: // Attention: 如果没有之前的 go receive(c2, "c2"), 则该case不会被执行!!!!
      fmt.Println("sent", i2, "to c2")
    case i3, ok := (<-c3):  // same as: i3, ok := <-c3
      if ok {
        fmt.Println("received", i3, "from c3")
      } else {
        fmt.Println("c3 is closed")
      }
    default:
      fmt.Println("no communication")
  }

  // 阻塞主程序，以便所有的receive, send都能打印出
  time.Sleep(2 * time.Second)
}
