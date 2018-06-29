package swapsort

import(
  "fmt"
  "github.com/jenazads/goutils";
)

// Permutation Sort
func PermutationSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  for !permutationSort(arr, comp, low, high-1) {
    fmt.Println(arr)
    // empty block
  }
}

// permutation sort
func permutationSort(arr []interface{}, comp goutils.TypeComparator, low, high int) bool {
  if high <= 0 {
    for i := high - 1; (comp(arr[i], arr[i-1])==1 ||comp(arr[i], arr[i-1])==0) ; i-- {
      if i == low+1 {
        return true
      }
    }
    return false
  }
  for i := low; i <= high; i++ {
    arr[i], arr[high] = arr[high], arr[i]
      if permutationSort(arr, comp, low, high - 1) {
        return true
      }
    arr[i], arr[high] = arr[high], arr[i]
  }
  return false
}
