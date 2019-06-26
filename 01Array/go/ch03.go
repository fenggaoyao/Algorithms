package main

import "fmt"

func mergeOrderedArray(a, b []int) (c []int) {
   i, j, k := 0, 0, 0
   mergedOrderedArrayLength := len(a) + len(b)
   c = make([]int, mergedOrderedArrayLength)
for {
   if i >= len(a) || j >= len(b) {
   break
}

if a[i] <= b[j] {
  c[k] = a[i]
  i++
} 
 else {
  c[k] = b[j]
  j++
}
   k++
}

for ; i < len(a); i++ {
  c[k] = a[i]
  k++
}

for ; j < len(b); j++ {
    c[k] = a[j]
    k++
}

return
}

func main() {
   a := []int{1, 3, 5, 7, 9, 10, 11, 13, 15}
   b := []int{2, 4, 6, 8}
   fmt.Println("ordered array a: ", a)
   fmt.Println("ordered array b: ", b)
   fmt.Println("merged ordered array: ", mergeOrderedArray(a, b))
}