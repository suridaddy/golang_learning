package main

import "fmt"

// 实现自定义Error

type MyStruct struct {
  error bool
  errorDesc string
  // others
}

func (myStruct MyStruct) Error() string {
  strFormat := `
  Cannot proceed, error met.
  error: %s
  `
  return fmt.Sprintf(strFormat, myStruct.errorDesc)
}

func handleStruct(myStruct MyStruct) {
  if (myStruct.error) {
    fmt.Println(myStruct.Error())
  }
  // 实现
}

func main() {
  var eStruct = MyStruct{error: true, errorDesc: "critical issue met: OutOfMemory"}
  handleStruct(eStruct)
}
