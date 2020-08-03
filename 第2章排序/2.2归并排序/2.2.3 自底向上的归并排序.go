package main

import (
  "fmt"
)

var aux []int

func isSorted(a []int) bool {
  N := len(a)
  for i := 1; i < N; i++ {
    if less(a[i], a[i - 1]) {
      return false
    }
  }
  return true
}

func less(v, w int) bool {
  return v < w
}

func merge(a []int, lo, mid, hi int) {
  i, j := lo, mid + 1


  // copy array
  for k := lo; k <= hi; k++ {
    aux[k] = a[k]
  }

  for k := lo; k <= hi; k++ {
    if i > mid {
      a[k] = aux[j]
      j++
    } else if j > hi {
      a[k] = aux[i]
      i++
    } else if less(aux[j], aux[i]) {
      a[k] = aux[j]
      j++
    } else {
      a[k] = aux[i]
      i++
    }
  }
}

func min(v, w int) int {
  if v > w {
    return w
  } else {
    return v
  }
}

func sort(a []int) {
  aux = make([]int, len(a))
  N := len(a)
  for sz := 1; sz < N; sz = sz + sz {
    for lo := 0; lo < N - sz; lo += sz {
      merge(a, lo, lo + sz - 1, min(lo + sz + sz -1, N - 1))
    }
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
