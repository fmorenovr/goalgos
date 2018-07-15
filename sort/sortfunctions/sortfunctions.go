package sortfunctions

import (
  "math/rand";
  "github.com/jenazads/goutils";
)

// return min between a, b
func EvaluateMin(a, b interface{}, comp goutils.TypeComparator) interface{} {
  if comp(a ,b) == 1 {
    return b
  } else {
    return a
  }
}

// Returns the max_min element in arr[0..size-1] 
func FindMaxMinElementIndex(arr []interface{}, comp goutils.TypeComparator, low, high int) (interface{}, interface{}){
  index_min, index_max := low, low
  for i := low+1; i < high; i++ {
    if comp(arr[i], arr[index_max]) == 1 {
      index_max = i
    }
    if comp(arr[i], arr[index_min]) == -1{
      index_min = i
    }
  }
  return arr[index_min], arr[index_max];
}

// Return min index
func GetMinIndex(arr []interface{}, comp goutils.TypeComparator, i, j int) int {
  if i == j {
    return i
  }
  k := GetMinIndex(arr, comp, i + 1, j)
  if comp(arr[i], arr[k]) == -1 {
    return i
  } else {
    return k
  }
}

// Return max index
func GetMaxIndex(arr []interface{}, comp goutils.TypeComparator, i, j int) int {
  if i == j {
    return i
  }
  k := GetMaxIndex(arr, comp, i + 1, j)
  if comp(arr[i], arr[k]) == 1 {
    return i
  } else {
    return k
  }
}

// Generate random permutation in the array
func Shuffle(arr []interface{}, comp goutils.TypeComparator, low, high int){
  for i:=low; i<high; i++ {
    rand_index := rand.Intn(high-low)+low
    if comp(arr[i], arr[rand_index]) == 1 {
      arr[rand_index], arr[i] = arr[i], arr[rand_index]
    }
  }
}

// Generate permutation in the array
func Permute(arr []interface{}, comp goutils.TypeComparator, low, high int) {
  for j:=low;j<high;j++ {
    if comp(arr[j], arr[high-1]) == 1 {
      arr[high-1], arr[j] = arr[j], arr[high-1]
    }
    Permute(arr, comp, low, high-1)
    if comp(arr[j], arr[high-1]) == 1 {
      arr[high-1], arr[j] = arr[j], arr[high-1]
    }
  }
}

// verify is Sorted
func IsSorted(arr []interface{}, comp goutils.TypeComparator, low, high int) (bool){
  for i:=low; i<high-1; i++ {
    if comp(arr[i], arr[i+1]) == 1 {
      return false;
    }
  }
  return true;
}

// Reverses arr[low..high]
func Reverse(arr []interface{}, low, high int){
  for i:=low; i < high ; i, high = i+1, high-1 {
    arr[high], arr[i] = arr[i], arr[high]
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
  return 2<<(x-1)
}

// return the highest number less than N
func HighestPowerofTwoLessThan(n int) (int) {
  res := 0;
  for i:=n; i>=1; i-- {
    if (i & (i-1)) == 0 {
      res = i;
      break;
    }
  }
  return res;
}
