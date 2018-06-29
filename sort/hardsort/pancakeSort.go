package hardsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgos/sort/sortfunctions";
)

// Pancake Sort
func PancakeSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  for curr_size := high; curr_size > 1; curr_size-- {
    index_max, _ := sortfunctions.FindMaxMinElementIndex(arr, comp, low, curr_size);
    if index_max != curr_size-1 {
      sortfunctions.Reverse(arr, low, index_max);
      sortfunctions.Reverse(arr, low, curr_size-1);
    }
  }
}
