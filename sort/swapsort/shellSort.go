package swapsort

import(
  "github.com/jenazads/goutils";
)

// Shell Sort
func ShellSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  if low < high {
    size := high - low;
    for gap := size/2; gap > 0; gap/=2 {
      for i := gap+low; i < high; i++ {
        j := i;
        temp := arr[i];
        for (j >= gap+low) && (comp(arr[j-gap], temp) == 1) {
          arr[j] = arr[j - gap];
          j -= gap;
        }
        arr[j] = temp;
      }
    }
  }
}
