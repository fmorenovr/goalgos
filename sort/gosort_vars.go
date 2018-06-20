package gosort

import(
  "github.com/jenazads/golgos/sort/simplesort";
  "github.com/jenazads/golgos/sort/efficientsort";
)

/**********************
*                     *
*    Simple Sort      *
*                     *
**********************/

// Call SlowSort in SimpleSort package
// Run as SlowSort(begin, end-1)
func (s* GoSortObject) SlowSort(low, high int){
  simplesort.SlowSort(s.values, s.comparator, low, high)
}

// Call GnomeSort in SimpleSort package
// Run as GnomeSort(begin, end)
func (s* GoSortObject) GnomeSort(low, high int){
  simplesort.GnomeSort(s.values, s.comparator, low, high)
}

// Call InsertionSort in SimpleSort package
// Run as InsertionSort(begin, end)
func (s* GoSortObject) InsertionSort(low, high int){
  simplesort.InsertionSort(s.values, s.comparator, low, high)
}

// Call SelectionSort in SimpleSort package
// Run as SelectionSort(begin, end)
func (s* GoSortObject) SelectionSort(low, high int){
  simplesort.SelectionSort(s.values, s.comparator, low, high)
}

// Call LibrarySort in SimpleSort package
// Run as LibrarySort(begin, end)
func (s* GoSortObject) LibrarySort(low, high int){
  simplesort.LibrarySort(s.values, s.comparator, low, high)
}

// Call PatienceSort in SimpleSort package
func (s* GoSortObject) PatienceSort(low, high int){
  simplesort.PatienceSort(s.values, s.comparator, low, high)
}


/**********************
*                     *
*   Efficient Sort    *
*                     *
**********************/

// Call QuickSort in EfficientSort package
// Run as QuickSort(begin, end-1)
func (s* GoSortObject) QuickSort(low, high int){
  efficientsort.QuickSort(s.values, s.comparator, low, high)
}

// Call MergeSort in EfficientSort package
// Run as MergeSort(begin, end-1)
func (s* GoSortObject) MergeSort(low, high int){
  efficientsort.MergeSort(s.values, s.comparator, low, high)
}

// Call HeapSort in EfficientSort package
// Run as HeapSort(begin, end)
func (s* GoSortObject) HeapSort(low, high int){
  efficientsort.HeapSort(s.values, s.comparator, low, high)
}
