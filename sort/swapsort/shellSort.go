package swapsort

import(
  "github.com/jenazads/goutils";
)

// Shell Sort
func ShellSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  size := high - low;
  gap:=size/2;
  for gap > 0 {
    for i := gap; i < size; i += 1 {
      j := i;
      temp := arr[i];
      for j >= gap && (comp(arr[j-gap], temp) == 1) {
        arr[j] = arr[j - gap];
        j -= gap;
      }
      arr[j] = temp;
    }
    if gap == 2 {
      gap = 1
    } else {
      gap = gap*10/22
    }
  }
}
