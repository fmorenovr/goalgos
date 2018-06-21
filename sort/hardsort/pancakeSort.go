package hardsort

import(
  "github.com/jenazads/goutils";
)

// Pancake Sort
func PancakeSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  // Start from the complete array and one by one 
  // reduce current size by one
  for curr_size := high; curr_size > 1; curr_size-- {
    // Find index of the maximum element in arr[low..curr_size-1]
    index_max, _ := findMaxMinElementIndex(arr, comp, low, curr_size);
    // Move the maximum element to end of 
    // current array if it's not already at the end
    if index_max != curr_size-1 {
      // To move at the end, first move maximum number to beginning 
      flip(arr, low, index_max);
      // Now move the maximum number to end by reversing current array
      flip(arr, low, curr_size-1);
    }
  }
}

// Reverses arr[low..high]
func flip(arr []interface{}, low, high int){
  for i:=low; i < high ; i++ {
    arr[high], arr[i] = arr[i], arr[high]
    high--;
  }
}

// Returns index of the maximum element in arr[0..size-1] 
func findMaxMinElementIndex(arr []interface{}, comp goutils.TypeComparator, low, high int) (int, int){
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
