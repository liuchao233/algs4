package main

import "fmt"

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

func sort(a []int) {
  aux = make([]int, len(a))
  _sort(a, 0, len(a) - 1)
}

func _sort(a []int, lo, hi int) {
  if hi <= lo {
    return
  }

  mid := lo + (hi - lo) / 2
  _sort(a, lo, mid)
  _sort(a, mid + 1, hi)
  merge(a, lo, mid, hi)
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
