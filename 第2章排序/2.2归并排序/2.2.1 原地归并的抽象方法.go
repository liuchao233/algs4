package main

var aux []int

func less(v, w int) bool {
  return v < w
}

// 将 a[lo, mid] 和 a[mid+1, hi]进行归并
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

func main() {

}
