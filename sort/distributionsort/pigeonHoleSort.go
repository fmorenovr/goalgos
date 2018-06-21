package distributionsort

import(
  "github.com/jenazads/goutils";
)

// PigeonHole Sort
func PigeonHoleSort(arr []interface{}, comp goutils.TypeComparator, op goutils.TypeOperator, low, high int) () {
  maxValueIndex, minValueIndex := FindMaxMinElementIndex(arr, comp, low, high)
  minValue, maxValue := arr[minValueIndex], arr[maxValueIndex]
  keysRange := goutils.ToInt(op(op(maxValue, minValue, "-"), 1, "+"));
  holes := make([][]interface{}, keysRange)
  
  for i:=low; i<high; i++ {
    index_curr := goutils.ToInt(op(arr[i], minValue, "-"))
    holes[index_curr] = append(holes[index_curr], arr[i])
  }
  index := low;
  for i:= 0; i < keysRange; i++ {
    holeSize := len(holes[i])
    for j := 0; j < holeSize; j++ {
      arr[index] = holes[i][j];
      index++;
    }
  }
}
