package simplesort

import(
  "github.com/jenazads/goutils";
)

// Insertion Sort
func InsertionSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if (low == high-1){
    return;
  } else {
    InsertionSort(arr, comp, low+1, high);
    temp := arr[low];
    for j := low+1; j<high; j++ {
      if comp(temp, arr[j]) == 1 {
        arr[j-1] = arr[j];
        arr[j] = temp;
      }
    }
  }
}
