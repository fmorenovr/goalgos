package distributionsort

import(
  "github.com/jenazads/goutils";
)

// Counting Sort
func CountingSort(arr []interface{}, comp goutils.TypeComparator, op goutils.TypeOperator, low, high int) () {
  maxValueIndex, minValueIndex := FindMaxMinElementIndex(arr, comp, low, high)
  minValue, maxValue := arr[minValueIndex], arr[maxValueIndex]
  numOfCounts := goutils.ToInt(op(op(maxValue, minValue, "-"), 1, "+"));
  counts := make([]interface{}, numOfCounts)
  
  for i:=low; i<high; i++ {
    index_curr := goutils.ToInt(op(arr[i], minValue, "-"))
    counts[index_curr] = goutils.ToInt(op(counts[index_curr], 1, "+"));
  }
  /*j:=low;
  for i:=minValue; i<=maxValue; i++ {
    for z:=0; z<counts[i-minValue]; z++ {
      arr[j] = i;
      j+=1
    }
  }*/
}
