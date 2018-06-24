package distributionsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgs/sort/efficientsort";
  "github.com/jenazads/goalgs/sort/sortfunctions";
)

// PigeonHole Sort
func PigeonHoleSort(arr []interface{}, comp goutils.TypeComparator, op goutils.TypeOperator, low, high int) () {
  maxValueIndex, minValueIndex := sortfunctions.FindMaxMinElementIndex(arr, comp, low, high)
  minValue, maxValue := arr[minValueIndex], arr[maxValueIndex]
  keysRange := goutils.ToInt(op(op(maxValue, minValue, "-"), 1, "+"));
  holes := make([][]interface{}, keysRange)
  
  for i:=low; i<high; i++ {
    curr_index := goutils.ToInt(op(arr[i], minValue, "-"))
    holes[curr_index] = append(holes[curr_index], arr[i])
  }
  index := low;
  for i:= 0; i < keysRange; i++ {
    holeSize := len(holes[i])
    if holeSize > 0 {
      efficientsort.QuickSort(holes[i], comp, 0, holeSize-1)
    }
    for j := 0; j < holeSize; j++ {
      arr[index] = holes[i][j];
      index++;
    }
  }
}
