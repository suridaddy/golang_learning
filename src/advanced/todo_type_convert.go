package main

import (
  "fmt"
  "strconv"
)

/*
 https://blog.csdn.net/qwdafedv/article/details/80453511
*/
func main() {
  var strInt = "1111"
  // #string到int
  i, _ := strconv.Atoi(strInt)
  fmt.Printf("\"1111\" is converted to int: %d \n", i)

  // #string到int64
  i64, _ := strconv.ParseInt(strInt, 10, 64)
  // 第二个参数为基数（2~36），
  // 第三个参数位大小表示期望转换的结果类型，其值可以为0, 8, 16, 32和64，
  // 分别对应 int, int8, int16, int32和int64
  fmt.Printf("\"1111\" is converted to int64: %d \n", i)

  // #int到string
  str := strconv.Itoa(i)
  // 等价于
  // str := strconv.FormatInt(int64(i),10)
  fmt.Printf("int(1111) is converted to string: %q \n", str)

  // #int64到string
  str = strconv.FormatInt(i64,10)
  //第二个参数为基数，可选2~36
  //对于无符号整形，可以使用FormatUint(i uint64, base int)
  fmt.Printf("int64(1111) is converted to string: %q \n", str)

  // #float到string
  //string := strconv.FormatFloat(float32,'E',-1,32)
  //string := strconv.FormatFloat(float64,'E',-1,64)
  // 'b' (-ddddp±ddd，二进制指数)
  // 'e' (-d.dddde±dd，十进制指数)
  // 'E' (-d.ddddE±dd，十进制指数)
  // 'f' (-ddd.dddd，没有指数)
  // 'g' ('e':大指数，'f':其它情况)
  // 'G' ('E':大指数，'f':其它情况)

  // #string到float64
  //float,err := strconv.ParseFloat(string,64)

  // #string到float32
  //float,err := strconv.ParseFloat(string,32)

  // #int到int64
  //int64_ := int64(1234)
}
