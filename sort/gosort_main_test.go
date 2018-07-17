package gosort_test

import(
  "github.com/jenazads/goalgos/sort";
  "github.com/jenazads/goutils";
  "fmt";
)

func main() {
  
  begin:=24
  end:=34

  strings := []interface{}{"d", "a", "b", "c"}                  // []
  sortObject,_:=gosort.NewGoSort(strings, goutils.StringComparator)
  fmt.Println("Is Ordered ? ",strings)
  fmt.Println("Is Ordered ? ", sortObject.IsSortedPart(1,2))
  sortObject.GoSort() // ["a","b","c","d"]
  fmt.Println("Is Ordered ? ",strings)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr := []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("GnomeSort")
  sortObject.GnomeSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("InsertionSort")
  sortObject.InsertionSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("SlowSort")
  sortObject.SlowSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("SelectionSort")
  sortObject.SelectionSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("SillySort")
  sortObject.SillySort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("QuickSort")
  sortObject.QuickSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("MergeSort")
  sortObject.MergeSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("HeapSort -> revisar el sub array")
  sortObject.HeapSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("TimSort")
  sortObject.TimSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("BubbleSort")
  sortObject.BubbleSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("CockTailSort")
  sortObject.CockTailSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("ShellSort")
  sortObject.ShellSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("CombSort")
  sortObject.CombSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("CycleSort")
  sortObject.CycleSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("CircleSort")
  sortObject.CircleSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("BrickSort")
  sortObject.BrickSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSort(arr, goutils.IntComparator)
  fmt.Println("SleepSort -> revisar el algoritmo")
  sortObject.SleepSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSortOp(arr, goutils.IntComparator, goutils.IntOperator)
  fmt.Println("BucketSort")
  sortObject.BucketSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSortOp(arr, goutils.IntComparator, goutils.IntOperator)
  fmt.Println("PigeonHoleSort")
  sortObject.PigeonHoleSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSortOp(arr, goutils.IntComparator, goutils.IntOperator)
  fmt.Println("CountingSort")
  sortObject.CountingSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSortOp(arr, goutils.IntComparator, goutils.IntOperator)
  fmt.Println("StoogeSort")
  sortObject.StoogeSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSortOp(arr, goutils.IntComparator, goutils.IntOperator)
  fmt.Println("PancakeSort")
  sortObject.PancakeSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4, 8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSortOp(arr, goutils.IntComparator, goutils.IntOperator)
  fmt.Println("IntroSort")
  sortObject.IntroSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSortOp(arr, goutils.IntComparator, goutils.IntOperator)
  fmt.Println("BitonicSort")
  sortObject.BitonicSort(0,sortObject.Len())
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4, 8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSortOp(arr, goutils.IntComparator, goutils.IntOperator)
  fmt.Println("PermutationSort")
  sortObject.PermutationSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
  
  arr = []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4, 8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSortOp(arr, goutils.IntComparator, goutils.IntOperator)
  fmt.Println("BogoSort")
  sortObject.BogoSort(begin,sortObject.Len()-end)
  fmt.Println("Is Ordered ? ", arr)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted(), "\n\n")
}
