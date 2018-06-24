package gosort

import(
  "github.com/jenazads/goalgs/sort/hardsort";
  "github.com/jenazads/goalgs/sort/swapsort";
  "github.com/jenazads/goalgs/sort/simplesort";
  "github.com/jenazads/goalgs/sort/parallelsort";
  "github.com/jenazads/goalgs/sort/efficientsort";
  "github.com/jenazads/goalgs/sort/distributionsort";
)

/**********************
*                     *
*    Simple Sort      *
*                     *
**********************/

// Call SlowSort from SimpleSort package
// Run as SlowSort(begin, end)
func (s* GoSortObject) SlowSort(low, high int){
  simplesort.SlowSort(s.values, s.comparator, low, high-1)
}

// Call GnomeSort from SimpleSort package
// Run as GnomeSort(begin, end)
func (s* GoSortObject) GnomeSort(low, high int){
  simplesort.GnomeSort(s.values, s.comparator, low, high)
}

// Call InsertionSort from SimpleSort package
// Run as InsertionSort(begin, end)
func (s* GoSortObject) InsertionSort(low, high int){
  simplesort.InsertionSort(s.values, s.comparator, low, high)
}

// Call SelectionSort from SimpleSort package
// Run as SelectionSort(begin, end)
func (s* GoSortObject) SelectionSort(low, high int){
  simplesort.SelectionSort(s.values, s.comparator, low, high)
}

// Call LibrarySort from SimpleSort package
// Run as LibrarySort(begin, end)
func (s* GoSortObject) LibrarySort(low, high int){
  simplesort.LibrarySort(s.values, s.comparator, low, high)
}

// Call PatienceSort from SimpleSort package
func (s* GoSortObject) PatienceSort(low, high int){
  simplesort.PatienceSort(s.values, s.comparator, low, high)
}

/**********************
*                     *
*   Efficient Sort    *
*                     *
**********************/

// Call QuickSort from EfficientSort package
// Run as QuickSort(begin, end)
func (s* GoSortObject) QuickSort(low, high int){
  efficientsort.QuickSort(s.values, s.comparator, low, high-1)
}

// Call MergeSort from EfficientSort package
// Run as MergeSort(begin, end)
func (s* GoSortObject) MergeSort(low, high int){
  efficientsort.MergeSort(s.values, s.comparator, low, high-1)
}

// Call HeapSort from EfficientSort package
// Run as HeapSort(begin, end)
func (s* GoSortObject) HeapSort(low, high int){
  efficientsort.HeapSort(s.values, s.comparator, low, high)
}

// Call TimSort from EfficientSort package
// Run as TimSort(begin, end)
func (s* GoSortObject) TimSort(low, high int){
  efficientsort.TimSort(s.values, s.comparator, low, high)
}

/**********************
*                     *
*      Swap Sort      *
*                     *
**********************/

// Call BubbleSort from SwapSort package
// Run as BubbleSort(begin, end)
func (s* GoSortObject) BubbleSort(low, high int){
  swapsort.BubbleSort(s.values, s.comparator, low, high)
}

// Call CockTailSort from SwapSort package
// Run as CockTailSort(begin, end)
func (s* GoSortObject) CockTailSort(low, high int){
  swapsort.CockTailSort(s.values, s.comparator, low, high)
}

// Call ShellSort from SwapSort package
// Run as ShellSort(begin, end)
func (s* GoSortObject) ShellSort(low, high int){
  swapsort.ShellSort(s.values, s.comparator, low, high)
}

// Call CombSort from SwapSort package
// Run as CombSort(begin, end)
func (s* GoSortObject) CombSort(low, high int){
  swapsort.CombSort(s.values, s.comparator, low, high)
}

// Call CycleSort from SwapSort package
// Run as CycleSort(begin, end)
func (s* GoSortObject) CycleSort(low, high int){
  swapsort.CycleSort(s.values, s.comparator, low, high)
}

// Call CircleSort from SwapSort package
// Run as CircleSort(begin, end)
func (s* GoSortObject) CircleSort(low, high int){
  swapsort.CircleSort(s.values, s.comparator, low, high-1)
}

/**********************
*                     *
*  Distribution Sort  *
*                     *
**********************/

// Call BucketSort from DistributionSort package
// Run as BucketSort(begin, end)
func (s* GoSortObject) BucketSort(low, high int){
  distributionsort.BucketSort(s.values, s.comparator, s.operator, low, high)
}

// Call PigeonHoleSort from DistributionSort package
// Run as PigeonHoleSort(begin, end)
func (s* GoSortObject) PigeonHoleSort(low, high int){
  distributionsort.PigeonHoleSort(s.values, s.comparator, s.operator, low, high)
}

// Call CountingSort from DistributionSort package
// Run as CountingSort(begin, end)
func (s* GoSortObject) CountingSort(low, high int){
  distributionsort.CountingSort(s.values, s.comparator, s.operator, low, high)
}

/**********************
*                     *
*     Hard Sort       *
*                     *
**********************/

// Call StoogeSort from HardSort package
// Run as StoogeSort(begin, end)
func (s* GoSortObject) StoogeSort(low, high int){
  hardsort.StoogeSort(s.values, s.comparator, low, high-1)
}

// Call PancakeSort from HardSort package
// Run as PancakeSort(begin, end)
func (s* GoSortObject) PancakeSort(low, high int){
  hardsort.PancakeSort(s.values, s.comparator, low, high)
}

// Call BitonicSort from HardSort package
// Run as BitonicSort(begin, end), (end-high) must be a power of 2.
func (s* GoSortObject) BitonicSort(low, high int){
  hardsort.BitonicSort(s.values, s.comparator, low, high, true)
}

// Call PancakeSort from HardSort package
// Run as PancakeSort(begin, end)
func (s* GoSortObject) BogoSort(low, high int){
  hardsort.BogoSort(s.values, s.comparator, low, high)
}

/**********************
*                     *
*    Parallel Sort    *
*                     *
**********************/

// Call SleepSort from ParallelSort package
// Run as SleepSort(begin, end)
func (s* GoSortObject) SleepSort(low, high int){
  parallelsort.SleepSort(s.values, s.comparator, low, high)
}

// Call BrickSort from ParallelSort package
// Run as BrickSort(begin, end)
func (s* GoSortObject) BrickSort(low, high int){
  parallelsort.BrickSort(s.values, s.comparator, low, high-1)
}
