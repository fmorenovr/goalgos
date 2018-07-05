package hardsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgos/sort/sortfunctions";
)

// Bogo Sort or Shotgun Sort
func BogoSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low < high {
    for !sortfunctions.IsSorted(arr, comp, low, high) {
      sortfunctions.Shuffle(arr, low, high);
    }
  }
}
