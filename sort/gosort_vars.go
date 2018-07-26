package gosort

import(
  "github.com/jenazads/goalgos/sort/hardsort";
  "github.com/jenazads/goalgos/sort/swapsort";
  "github.com/jenazads/goalgos/sort/simplesort";
  "github.com/jenazads/goalgos/sort/parallelsort";
  "github.com/jenazads/goalgos/sort/efficientsort";
  "github.com/jenazads/goalgos/sort/distributionsort";
)

/**********************
*                     *
*    Simple Sort      *
*                     *
**********************/

// Call SlowSort from SimpleSort package
// Run as SlowSort(begin, end)
func (s* GoSort) SlowSort(low, high int){
  simplesort.SlowSort(s.values, s.comparator, low, high-1)
}

// Call SillySort from SimpleSort package
// Run as SillySort(begin, end)
func (s* GoSort) SillySort(low, high int){
  simplesort.SillySort(s.values, s.comparator, low, high-1)
}

// Call GnomeSort from SimpleSort package
// Run as GnomeSort(begin, end)
func (s* GoSort) GnomeSort(low, high int){
  simplesort.GnomeSort(s.values, s.comparator, low, high)
}

// Call InsertionSort from SimpleSort package
// Run as InsertionSort(begin, end)
func (s* GoSort) InsertionSort(low, high int){
  simplesort.InsertionSort(s.values, s.comparator, low, high)
}

// Call SelectionSort from SimpleSort package
// Run as SelectionSort(begin, end)
func (s* GoSort) SelectionSort(low, high int){
  simplesort.SelectionSort(s.values, s.comparator, low, high)
}
/*
// Call LibrarySort from SimpleSort package
// Run as LibrarySort(begin, end)
func (s* GoSort) LibrarySort(low, high int){
  simplesort.LibrarySort(s.values, s.comparator, low, high)
}

// Call PatienceSort from SimpleSort package
// Run as PatienceSort(begin, end)
func (s* GoSort) PatienceSort(low, high int){
  simplesort.PatienceSort(s.values, s.comparator, low, high)
}
*/
/**********************
*                     *
*   Efficient Sort    *
*                     *
**********************/

// Call QuickSort from EfficientSort package
// Run as QuickSort(begin, end)
func (s* GoSort) QuickSort(low, high int){
  efficientsort.QuickSort(s.values, s.comparator, low, high-1)
}

// Call MergeSort from EfficientSort package
// Run as MergeSort(begin, end)
func (s* GoSort) MergeSort(low, high int){
  efficientsort.MergeSort(s.values, s.comparator, low, high-1)
}

// Call IntroSort from EfficientSort package
// Run as IntroSort(begin, end)
func (s* GoSort) IntroSort(low, high int){
  efficientsort.IntroSort(s.values, s.comparator, low, high-1)
}

// Call HeapSort from EfficientSort package
// Run as HeapSort(begin, end)
func (s* GoSort) HeapSort(low, high int){
  efficientsort.HeapSort(s.values, s.comparator, low, high)
}

// Call TimSort from EfficientSort package
// Run as TimSort(begin, end)
func (s* GoSort) TimSort(low, high int){
  efficientsort.TimSort(s.values, s.comparator, low, high)
}

// Call TreeSort from EfficientSort package
// Run as TreeSort(begin, end)
func (s* GoSort) TreeSort(low, high int, f func(interface{})(interface{})){
  efficientsort.TreeSort(s.values, s.comparator, s.operator, f, low, high)
}

/**********************
*                     *
*      Swap Sort      *
*                     *
**********************/

// Call BubbleSort from SwapSort package
// Run as BubbleSort(begin, end)
func (s* GoSort) BubbleSort(low, high int){
  swapsort.BubbleSort(s.values, s.comparator, low, high)
}

// Call CockTailSort from SwapSort package
// Run as CockTailSort(begin, end)
func (s* GoSort) CockTailSort(low, high int){
  swapsort.CockTailSort(s.values, s.comparator, low, high)
}

// Call ShellSort from SwapSort package
// Run as ShellSort(begin, end)
func (s* GoSort) ShellSort(low, high int){
  swapsort.ShellSort(s.values, s.comparator, low, high)
}

// Call CombSort from SwapSort package
// Run as CombSort(begin, end)
func (s* GoSort) CombSort(low, high int){
  swapsort.CombSort(s.values, s.comparator, low, high)
}

// Call CycleSort from SwapSort package
// Run as CycleSort(begin, end)
func (s* GoSort) CycleSort(low, high int){
  swapsort.CycleSort(s.values, s.comparator, low, high)
}

// Call CircleSort from SwapSort package
// Run as CircleSort(begin, end)
func (s* GoSort) CircleSort(low, high int){
  swapsort.CircleSort(s.values, s.comparator, low, high)
}

// Call PermutationSort from SwapSort package
// Run as PermutationSort(begin, end)
func (s* GoSort) PermutationSort(low, high int){
  swapsort.PermutationSort(s.values, s.comparator, low, high)
}

/**********************
*                     *
*  Distribution Sort  *
*                     *
**********************/

// Call BucketSort from DistributionSort package
// Run as BucketSort(begin, end)
func (s* GoSort) BucketSort(low, high int){
  distributionsort.BucketSort(s.values, s.comparator, s.operator, low, high)
}

// Call BucketSort from DistributionSort package
// Run as BucketSort(begin, end)
func (s* GoSort) ShuffleSort(low, high int){
  distributionsort.ShuffleSort(s.values, s.comparator, s.operator, low, high)
}

// Call PigeonHoleSort from DistributionSort package
// Run as PigeonHoleSort(begin, end)
func (s* GoSort) PigeonHoleSort(low, high int){
  distributionsort.PigeonHoleSort(s.values, s.comparator, s.operator, low, high)
}

// Call CountingSort from DistributionSort package
// Run as CountingSort(begin, end)
func (s* GoSort) CountingSort(low, high int){
  distributionsort.CountingSort(s.values, s.comparator, s.operator, low, high)
}

/**********************
*                     *
*     Hard Sort       *
*                     *
**********************/

// Call StoogeSort from HardSort package
// Run as StoogeSort(begin, end)
func (s* GoSort) StoogeSort(low, high int){
  hardsort.StoogeSort(s.values, s.comparator, low, high-1)
}

// Call PancakeSort from HardSort package
// Run as PancakeSort(begin, end)
func (s* GoSort) PancakeSort(low, high int){
  hardsort.PancakeSort(s.values, s.comparator, low, high)
}

// Call BitonicSort from HardSort package
// Run as BitonicSort(begin, end), (end-high) must be a power of 2.
func (s* GoSort) BitonicSort(low, high int){
  hardsort.BitonicSort(s.values, s.comparator, low, high, true)
}

// Call PancakeSort from HardSort package
// Run as PancakeSort(begin, end)
func (s* GoSort) BogoSort(low, high int){
  hardsort.BogoSort(s.values, s.comparator, low, high)
}

/**********************
*                     *
*    Parallel Sort    *
*                     *
**********************/

// Call BrickSort from ParallelSort package
// Run as BrickSort(begin, end)
func (s* GoSort) BrickSort(low, high int){
  parallelsort.BrickSort(s.values, s.comparator, low, high-1)
}

// Call SleepSort from ParallelSort package
// Run as SleepSort(begin, end)
func (s* GoSort) SleepSort(low, high int){
  parallelsort.SleepSort(s.values, s.comparator, low, high)
}
