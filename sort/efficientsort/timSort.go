package efficientsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgos/sort/simplesort";
  "github.com/jenazads/goalgos/sort/sortfunctions";
)

// Tim Sort
// Divide the array in blocks of MINSIZERUN that should be a power of 2.
func TimSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  MINSIZERUN := high-low
  if MINSIZERUN > 64 {
    MINSIZERUN = GetMinSizeRun(high-low)
    if MINSIZERUN < 32 {
      MINSIZERUN = 32
    }
  }
  // Sort individual subarrays of MINSIZERUN
  for i := low; i < high; i+=MINSIZERUN {
    min := goutils.ToInt(sortfunctions.EvaluateMin(i+MINSIZERUN, high, goutils.IntComparator))
    simplesort.InsertionSort(arr, comp, i, min);
  }
  // start merging array in size of MINSIZERUN
  for j := MINSIZERUN; j < high; j = 2*j {
    // We are going to merge arr[left : left+j-1] + arr[left+j : left+2*j-1].
    for left := low; left < high; left += 2*j {
      // find ending point of left sub array mid+1 is starting point of right sub array
      mid := left + j - 1;
      right := goutils.ToInt(sortfunctions.EvaluateMin((left + 2*j - 1), high-1, goutils.IntComparator))
      // merge sub array arr[left.....mid] & arr[mid+1....right]
      Merge(arr, comp, left, mid, right);
    }
  }
}

// Returns the min size run
func GetMinSizeRun(num int) (int) {
  r := 0;  /* becomes 1 if the least significant bits contain at least one off bit */
  for num >= 64 {
    r |= num & 1;
    num >>= 1;
  }
  return num + r;
}
