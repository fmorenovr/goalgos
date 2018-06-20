package efficientsort

import(
  "github.com/jenazads/goutils";
)

// Quick Sort
func QuickSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low < high {     
    partition := QuickPartition(arr, comp, low, high); 
    QuickSort(arr, comp, low, partition - 1); 
    QuickSort(arr, comp, partition + 1, high);
  }
}

// Fast pivoting selection
func QuickPartition(arr []interface{}, comp goutils.TypeComparator, low, high int) int{
  i := low
  j := high + 1;
  for true {
    for i=i+1;comp(arr[i], arr[low]) == -1; i+=1{
      if i == high {
        break
      }
    }
    for j=j-1;comp(arr[low], arr[j]) == -1; j-=1{
      if j == low {
        break
      }
    }
    if i >= j {
      break
    }
    arr[j], arr[i] = arr[i], arr[j]
  }
  arr[j], arr[low] = arr[low], arr[j]
  return j;
  /*int x = arr[high];
  int i = (low - 1);

  for (int j = low; j <= high-1; j++){
    if (arr[j] <= x){
      i++;
      swap(arr[i], arr[j]);
    }
  }
  swap(arr[i + 1], arr[high]);
  return (i + 1);*/
}
