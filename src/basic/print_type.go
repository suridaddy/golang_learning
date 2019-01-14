package main

import . "fmt"

func main() {
  var untype = 3.0
  Printf("%T\n", untype)

/* error: can't assign string to float64 variable
  untype = "abc"
  Printf("%T\n", untype)
*/
  untype = 'a'
  Printf("%T\n", untype)

  untype1 := 3.0
  Printf("%T\n", untype1)

  untype2 := "abc"
  Printf("%T\n", untype2)

  untype3 := 'a'
  Printf("%T\n", untype3)

}
