package hardsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgos/sort/sortfunctions";
)

// Pancake Sort
func PancakeSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low < high {
    for curr_index := high-1; curr_index > low; curr_index-- {
      max_index := sortfunctions.GetMaxIndex(arr, comp, low, curr_index);
      sortfunctions.Reverse(arr, low, max_index);
      sortfunctions.Reverse(arr, low, curr_index);
    }
  }
}
