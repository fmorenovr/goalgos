package gosort

import(
  "sort";
  "github.com/jenazads/goutils"
)

type GoSortObject struct {
  values     []interface{}
  comparator goutils.TypeComparator
}

func NewGoSortObject(arr []interface{}, comp goutils.TypeComparator) (*GoSortObject, error){
  if arr == nil {
    return nil, GoSortEmptyArray
  }
  return &GoSortObject{
    values: arr,
    comparator: comp}, nil
}

// Sort sorts values (in-place) with respect to the given comparator.
func (s *GoSortObject) GoSort() {
  sort.Sort(s)
}

// return True if is sorted
func (s *GoSortObject) IsSorted() bool {
	n := s.Len()
	for i := n - 1; i > 0; i-- {
		if s.Less(i, i-1) {
			return false
		}
	}
	return true
}

// return True if is sorted in this part
func (s *GoSortObject) IsSortedPart(begin, end int) bool {
	for i := end - 1; i > begin; i-- {
		if s.Less(i, i-1) {
			return false
		}
	}
	return true
}

// return the len of array
func (s *GoSortObject) Len() int {
  return len(s.values)
}

// swap elements
func (s *GoSortObject) Swap(i, j int) {
  s.values[i], s.values[j] = s.values[j], s.values[i]
}

// return if is less
func (s *GoSortObject) Less(i, j int) bool {
  return s.comparator(s.values[i], s.values[j]) == -1
}

// return if is less
func (s *GoSortObject) LessThan(i, j int) bool {
  return s.comparator(s.values[i], s.values[j]) == -1
}

// return if is greater
func (s *GoSortObject) GreaterThan(i, j int) bool {
  return s.comparator(s.values[i], s.values[j]) == 1
}

// return if is equals
func (s *GoSortObject) EqualsThan(i, j int) bool {
  return s.comparator(s.values[i], s.values[j]) == 0
}

// return if is less or equals
func (s *GoSortObject) LessEqualsThan(i, j int) bool {
  return (s.comparator(s.values[i], s.values[j]) == -1 || s.comparator(s.values[i], s.values[j]) == 0)
}

// return if is greater or equals
func (s *GoSortObject) GreaterEqualsThan(i, j int) bool {
  return (s.comparator(s.values[i], s.values[j]) == 1 || s.comparator(s.values[i], s.values[j]) == 0)
}

// find the min element and respective index
func (s *GoSortObject) FindMinElementAndIndex() (interface{}, int) {
  var index = 0
  
  for i := 1; i < s.Len(); i++ {
    if s.LessThan(i, index) {
      index = i
    }
  }
  return s.values[index], index
}

// find the max element and respective index
func (s *GoSortObject) FindMaxElementAndIndex() (interface{},int) {
  var index = 0
  
  for i := 1; i < s.Len(); i++ {
    if s.GreaterThan(i, index) {
      index = i
    }
  }
  return s.values[index], index
}

