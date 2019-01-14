package main

import "fmt"

/*
goto 语法格式如下：

goto label;
..
.
label: statement;
*/
func main() {

  var a int = 10

  LOOP: for a < 20 {
    if a == 15 {
      a ++
      goto LOOP
    }
    fmt.Printf("a的值是: %d\n", a)
    a ++
  }
}
