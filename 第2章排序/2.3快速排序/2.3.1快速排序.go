package main

import "fmt"

func exch(a []int, i, j int) {
  t := a[i]
  a[i] = a[j]
  a[j] = t
}

func partition(a []int, lo, hi int) int {
  i, j := lo + 1, hi
  v := a[lo]

  for {
    for a[i] < v {
      i++
      if i == hi {
        break
      }
    }

    for v < a[j] {
      j--
      if j == lo {
        break
      }
    }

    if i >= j {
      break
    }

    exch(a, i, j)
  }

  // 交换两个元素，确保切分位置的正确
  exch(a, lo, j)
  return j
}


func sort(a []int, lo, hi int) {
  if hi <= lo {
    return
  }

  j := partition(a, lo, hi)
  sort(a, lo, j - 1)
  sort(a, j + 1, hi)
}

func main() {
  a := []int{ 2, 4, 7, 1 ,3, 6, 5}
  sort(a, 0, len(a) - 1)
  fmt.Println(a)
}
