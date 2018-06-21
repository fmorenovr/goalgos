package distributionsort

import(
  "github.com/jenazads/goutils";
)

// Get max and min value with index from array.
func FindMaxMinElementIndex(arr []interface{}, comp goutils.TypeComparator, low, high int) (int, int){
  index_min, index_max := low, low
  for i := low+1; i < high; i++ {
    if comp(arr[i], arr[index_max]) == 1 {
      index_max = i
    }
    if comp(arr[i], arr[index_min]) == -1{
      index_min = i
    }
  }
  return index_max, index_min;
}
