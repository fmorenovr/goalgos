package simplesort

import(
  "github.com/jenazads/goutils";
)

// Gnome Sort or Stupid Sort
func GnomeSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low < high {
    i := low
    for i < high {
      if i==low {
        i=i+1
      }
      if comp(arr[i-1], arr[i]) == -1 || comp(arr[i-1], arr[i]) == 0 {
        i = i+1
      } else {
        arr[i], arr[i-1] = arr[i-1], arr[i]
        i=i-1
        if i == -1 {
          i = low
        }
      }
    }
  }
}
