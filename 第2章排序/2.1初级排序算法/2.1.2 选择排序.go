package main

import "fmt"

func less(v, w int) bool {
  return v < w
}

func exch(a []int, i, j int) {
  t := a[i]
  a[i] = a[j]
  a[j] = t
}

func isSorted(a []int) bool {
  N := len(a)
  for i := 1; i < N; i++ {
    if less(a[i], a[i - 1]) {
      return false
    }
  }
  return true
}


func sort(a []int) {
  N := len(a)
  for i := 0; i < N; i++ {
    min := i
    for j := i + 1; j < N; j++ {
     if less(a[j], a[min]) {
       min = j
     }
    }
    exch(a, i, min)
  }
}

func main() {
  a := []int{ 2, 4, 7,1 ,3, 6, 5}
  sort(a)
  if !isSorted(a) {
    fmt.Println("sort failed", a)
  }
  fmt.Println("sort success", a)
}
