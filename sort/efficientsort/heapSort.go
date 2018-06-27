package efficientsort

import(
  "github.com/jenazads/goutils";
)

// Heap Sort
func HeapSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  BuildMaxHeap(arr,comp,low,high);
  // One by one extract an element from heap
  for i:=high-1; i>=low; i-=1 {
    arr[i], arr[low] = arr[low], arr[i]
    // call max heapify on the reduced heap
    MaxHeapify(arr, comp, low, i);
  }
}

// Build Max heap (rearrange array)
func BuildMaxHeap(arr []interface{}, comp goutils.TypeComparator, low, high int){
  size:=high-low
  for i:=size/2-1; i>=low; i-=1 {
    MaxHeapify(arr, comp, i, high);
  }
}

// Build Min heap (rearrange array)
func BuildMinHeap(arr []interface{}, comp goutils.TypeComparator, low, high int){
  size:=high-low
  for i:=size/2-1; i>=low; i-=1 {
    MinHeapify(arr, comp, i, high);
  }
}

// To max heapify a subtree rooted with node i which is an index in arr[]
func MaxHeapify(arr []interface{}, comp goutils.TypeComparator, low, high int){
  largest := low  // Initialize largest as root
  left := 2*low + 1
  right := 2*low + 2

  // If left child is larger than root
  if left < high && comp(arr[left], arr[largest]) == 1 {
    largest = left;
  }
  // If right child is larger than largest so far
  if right < high && comp(arr[right], arr[largest]) == 1 {
    largest = right;
  }
  // If largest is not root
  if largest != low {
    arr[largest], arr[low] = arr[low], arr[largest]
    // Recursively heapify the affected sub-tree
    MaxHeapify(arr, comp, largest, high);
  }
}

// To min heapify a subtree rooted with node i which is an index in arr[]
func MinHeapify(arr []interface{}, comp goutils.TypeComparator, low, high int){
  smallest := low  // Initialize largest as root
  left := 2*low + 1
  right := 2*low + 2

  // If left child is larger than root
  if left < high && comp(arr[left], arr[smallest]) == -1 {
    smallest = left;
  }
  // If right child is larger than largest so far
  if right < high && comp(arr[right], arr[smallest]) == -1 {
    smallest = right;
  }
  // If largest is not root
  if smallest != low {
    arr[smallest], arr[low] = arr[low], arr[smallest]
    // Recursively heapify the affected sub-tree
    MinHeapify(arr, comp, smallest, high);
  }
}
