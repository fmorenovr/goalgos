package gosort_test

import(
  "github.com/jenazads/golgos/sort";
  "github.com/jenazads/goutils";
  "fmt";
)

func main() {
  strings := []interface{}{"d", "a", "b", "c"}                  // []
  sortObject,_:=gosort.NewGoSortObject(strings, goutils.StringComparator)
  fmt.Println("Is Ordered ? ",strings)
  fmt.Println("Is Ordered ? ", sortObject.IsSortedPart(1,2))
  sortObject.GoSort() // ["a","b","c","d"]
  fmt.Println("Is Ordered ? ",strings)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted())
  
  arr1 := []interface{}{10, 6, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 10, 3, 5, 8, 8, 3, 2, 7, 4,8, 8, 3, 2, 7, 4, 6, 8, 8, 3, 2, 7, 4,9, 11, 10, 9, 6, 8, 8, 3, 5, 5, 7, 8, 15, 49, 65, 2}
  
  sortObject,_=gosort.NewGoSortObject(arr1, goutils.IntComparator)
  fmt.Println("Is Ordered ? ", arr1)
  sortObject.GnomeSort(0,sortObject.Len())
  fmt.Println("Is Ordered ? ", arr1)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted())
}
