package swapsort

import(
  "fmt";
  "github.com/jenazads/goutils";
)

// Cycle Sort
func CycleSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low < high {
    writes := 0;
    for cycle_start:=low; cycle_start <= high - 2; cycle_start++ {
      item := arr[cycle_start];
      pos := cycle_start;
      for i := cycle_start + 1; i < high; i++ {
        if comp(arr[i], item) == -1 {
          pos++;
        }
      }
      if pos == cycle_start{
        continue;
      }
      for item == arr[pos] {
        pos += 1;
      }
      if pos != cycle_start {
        arr[pos], item = item, arr[pos];
        writes++;
      }
      for pos != cycle_start {
        pos = cycle_start;
        for i := cycle_start + 1; i < high; i++ {
          if comp(arr[i], item) == -1{
            pos += 1;
          }
        }
        for comp(item, arr[pos])==0 {
          pos += 1;
        }
        if !(comp(item, arr[pos]) ==0) {
          arr[pos], item = item, arr[pos]
          writes++;
        }
      }
    }
    fmt.Println("Number of memory writes or swaps: ", writes);
  }
}
