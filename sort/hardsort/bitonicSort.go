package hardsort

import(
  "github.com/jenazads/goutils";
)

// Bitonic Sort
// high-low must be power of 2
func BitonicSort(arr []interface{}, comp goutils.TypeComparator, low, high int, order bool) () {
  // order == true increase
  // order == false decrease
  if high>1 {
    mid := high/2;
    // sort in ascending order since order here is true
    BitonicSort(arr, comp, low, mid, true);
    // sort in descending order since order here is false
    BitonicSort(arr, comp, low+mid, mid, false);
    // Will merge wole sequence in ascending order since order.
    BitonicMerge(arr, comp, low, high, order);
  }
}

// compare and swap
func BitonicCompare(arr []interface{}, comp goutils.TypeComparator, low, distance int, order bool){
  for i := low; i < low+distance; i++ {
    if order==(comp(arr[i], arr[i+distance])==1) {
      arr[i+distance], arr[i] = arr[i], arr[i+distance]
    }
  }
}

// The sequence to be sorted starts at index position low, high is the number of elements to be sorted
func BitonicMerge(arr []interface{}, comp goutils.TypeComparator, low, high int, order bool){
  if high > 1 {
    mid := high/2;
    BitonicCompare(arr, comp, low, mid, order);
    BitonicMerge(arr, comp, low, mid, order);
    BitonicMerge(arr, comp, low+mid, mid, order);
  }
}
