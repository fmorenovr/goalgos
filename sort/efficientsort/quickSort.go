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
  x := arr[low];
  i := low+1;
  j := high;
  for i < j {
    for (i < j) && (comp(arr[i], x)==-1 || comp(arr[i], x)==0) {
      i++
    }
    for (i < j) && (comp(arr[j], x)==1) {
      j--
    }
    if i == j {
      break;
    }
    arr[j], arr[i] = arr[i], arr[j]
  }
  if comp(arr[i], x)==1 {
    i--;
  }
  arr[low], arr[i] = arr[i], x;
  return i;
  /*
  //Hoare Partition
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
  */
  /*
  //Lomuto Partition
  x := arr[high];
  i := low-1;
  for j := low; j <= high-1; j++ {
    if comp(arr[j], x) == -1 ||  comp(arr[j], x) == 0 {
      i++;
      arr[j], arr[i] = arr[i], arr[j];
    }
  }
  arr[high], arr[i+1] = arr[i+1], arr[high];
  return (i + 1);
  */
}
