package efficientsort

import(
  "math";
  "github.com/jenazads/goutils";
)

// Intro Sort
func IntroSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  maxDepth:=int(math.Floor(math.Log2(float64(high-low))))*2
  introSort(arr, comp, maxDepth, low, high)
}

// introSort
func introSort(arr []interface{}, comp goutils.TypeComparator, maxDepth, low, high int) () {
  size:=high-low
  if size <= 1 {
    return
  } else if maxDepth == 0{
    HeapSort(arr, comp, low, high)
  } else {
    mid := QuickPartition(arr, comp, low, high)
    introSort(arr, comp, maxDepth-1, low, mid-1)
    introSort(arr, comp, maxDepth-1, mid+1, high)
  }
}
