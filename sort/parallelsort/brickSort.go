package parallelsort

import(
  "github.com/jenazads/goutils";
)

// Brick Sort or Odd-Even Sort
func BrickSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low < high {
    isSorted := false;
    for !isSorted {
      isSorted = true;
      // Perform Bubble sort on odd indexed element
      for i:=low+1; i<=high-1; i=i+2 {
        if comp(arr[i], arr[i+1])==1 {
          arr[i+1], arr[i] = arr[i], arr[i+1]
          isSorted = false;
        }
      }
      // Perform Bubble sort on even indexed element
      for i:=low; i<=high-1; i=i+2 {
        if comp(arr[i], arr[i+1])==1 {
          arr[i+1], arr[i] = arr[i], arr[i+1]
          isSorted = false;
        }
      }
    }
  }
}
