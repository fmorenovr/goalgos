package gosort

import(
  "sort";
  "github.com/jenazads/goutils"
)

type GoSort struct {
  values     []interface{}
  comparator goutils.TypeComparator
  operator   goutils.TypeOperator
}

// new GoSort 
func NewGoSort(arr []interface{}, comp goutils.TypeComparator) (*GoSort, error){
  if arr == nil {
    return nil, GoSortEmptyArray
  }
  return &GoSort{
    values: arr,
    comparator: comp,
    operator: nil}, nil
}

func NewGoSortOp(arr []interface{}, comp goutils.TypeComparator, op goutils.TypeOperator) (*GoSort, error){
  if arr == nil {
    return nil, GoSortEmptyArray
  }
  return &GoSort{
    values: arr,
    comparator: comp,
    operator: op}, nil
}

// Sort sorts values (in-place) with respect to the given comparator.
func (s *GoSort) GoSort() {
  sort.Sort(s)
}

// return True if is sorted
func (s *GoSort) IsSorted() bool {
	n := s.Len()
	for i := n - 1; i > 0; i-- {
		if s.Less(i, i-1) {
			return false
		}
	}
	return true
}

// return sub array
func (s *GoSort) SubArray(low, high int) (*GoSort) {
  gosortObj, _ := NewGoSortOp(s.values[low:high], s.comparator, s.operator)
  return gosortObj
}

// return True if is sorted in this part
func (s *GoSort) IsSortedPart(begin, end int) bool {
	for i := end - 1; i > begin; i-- {
		if s.Less(i, i-1) {
			return false
		}
	}
	return true
}

// return the len of array
func (s *GoSort) Len() int {
  return len(s.values)
}

// swap elements
func (s *GoSort) Swap(i, j int) {
  s.values[i], s.values[j] = s.values[j], s.values[i]
}

// return if is less
func (s *GoSort) Less(i, j int) bool {
  return s.comparator(s.values[i], s.values[j]) == -1
}

// return if is less
func (s *GoSort) LessThan(i, j int) bool {
  return s.comparator(s.values[i], s.values[j]) == -1
}

// return if is greater
func (s *GoSort) GreaterThan(i, j int) bool {
  return s.comparator(s.values[i], s.values[j]) == 1
}

// return if is equals
func (s *GoSort) EqualsThan(i, j int) bool {
  return s.comparator(s.values[i], s.values[j]) == 0
}

// return if is less or equals
func (s *GoSort) LessEqualsThan(i, j int) bool {
  return (s.comparator(s.values[i], s.values[j]) == -1 || s.comparator(s.values[i], s.values[j]) == 0)
}

// return if is greater or equals
func (s *GoSort) GreaterEqualsThan(i, j int) bool {
  return (s.comparator(s.values[i], s.values[j]) == 1 || s.comparator(s.values[i], s.values[j]) == 0)
}

// find the min element and respective index
func (s *GoSort) FindMinElementAndIndex() (interface{}, int) {
  var index = 0
  
  for i := 1; i < s.Len(); i++ {
    if s.LessThan(i, index) {
      index = i
    }
  }
  return s.values[index], index
}

// find the max element and respective index
func (s *GoSort) FindMaxElementAndIndex() (interface{},int) {
  var index = 0
  
  for i := 1; i < s.Len(); i++ {
    if s.GreaterThan(i, index) {
      index = i
    }
  }
  return s.values[index], index
}

