package efficientsort

import(
  "github.com/jenazads/goutils";
)

// Heap Sort
func HeapSort(arr []interface{}, comp goutils.TypeComparator, low, high int) () {
  // Build heap (rearrange array)
  for i:=high/2-1; i>=low; i-=1 {
    Heapify(arr, comp, i, high);
  }
  // One by one extract an element from heap
  for i:=high-1; i>=low; i-=1 {
    arr[i], arr[low] = arr[low], arr[i]
    // call max heapify on the reduced heap
    Heapify(arr, comp, low, i);
  }
}

// To heapify a subtree rooted with node i which is an index in arr[]. n is size of heap
func Heapify(arr []interface{}, comp goutils.TypeComparator, low, high int){
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
    Heapify(arr, comp, largest, high);
  }
}
