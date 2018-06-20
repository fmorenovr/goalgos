package simplesort

import(
  "github.com/jenazads/goutils";
)

// Selection Sort
func SelectionSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low == high {
    return
  }
  k := MinIndex(arr, comp, low, high-1);
  if k != low {
    arr[low], arr[k] = arr[k], arr[low]
  }
  SelectionSort(arr, comp, low + 1, high);
}

// Return min index
func MinIndex(arr []interface{}, comp goutils.TypeComparator, i, j int) int {
  if i == j {
    return i
  }
  k := MinIndex(arr, comp, i + 1, j)
  if comp(arr[i], arr[k]) == -1 {
    return i
  } else {
    return k
  }
}
