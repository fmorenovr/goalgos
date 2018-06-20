package simplesort

import(
  "github.com/jenazads/golgos/sort";
)

func GnomeSort(arr []interface{}, comparator TypeComparator, low, high int){
  size := len(arr)
  i := 0
  for i < size {
    if i==0 || (arr[i-1] <= arr[i]) {
      i = i+1;
    } else {
      arr[i], arr[i-1] = swap(arr[i-1], arr[i]);
      i = i-1;
      if i == -1 {
        i = 0;
      }
    }
  }
}
