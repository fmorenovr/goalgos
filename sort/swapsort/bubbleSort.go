package swapsort

import(
  "github.com/jenazads/goutils";
)

// Bubble Sort
func BubbleSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low < high {
    for i:=low; i<high-1; i++ {
      if comp(arr[i], arr[i+1])==1 {
        arr[i+1], arr[i] = arr[i], arr[i+1];
      }
    }
    BubbleSort(arr, comp, low, high-1);
  }
}
