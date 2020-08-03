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
  h := 1
  for h < N / 3 {
    h = 3 * h + 1
  }
  for h >= 1 {
    for i := h; i < N; i++ {
      // 插入排序
      for j := i; j >= h && less(a[j], a[j - h]); j -= h {
        exch(a, j, j - h)
      }
    }
    h = h / 3
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
