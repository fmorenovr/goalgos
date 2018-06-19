# Golang + Algorithm = goAlg

## Sorting Algorithms

In this section I introduce Sorting Algorithms

### Distribution Sort

* Note: k in Bucket Sort means number of buckets.  
* Note: k in CountingSort means the number of different numbers in array.  
* Note: k in PigeonHoleSort means the number of range number.  
* Note: k in RadixSort means the number of range number and d means the digits of the numbers.  
* Note: k in SpreadSort means the number of range number and d means the digits of the numbers.  

<center>

Algorithm | Worst Case | Average Case | Best case |
--- | --- | ---
BucketSort | O(n².k) | O(n+k) | O(nlog(n)) |
CountingSort | O(n+k) | O(n+k) | O(n+k) |
PigeonHoleSort | O(n+k) | O(n+k) | O(n+k) |
RadixSort | O(n.k/d) | O(n.k/d) | O(nlog(n)) |

Falta
SpreadSort | O(n.(k/s+d)) | O(n.k/d) | O(n) |
BurstSort | O(n.k/d) | O(n.k/d) | O(n) |
FlashSort | O(n²) | O(n.+r) | O(n) |
PostmanSort | O(n.k/d) | O(n.k/d) | O(n) |

</center>


