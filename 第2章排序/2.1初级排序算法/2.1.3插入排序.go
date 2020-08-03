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
  for i := 1; i < N; i++ {
    //for j := i; j > 0 && less(a[j], a[j - 1]); j-- {
    //  exch(a, j, j - 1)
    //}

    num := a[i]
    j := i
    for ; j > 0 && less(num, a[j - 1]); j-- {
      a[j] =  a[j - 1]
    }
    a[j] = num
  }
}

func main() {
  a := []int{ 2, 4, 7, 1 ,3, 6, 5}
  sort(a)
  if !isSorted(a) {
    fmt.Println("sort failed", a)
  } else {
    fmt.Println("sort success", a)
  }
}
