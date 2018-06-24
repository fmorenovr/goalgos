package efficientsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgs/sort/sortfunctions";
)

// Tim Sort
// Divide the array in blocks of BLOCKSIZE that should be a power of 2.
func TimSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  var BLOCKSIZE int
  var mid, right, min int
  power, expo := sortfunctions.IsPowerOfTwo(high)
  if power && high > 32{
    BLOCKSIZE = high/sortfunctions.SetPowerOfTwo(expo-2)
  } else if high > 32{
    BLOCKSIZE=32*(high%sortfunctions.SetPowerOfTwo(4)+1)
  } else if high < 32 {
    BLOCKSIZE=32
  }
  // Sort individual subarrays of BLOCKSIZE
  for i := low; i < high; i+=BLOCKSIZE {
    min = goutils.ToInt(sortfunctions.EvaluateMin((i+(BLOCKSIZE-1)), high -1, goutils.IntComparator))
    QuickSort(arr, comp, i, min+1);
  }
  // start merging array in size of BLOCKSIZE
  for j := BLOCKSIZE; j < high; j = 2*j {
    // We are going to merge arr[left : left+j-1] + arr[left+j : left+2*j-1].
    for left := low; left < high; left += 2*j {
      // find ending point of left sub array mid+1 is starting point of right sub array
      mid = left + j - 1;
      right = goutils.ToInt(sortfunctions.EvaluateMin((left + 2*j - 1), high -1, goutils.IntComparator))
      // merge sub array arr[left.....mid] & arr[mid+1....right]
      Merge(arr, comp, left, mid, right);
    }
  }
}
