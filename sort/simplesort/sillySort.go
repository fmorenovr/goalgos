package simplesort

import(
  "github.com/jenazads/goutils";
)

// Silly Sort
func SillySort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low < high {
    for i,j:=low,high; i!=j; j-- {
      mid:=(i+j)/2
      SillySort(arr, comp, i, mid)
      SillySort(arr, comp, mid+1,j)
      if comp(arr[mid], arr[j]) == 1 {
        arr[mid], arr[j] = arr[j], arr[mid]
      }
    }
  }
}
