package hardsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgos/sort/swapsort";
)

// Bad Sort or Shotgun Sort
func BadSort(arr []interface{}, comp goutils.TypeComparator, low, high, k int) () {
  if low < high {
    if k==0 {
      swapsort.BubbleSort(arr, comp, low, high)
    } else {
      /*P:=Permutations(arr, low, high)
      BadSort(P, comp, low, high, k-1)*/
    }
  }
}
