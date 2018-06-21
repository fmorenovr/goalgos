package parallelsort

import(
  "time";
  "github.com/jenazads/goutils";
)

// Sleep Sort or Lazy Sort
func SleepSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
	out := make(chan interface{})
  for i:=low; i<high; i++ {
    go func(index int, arr []interface{}) {
      time.Sleep(time.Duration(arr[index].(int)) * time.Millisecond)
      out <- arr[index]
    }(i, arr)
  }
  for i:=low; i<high; i++ {
    arr[i] = <- out
  }
}
