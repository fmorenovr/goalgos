package efficientsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/golgos/sort/simplesort"
)

// Tim Sort
// Divide the array in blocks of BLOCKSIZE that should be a power of 2.
func TimSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  var BLOCKSIZE int
  var mid, right, min int
  power, expo := IsPowerOfTwo(high)
  if power && high > 32{
    BLOCKSIZE = high/SetPowerOfTwo(expo-2)
  } else if high > 32{
    BLOCKSIZE=32*(high%SetPowerOfTwo(4)+1)
  } else if high < 32 {
    BLOCKSIZE=32
  }
  // Sort individual subarrays of BLOCKSIZE
  for i := low; i < high; i+=BLOCKSIZE {
    if (i+(BLOCKSIZE-1)) > (high-1) {
      min = high -1
    } else {
      min = i+(BLOCKSIZE-1)
    }
    simplesort.InsertionSort(arr, comp, i, min+1);
  }
  // start merging array in size of BLOCKSIZE
  for j := BLOCKSIZE; j < high; j = 2*j {
    // We are going to merge arr[left : left+j-1] + arr[left+j : left+2*j-1].
    for left := low; left < high; left += 2*j {
      // find ending point of left sub array mid+1 is starting point of right sub array
      mid = left + j - 1;
      if (left + 2*j - 1) > (high-1) {
        right = high-1
      } else {
        right = left + 2*j - 1
      }
      // merge sub array arr[left.....mid] & arr[mid+1....right]
      Merge(arr, comp, left, mid, right);
    }
  }
}

// return True if is power of 2
func IsPowerOfTwo(x int) (bool, uint){
  var expo uint = 0
  for ((x % 2) == 0) && (x > 1) {
    x /= 2
    expo++
  }
  return ((x != 0) && ((x & (^x + 1)) == x)), expo
}

// Power Two
func SetPowerOfTwo(x uint) int {
  return 2<<x
}
