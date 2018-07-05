package distributionsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgos/sort/simplesort";
  "github.com/jenazads/goalgos/sort/sortfunctions";
)

// PigeonHole Sort
func PigeonHoleSort(arr []interface{}, comp goutils.TypeComparator, op goutils.TypeOperator, low, high int) () {
  if(low<high){
    minValue, maxValue := sortfunctions.FindMaxMinElementIndex(arr, comp, low, high)
    numOfHoles := goutils.ToInt(op(op(maxValue, minValue, "-"), 1, "+"));
    holes := make([][]interface{}, numOfHoles)
    
    for i:=low; i<high; i++ {
      curr_index := goutils.ToInt(op(arr[i], minValue, "-"))
      holes[curr_index] = append(holes[curr_index], arr[i])
    }
    index := low;
    for i:= 0; i < numOfHoles; i++ {
      holeLength := len(holes[i])
      if holeLength > 0 {
        simplesort.InsertionSort(holes[i], comp, 0, holeLength)
      }
      for j := 0; j < holeLength; j++ {
        arr[index] = holes[i][j];
        index++;
      }
    }
  }
}
