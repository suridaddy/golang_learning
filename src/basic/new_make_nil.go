package main

import "fmt"

func main()  {
    //new返回指针
    p1 := new(int)    //分配一个地址空间，new函数会返回地址p1,并会初始化值为0（类型的空值），地址指向（引用）的值可变
    fmt.Printf(" p1 --> %#v \n ", p1)  //nil只能赋值给指针、channel、func、interface、map或slice类型的变量。如果将nil赋值给其他变量的时候将会引发panic
    fmt.Printf("p1 point to --> %#v \n ", *p1)
    *p1 = 10
    fmt.Printf("p1 --> %#v \n ", p1)
    fmt.Printf("p1 point to --> %#v \n ", *p1)
    var p2 *int
    fmt.Printf("p2 --> %#v \n ", p2)    //即没有分配地址空间，不能引用
    //fmt.Printf("p2 point to --> %#v \n ", *p2) //空指针的引用会导致编译错误

    var p3 *int         //即没有分配地址空间，不能引用
    j := 5
    jj := 10
    p3 = &j             //此时，p3指向地址，可以引用
    fmt.Printf("p2 --> %#v \n ", p3)  //0xc0420620b8
    fmt.Printf("p2 point to --> %#v \n ", *p3) //5
    p3 = &jj
    fmt.Printf("p3 --> %#v \n ", p3)   //0xc0420620c0
    fmt.Printf("p3 point to --> %#v \n ", *p3) //golang中变量初始化分配内存地址，j的地址没有指针指向，0xc0420620b8地址指向的值暂时无法改变

    var t int
    fmt.Printf("t address --> %#v \n ", &t)   //0xc042008110 由此可知，变量定义是会初始化地址，而且会赋予类型零值,
    fmt.Printf("t value is --> %#v \n ", t)
    t = 10
    fmt.Printf("t --> %#v \n ", &t)   //t 是0xc042008110指向值的引用，要使函数活方法中改变外部变量的值，必须传递指针，否则改变都是局部变量地址指向的值
    fmt.Printf("t value now --> %#v \n ", t)

    changeInt(t)
    //x := 11
    //&t = &x    //指针不能运算

    //make 用法，只能用于slice，map，channel，可用nil赋值
    var s1 []int

    if s1 == nil {    //未初始化
        fmt.Printf("s1 address  --> %#v \n ", &s1)
        fmt.Printf("s1 is nil --> %#v \n ", s1) // []int(nil)
    }
    s2 := make([]int,3)    // 返回的是引用类型，即可在函数中直接改变s2的值
    if s2 == nil {   //地址不为空，等号不成立
        fmt.Printf("s2 is nil --> %#v \n ", s2)
    } else {   //初始化后，分配地址空间
        fmt.Printf("s2 address  --> %#v \n ", &s2)
        fmt.Printf("s2 is not nil --> %#v \n ", &s2[0])// 0xc0420600c0
        fmt.Printf("s2 is not nil --> %#v \n ", &s2[1])// 0xc0420600c8
        fmt.Printf("s2 is not nil --> %#v \n ", s2)// []int{0, 0, 0}
    }
    changeSlice(s2)

    x := 11
    var addr = &x
    fmt.Printf("addr is not nil --> %#v \n", addr)
}

func changeSlice(s []int) {
    fmt.Printf("s[0] address  --> %#v \n ", &s[0]) //0xc0420600c0   地址相同，说明是地址传递，map channel 逻辑相同，golang源码中实现的逻辑
    s[0] = 5
}
func changeInt(s int) {
    fmt.Printf("t in func address  --> %#v \n ", &s)
    s = 5
}
