package swapsort

import(
  "github.com/jenazads/goutils";
)

// Bubble Sort
func BubbleSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  // Base case
  if high == 1{
      return;
  }
  // One pass of bubble sort. After this pass, the largest element is moved (or bubbled) to end.
  for i:=low; i<high-1; i++ {
    if comp(arr[i], arr[i+1])==1 {
      arr[i+1], arr[i] = arr[i], arr[i+1];
    }
  }
  BubbleSort(arr, comp, low, high-1);
}
