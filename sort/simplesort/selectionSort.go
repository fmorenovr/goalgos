package simplesort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgos/sort/sortfunctions";
)

// Selection Sort
func SelectionSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low == high {
    return
  }
  k := sortfunctions.GetMinIndex(arr, comp, low, high-1);
  if k != low {
    arr[low], arr[k] = arr[k], arr[low]
  }
  SelectionSort(arr, comp, low + 1, high);
}
