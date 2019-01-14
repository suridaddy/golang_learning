package main

import "fmt"

/*
for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
  for key, value := range oldMap {
    newMap[key] = value
  }
*/
func main() {

  var iArray [5] int = [5] int {1, 2, 3, 4, 5}
  var slice [] int = [] int{1, 2, 3, 4, 5, 6, 7}
  var noSizeArray = [...] int {1, 2, 3, 4}

  fmt.Println("array[index]\t data")
  for idx, data := range iArray {
    fmt.Println(iArray[idx], "\t\t", data)
  }

  for idx, data := range slice {
    fmt.Println(slice[idx], "\t\t", data )
  }

  for idx, data := range noSizeArray {
    fmt.Println(noSizeArray[idx], "\t\t", data)
  }

  
}
