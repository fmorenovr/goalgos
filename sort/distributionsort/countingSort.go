package distributionsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgos/sort/sortfunctions";
)

// Counting Sort or Histogram Sort
func CountingSort(arr []interface{}, comp goutils.TypeComparator, op goutils.TypeOperator, low, high int) () {
  if(low<high){
    minValue, maxValue := sortfunctions.FindMaxMinElementIndex(arr, comp, low, high)
    numOfCounts := goutils.ToInt(op(op(maxValue, minValue, "-"), 1, "+"));
    counts := make([]interface{}, numOfCounts)
    
    for i:=0; i<numOfCounts;i++ {
      counts[i] = 0;
    }
    
    for i:=low; i<high; i++ {
      curr_index := goutils.ToInt(op(arr[i], minValue, "-"))
      counts[curr_index] = op(counts[curr_index], 1, "+");
    }
    index:=low;
    for i:=minValue; comp(maxValue,i) ==0 || comp(maxValue,i) ==1; i=op(i,1,"+") {
      curr_index := goutils.ToInt(op(i, minValue, "-"))
      for z:=0; z<goutils.ToInt(counts[curr_index]); z++ {
        arr[index] = i;
        index++;
      }
    }
  }
}
