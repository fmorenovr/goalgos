package swapsort

import(
  "github.com/jenazads/goutils";
)

// CockTail Sort
func CockTailSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  swapped := true;
  start := low;
  end := high - 1;
  for swapped {
    swapped = false;
    for i := start; i < end; i+=1 {
      if comp(arr[i], arr[i + 1]) == 1 {
        arr[i+1], arr[i] = arr[i], arr[i + 1]
        swapped = true;
      }
    }
    if !swapped {
      break;
    }
    swapped = false;
    end-=1;
    for i := end - 1; i >= start; i-=1 {
      if comp(arr[i], arr[i + 1])==1 {
        arr[i+1], arr[i] = arr[i], arr[i + 1]
        swapped = true;
      }
    }
    start+=1;
  }
}
