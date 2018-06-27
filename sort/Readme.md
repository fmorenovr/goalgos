# Golang + Sort = goSort

Golang implementation of sorting algothims.  
You can see an extended doc in [godocs](https://godoc.org/github.com/Jenazads/golgos/sort).

## Sorting Algorithms

In this section I introduce Sorting Algorithms

### Simple Sort

<center>

| Algorithm | Best Case | Average Case | Worst case |
| --- | --- | --- | --- |
| GnomeSort | O(n) | O(n²) | O(n²) |
| InsertionSort | O(n) | O(n²) | O(n²) |
| SelectionSort | O(n²) | O(n²) | O(n²) |
| SlowSort | O(n^(log(n)/2)) | O(np) | O(np) |
| LibrarySort | O(n) | O(nlog(n)) | O(n²) |
| PatienceSort | O(n) | O(nlog(n)) | O(nlog(n)) |

</center>

### Efficient Sort

<center>

| Algorithm | Best Case | Average Case | Worst case |
| --- | --- | --- | --- |
| MergeSort | O(nlog(n)) | O(nlog(n)) | O(nlog(n)) |
| QuickSort | O(nlog(n)) | O(nlog(n)) | O(n²) |
| HeapSort | O(nlog(n)) | O(nlog(n)) | O(nlog(n)) |
| TimSort | O(n) | O(nlog(n)) | O(nlog(n)) |
| TreeSort | O(n) | O(nlog(n)) | O(nlog(n)) |
| IntroSort | O(nlog(n)) | O(nlog(n)) | O(nlog(n)) |
| TournamentSort | O(nlog(n)) | O(nlog(n)) | O(nlog(n)) |
| SmoothSort | O(n) | O(nlog(n)) | O(nlog(n)) |
| BlockSort | O(n) | O(nlog(n)) | O(nlog(n)) |
| StrandSort | O(n) | O(n²) | O(n²) |

</center>

### Swap Sort

<center>

| Algorithm | Best Case | Average Case | Worst case |
| --- | --- | --- | --- |
| BubbleSort | O(n) | O(n²) | O(n²) |
| CockTailSort | O(n) | O(n²) | O(n²) |
| ShellSort | O(n) | O(n(log(n))²) | O((nlog(n))²) |
| CombSort | O(nlog(n)) | O(n²) | O(n²) |
| CycleSort | O(n²) | O(n²) | O(n²) |
| CircleSort | O() | O() | O() |
| PermutationSort | O() | O() | O() |

</center>

### Distribution Sort

* Note: k in Bucket Sort means number of buckets.  
* Note: k in CountingSort means the number of different numbers in array.  
* Note: k in PigeonHoleSort means the number of range number.  
* Note: k in RadixSort means the number of range number.
* Note: k in SpreadSort means the number of range number and d means the digits of the numbers.  
* Notes: Couting sort and Radix sort is only for integers.

<center>

| Algorithm | Best Case | Average Case | Worst case |
| --- | --- | --- | --- |
| BucketSort | O(n+k) | O(n+k) | O(n²) |
| CountingSort | O(n+k) | O(n+k) | O(n+k) |
| PigeonHoleSort | O(n+k) | O(n+k) | O(n+k) |
| RadixSort | O(n.k) | O(n.k) | O(nk) |
| SpreadSort | O(n) | O(n.k/d) | O(n.(k/s+d)) |
| BurstSort | O(n) | O(n.k/d) | O(n.k/d) |
| FlashSort | O(n) | O(n.+r) | O(n²) |
| PostmanSort | O(n) | O(n.k/d) | O(n.k/d) |
| ProxmapSort | O(n) | O(n) | O(n²) |

</center>

### Hard Sort

<center>

| Algorithm | Best Case | Average Case | Worst case |
| --- | --- | --- | --- |
| PancakeSort | O(n) | O(n) | O(n) |
| BitonicSort | O(log²(n)) | O(log²(n)) | O(log²(n)) |
| BogoSort | O(n) | O(n.n!) | O(Inf) |
| StoogeSort | O(n^(log(3)/log(1.5))) | O(n^(log(3)/log(1.5))) | O(n^(log(3)/log(1.5))) |
| BeadSort | O(n) | O(Sum (ai)) | O(Sum (ai)) |
| SpaghettiSort | O(n) | O(n) | O(n) |
| NetworkSort | O(log²(n)) | O(log²(n)) | O(log²(n)) |

</center>

### Parallel Sort

* Note: t in SleepSort means the time of seconds sleeped by the number value.  

<center>

| Algorithm | Best Case | Average Case | Worst case |
| --- | --- | --- | --- |
| BrickSort | O(n) | O(n²) | O(n²) |
| SleepSort | O(nt) | O(nt) | O(nt) |
| SampleSort | O(n) | O(Sum (ai)) | O(Sum (ai)) |
| CubeSort | O(n) | O(nlog(n)) | O(nlog(n)) |

</center>

