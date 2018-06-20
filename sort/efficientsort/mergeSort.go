package efficientsort

import(
  "github.com/jenazads/goutils";
)

// Merge Sort
func MergeSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if(low<high){
    mid:=low+(high-low)/2;
    // Divide and Conquer
    MergeSort(arr,comp,low,mid);
    MergeSort(arr,comp,mid+1,high);
    // Combine
    Merge(arr,comp,low,mid,high);
  }
}

// Merges two subarrays of arr[] = arr[low..mid] + arr[mid+1..high]
func Merge(arr []interface{}, comp goutils.TypeComparator, low, mid, high int)(){
  if low == high {
    return
  }
  n1 := mid - low + 1;
  n2 := high - mid;
  L := make([]interface{}, n1)
  H := make([]interface{}, n2)
  // Copy data to temp arrays Low[] and High[]
  for i := 0; i < n1; i++{
    L[i] = arr[low + i]
  }
  for j := 0; j < n2; j++ {
    H[j] = arr[mid+1+j]
  }
  // Merge the temp arrays back into arr[low..high]
  i := 0 // Initial index of first subarray
  j := 0 // Initial index of second subarray
  k := low // Initial index of merged subarray
  for i < n1 && j < n2 {
    if comp(L[i], H[j]) == -1 || comp(L[i], H[j]) == 0 {
      arr[k] = L[i]
      i+=1
    } else{
      arr[k] = H[j]
      j+=1
    }
    k+=1
  }
  // Copy the remaining elements of L[], if there are any
  for i < n1 {
    arr[k] = L[i];
    i+=1
    k+=1
  }
  // Copy the remaining elements of R[], if there are any
  for j < n2 {
    arr[k] = H[j];
    j+=1
    k+=1
  }
}
