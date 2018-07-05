package swapsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgos/sort/sortfunctions";
)

// Permutation Sort
func PermutationSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if(low<high){
    for !sortfunctions.IsSorted(arr, comp, low, high) {
      Permute(arr, comp, low, high)
    }
  }
}

// Generate permutation in the array
func Permute(arr []interface{}, comp goutils.TypeComparator, low, high int) {
  for j:=low;j<high;j++ {
    if comp(arr[j], arr[high-1]) == 1 {
      arr[high-1], arr[j] = arr[j], arr[high-1]
    }
    Permute(arr, comp, low, high-1)
    if comp(arr[j], arr[high-1]) == 1 {
      arr[high-1], arr[j] = arr[j], arr[high-1]
    }
  }
}
