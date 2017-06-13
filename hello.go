// package main
//
// import "fmt"
//
// func sum(x int, y int) int {
//   return x + y
// }
//
//
// func main() {
//   fmt.Println(sum(1, 1))
//
// }

// package main
//
// import "fmt"
//
// func main() {
//   var pointer *int
//   var n = 100
//   pointer = &n
//   fmt.Println(pointer, &n)
// }

// package main
//
// import "fmt"
//
// func main() {
//   a, b := 10, 10
//   called(a, &b)
//   fmt.Println(a, b)
// }
//
// func called(a int, b *int) {
//   a  = a + 1
//   *b = *b + 1
// }

// package main
//
// import "fmt"
//
// func main() {
//   a, b := 10, 10
//   sample := called(a, b)
//   fmt.Printf("sampleは%d\n", sample)
// }
//
// func called(a int, b int) int {
//   a  = a + 1
//   b = b + 1
//   return a + b
// }

// package main
// import "fmt"
//
// func main() {
//   if 12 == 12 {
//     fmt.Println("aaa")
//   }
//   sum := 0
//   for i := 0; i < 10; i++ {
//     sum += i
//     fmt.Println(sum)
//   }
//
//   items := map[int]int{0:10 , 1:20}
//   for key, value := range items {
//     fmt.Println("key:", key, " value:", value)
//   }
// }


package main
import "fmt"

func main() {
  nums := []int{1, 2, 3, 4, 5}
  for i, v :=  range nums {
    fmt.Println(i, v)
  }
  var slice []int = make([]int, 3, 3)
  if cap(slice) <= 5{
    fmt.Println("OK")
  }
}
