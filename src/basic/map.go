package main

import "fmt"

/*
 Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
 Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。 

 定义 Map
 可以使用内建函数 make 也可以使用 map 关键字来定义 Map:
   \/* 声明变量，默认 map 是 nil *\/
   var map_variable map[key_data_type]value_data_type

   \/* 使用 make 函数 *\/
   map_variable := make(map[key_data_type]value_data_type)

 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
*/
func main() {
  // 声明map
  var countryCapitalMap map[string]string
  println("var countryCapitalMap map[string]string, (countryCapitalMap == nil):",  (countryCapitalMap == nil))

  // nil map 赋值
  /* panic: assignment to entry in nil map */ // countryCapitalMap [ "France" ] = "Paris, nil"
  fmt.Println("nil map 不能直接赋值，like: countryCapitalMap [ \"France\" ] = \"Paris, nil\"")

  countryCapitalMap = make(map[string]string)
  fmt.Println("countryCapitalMap = make(map[string]string), (countryCapitalMap == nil):", (countryCapitalMap == nil))

  /* map插入key - value对,各个国家对应的首都 */
  countryCapitalMap [ "France" ] = "Paris"
  countryCapitalMap [ "Italy" ] = "罗马"
  countryCapitalMap [ "Japan" ] = "东京"
  countryCapitalMap [ "India " ] = "新德里"

  /*使用键输出地图值 */ for country := range countryCapitalMap {
    fmt.Println(country, "首都是", countryCapitalMap [country])
  }

  /*查看元素在集合中是否存在 */
  captial, ok := countryCapitalMap [ "美国" ] /*如果确定是真实的,则存在,否则不存在 */
  /*fmt.Println(captial) */
  /*fmt.Println(ok) */
  if (ok) {
    fmt.Println("美国的首都是", captial)
  } else {
    fmt.Println("美国的首都不存在")
  }

  // 先声明map
  var m1 map[string]string
  // 再使用make函数创建一个非nil的map，nil map不能赋值
  m1 = make(map[string]string)
  // 最后给已声明的map赋值
  m1["a"] = "aa"
  m1["b"] = "bb"

  // 直接创建
  m2 := make(map[string]string)
  // 然后赋值
  m2["a"] = "aa"
  m2["b"] = "bb"

  // ==========================================
  // 查找键值是否存在
  if v, ok := m1["a"]; ok {
    fmt.Println(v)
  } else {
    fmt.Println("Key Not Found")
  }

  // 遍历map
  for k, v := range m2 {
    fmt.Println(k, v)
  }

  // 初始化 + 赋值一体化
  m3 := map[string]string{
    "a": "aa",
    "b": "bb",
  }

  if v, ok := m3["a"]; ok {
    fmt.Println("m3[\"a\"] = ", v)
  }
}
