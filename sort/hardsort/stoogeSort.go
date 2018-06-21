package hardsort

import(
  "github.com/jenazads/goutils";
)

// Stooge Sort
func StoogeSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low >= high {
    return;
  } else if low < high{
    // If first element is smaller than last, swap them
    if comp(arr[low], arr[high]) == 1 {
      arr[high], arr[low] = arr[low], arr[high]
    }
    // If there are more than 2 elements in the array
    mid := high - low + 1
    if mid > 2 {
      midthree := (mid) / 3;
      // Recursively sort first 2/3 elements
      StoogeSort(arr, comp, low, high - midthree);
      // Recursively sort last 2/3 elements
      StoogeSort(arr, comp, low + midthree, high);
      // Recursively sort first 2/3 elements again to confirm
      StoogeSort(arr, comp, low, high - midthree);
    }
  }
}
