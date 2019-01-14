package main

import . "fmt"
import "strings"

func main() {
  str := "这里是 www\n.runoob\n.com"
  Println(str)

  str = strings.Replace(str, " ", "", -1)
  str = strings.Replace(str, "\n", "", -1)
  Println(str)
}
