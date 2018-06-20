package simplesort

import(
  "github.com/jenazads/goutils";
)

// Slow Sort or Parody Sort
func SlowSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low>=high {
    return;
  }
  var m = (low+high)/2;
  SlowSort(arr,comp,low,m);
  SlowSort(arr,comp,m+1,high);
  if comp(arr[high], arr[m]) == -1 {
    arr[m], arr[high] = arr[high], arr[m]
  }
  SlowSort(arr,comp,low,high-1);
}
