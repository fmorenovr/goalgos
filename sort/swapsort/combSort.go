package swapsort

import(
  "github.com/jenazads/goutils";
)

// Comb Sort
func CombSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  swapped := true;
  size:=high-low
  for gap:=size; gap != 1 || swapped == true; {
    gap = (gap*10)/13;
    if gap < 1 {
      gap = 1
    }
    swapped = false;
    for i:=low; i<high-gap; i++ {
      if comp(arr[i], arr[i+gap]) == 1{
        arr[i+gap], arr[i] = arr[i], arr[i+gap]
        swapped = true;
      }
    }
  }
}
