package gosort_test

import(
  "github.com/jenazads/golgos/sort";
  "fmt";
)

func main() {
  strings := []interface{}{}                  // []
  strings = append(strings, "d")              // ["d"]
  strings = append(strings, "a")              // ["d","a"]
  strings = append(strings, "b")              // ["d","a",b"
  strings = append(strings, "c")              // ["d","a",b","c"]
  sortObject,_:=gosort.NewGoSortObject(strings, gosort.StringComparator)
  fmt.Println("Is Ordered ? ",strings)
  fmt.Println("Is Ordered ? ", sortObject.IsSortedPart(1,2))
  sortObject.GoSort() // ["a","b","c","d"]
  fmt.Println("Is Ordered ? ",strings)
  fmt.Println("Is Ordered ? ", sortObject.IsSorted())
}
