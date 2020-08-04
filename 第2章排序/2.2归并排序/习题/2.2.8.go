package main

import "fmt"

var aux []int

func merge(a []int, lo, mid, hi int) {
  i, j := lo, mid+1

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
    } else if aux[i] < aux[j] {
      a[k] = aux[i]
      i++
    } else {
      a[k] = aux[j]
      j++
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
  if a[mid] > a[mid+1] {
    merge(a, lo, mid, hi)
  }
}

func main() {
  a := []int{ 2, 4, 7, 1 ,3, 6, 5}
  sort(a)
  fmt.Println(a)
}
