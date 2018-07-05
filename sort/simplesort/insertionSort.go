package simplesort

import(
  "github.com/jenazads/goutils";
)

// Insertion Sort
func InsertionSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low < high {
    InsertionSort(arr, comp, low, high-1);
    last := arr[high-1];
    j:=high-2
    for ;j>=low  && comp(arr[j], last) ==1;j-- {
      arr[j+1]=arr[j]
    }
    arr[j+1]=last
  }
}
