# Golang + Sort = goSort

Golang implementation of sorting algothims.  
You can see an extended doc in [godocs](https://godoc.org/github.com/Jenazads/golgos/sort).

## Sorting Algorithms

In this section I introduce Sorting Algorithms

### Simple Sort

<center>

| Algorithm | Worst Case | Average Case | Best case |
| --- | --- | --- | --- |
| GnomeSort | O(n²) | O(n²) | O(n) |
| InsertionSort | O(n²) | O(n²) | O(n) |
| SelectionSort | O(n²) | O(n²) | O(n²) |
| SlowSort | O(np) | O(np) | O(n^(log(n)/2)) |
| LibrarySort | O(n²) | O(nlog(n)) | O(n) |
| PatienceSort | O(nlog(n)) | O(nlog(n)) | O(n) |

</center>

### Efficient Sort

<center>

| Algorithm | Worst Case | Average Case | Best case |
| --- | --- | --- | --- |
| MergeSort | O(nlog(n)) | O(nlog(n)) | O(nlog(n)) |
| QuickSort | O(n²) | O(nlog(n)) | O(nlog(n)) |
| HeapSort | O(nlog(n)) | O(nlog(n)) | O(nlog(n)) |
| TimSort | O(nlog(n)) | O(nlog(n)) | O(n) |
| TreeSort | O(n²) | O(nlog(n)) | O(nlog(n)) |
| IntroSort | O(nlog(n)) | O(nlog(n)) | O(nlog(n)) |
| TournamentSort | O(nlog(n)) | O(nlog(n)) | O(nlog(n)) |
| SmoothSort | O(nlog(n)) | O(nlog(n)) | O(n) |
| BlockSort | O(nlog(n)) | O(nlog(n)) | O(n) |

</center>

### Swap Sort

<center>

| Algorithm | Worst Case | Average Case | Best case |
| --- | --- | --- | --- |
| BubbleSort | O(n²) | O(n²) | O(n) |
| CockTailSort | O(n²) | O(n²) | O(n) |
| ShellSort | O(n²) | depends | O(nlog(n)) |
| CombSort | O(n²) | O(n²) | O(nlog(n)) |
| CycleSort | O(n²) | O(n²) | O(n²) |

</center>

### Distribution Sort

* Note: k in Bucket Sort means number of buckets.  
* Note: k in CountingSort means the number of different numbers in array.  
* Note: k in PigeonHoleSort means the number of range number.  
* Note: k in RadixSort means the number of range number and d means the digits of the numbers.  
* Note: k in SpreadSort means the number of range number and d means the digits of the numbers.  

<center>

| Algorithm | Worst Case | Average Case | Best case |
| --- | --- | --- | --- |
| BucketSort | O(n².k) | O(n+k) | O(nlog(n)) |
| CountingSort | O(n+k) | O(n+k) | O(n+k) |
| PigeonHoleSort | O(n+k) | O(n+k) | O(n+k) |
| RadixSort | O(n.k/d) | O(n.k/d) | O(nlog(n)) |
| SpreadSort | O(n.(k/s+d)) | O(n.k/d) | O(n) |
| BurstSort | O(n.k/d) | O(n.k/d) | O(n) |
| FlashSort | O(n²) | O(n.+r) | O(n) |
| PostmanSort | O(n.k/d) | O(n.k/d) | O(n) |

</center>

### Hard Sort

<center>

| Algorithm | Worst Case | Average Case | Best case |
| --- | --- | --- | --- |
| PancakeSort | O(n) | O(n) | O(n) |
| BitonicSort | O(log²(n)) | O(log²(n)) | O(log²(n)) |
| BogoSort | O(Inf) | O(n.n!) | O(n) |
| StoogeSort | O(n^(log(3)/log(1.5))) | O(n^(log(3)/log(1.5))) | O(n^(log(3)/log(1.5))) |
| BeadSort | O(Sum (ai)) | O(Sum (ai)) | O(n) |
| SpaghettiSort | O(n) | O(n) | O(n) |

</center>

### Parallel Sort

<center>

| Algorithm | Worst Case | Average Case | Best case |
| --- | --- | --- | --- |
| SampleSort | O(Sum (ai)) | O(Sum (ai)) | O(n) |
| CubeSort | O(nlog(n)) | O(nlog(n)) | O(n) |

</center>

