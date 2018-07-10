package swapsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgos/sort/sortfunctions";
)

// Permutation Sort
func PermutationSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if(low<high){
    for !sortfunctions.IsSorted(arr, comp, low, high) {
      sortfunctions.Permute(arr, comp, low, high)
    }
  }
}
