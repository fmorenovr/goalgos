package swapsort

import(
  "github.com/jenazads/goutils";
)

// Circle Sort
func CircleSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  for circleSort(arr, comp, low, high-1) {
    // empty block
  }
}

// circle sort
func circleSort(arr []interface{}, comp goutils.TypeComparator, low, high int) bool {
  swapped := false
  if low == high {
    return false
  }
  hi, lo := high, low
  for lo < hi {
    if comp(arr[lo], arr[hi]) == 1 {
      arr[lo], arr[hi] = arr[hi], arr[lo]
      swapped = true
    }
    lo++
    hi--  
  }
  if lo == hi {
    if comp(arr[lo], arr[hi+1]) == 1 {
      arr[lo], arr[hi+1] = arr[hi+1], arr[lo]
      swapped = true
    }
  }
  mid   := (high - low) / 2
  right := circleSort(arr, comp, low, low+mid)
  left  := circleSort(arr, comp, low+mid+1, high)
  return swapped || right || left
}
