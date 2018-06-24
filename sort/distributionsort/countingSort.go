package distributionsort

import(
  "fmt"
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgs/sort/sortfunctions";
)

// Counting Sort or Histogram Sort
func CountingSort(arr []interface{}, comp goutils.TypeComparator, op goutils.TypeOperator, low, high int) () {
  maxValueIndex, minValueIndex := sortfunctions.FindMaxMinElementIndex(arr, comp, low, high)
  minValue, maxValue := arr[minValueIndex], arr[maxValueIndex]
  numOfCounts := goutils.ToInt(op(op(maxValue, minValue, "-"), 1, "+"));
  counts := make([]int, numOfCounts)
  
  for i:=low; i<high; i++ {
    curr_index := goutils.ToInt(op(arr[i], minValue, "-"))
    counts[curr_index] = goutils.ToInt(op(counts[curr_index], 1, "+"));
  }
  index:=low;
  fmt.Println(counts)
  for i:=goutils.ToInt(minValue); i<=goutils.ToInt(maxValue); i++ {
    curr_index := goutils.ToInt(op(i, minValue, "-"))
    fmt.Println("i: ", i, "min: ", goutils.ToInt(minValue), "curr_index: ",curr_index)
    for z:=0; z<goutils.ToInt(counts[curr_index]); z++ {
      arr[index] = i;
      index++;
    }
  }
}
