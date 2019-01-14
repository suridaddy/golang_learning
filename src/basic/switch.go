package main

import "fmt"

/*
switch sExpr {
case expr1:
    some instructions
case expr2:
    some other instructions
case expr3:
    some other instructions
default:
    other code
}

sExpr和expr1、expr2、expr3的类型必须一致。Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；而如果switch没有表达式，它会匹配true。 Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码，fallthrough不会判断下一条case的expr结果是否为true。
*/

//cannot fallthrough in type switch!!!!!!
//fallthrough 必须是 case 语句块中的最后一条语句!!!!!

func main() {
  /* 定义局部变量 */
  var grade string = "B"
  var marks int = 90

  switch marks {
    case 90: grade = "A"
    case 80: grade = "B"
    case 50,60,70 : grade = "C"
    default: grade = "D"
  }

  // expression为空，表示bool true
  switch {
    case grade == "A" :
      fmt.Printf("优秀!\n" )
    case grade == "B", grade == "C" :
      fmt.Printf("良好\n" )
    case grade == "D" :
      fmt.Printf("及格\n" )
    case grade == "F":
      fmt.Printf("不及格\n" )
    default:
      fmt.Printf("差\n" )
  }
  fmt.Printf("你的等级是 %s\n", grade )

  // 常量，只要类型相符
  switch 1 {
    case 1: fmt.Println("1")
    case 3: fmt.Println("3")
    default: fmt.Println("defualt")
  }

  // 当第一个case匹配成功，直接break，后面不再匹配
  var aa, bb int = 1, 1
  switch aa {
    case 1: fmt.Println("aa == 1; no fallthrough")
    case 3: fmt.Println("3333")
    case bb: fmt.Println("aa == bb == 1")
    default: fmt.Println("defualt")
  }

  // 通过fallthrough，后面的case body强制执行
  switch aa {
    case 1:
         fmt.Println("aa == 1; fallthrough")
         fallthrough //fallthrough 必须是 case 语句块中的最后一条语句
    case 3: fmt.Println("3, 之前case有fallthrough")
    case bb: fmt.Println("aa == bb == 1")
    default: fmt.Println("defualt")
  }

  // Type-Switch 
  var x interface{}
  switch i := x.(type) {
    case nil:
      fmt.Printf("x 的类型 :%T",i)
    case int:
      fmt.Printf("x 是 int 型")
    case float64:
      fmt.Printf("x 是 float64 型")
    case func(int) float64:
      fmt.Printf("x 是 func(int) 型")
    case bool, string:
      fmt.Printf("x 是 bool 或 string 型")
    default:
      fmt.Printf("未知型")
  }
  println()
}
