// 2. Efficent Sort
// Tim Sort -> run as timSort(arr,size);
void timSort(int *arr, int size, int RUN = 32){
  // Sort individual subarrays of sizeRUN
  for (int i = 0; i < size; i+=RUN)
    insertionSort(arr, i, min((i+31), (size-1)));
  // start merging from sizeRUN (or 32). It will merge
  // to form sizeRUN 64, then 128, 256 and so on ....
  for (int sizeRUN = RUN; sizeRUN < size; sizeRUN = 2*sizeRUN){
    // pick starting point of left sub array. We are going to merge arr[left..left+sizeRUN-1]
    // and arr[left+sizeRUN, left+2*sizeRUN-1]. After every merge, we increase left by 2*sizeRUN
    for (int left = 0; left < size; left += 2*sizeRUN){
      // find ending point of left sub array mid+1 is starting point of right sub array
      int mid = left + sizeRUN - 1;
      int right = min((left + 2*sizeRUN - 1), (size-1));
      // merge sub array arr[left.....mid] & arr[mid+1....right]
      merge(arr, left, mid, right);
    }
  }
}

// Tree Sort -> run as treeSort(arr,size);
void treeSort(int *arr, int size){
  struct Node *root = NULL;

  // Construct the BST
  root = insert(root, arr[0]);
  for (int i=1; i<size; i++)
    insert(root, arr[i]);

  // Store inoder traversal of the BST
  // in arr[]
  int i = 0;
  storeSorted(root, arr, i);
}

// Stores inoder traversal of the BST in arr[]
void storeSorted(Node *root, int *arr, int &i){
  if (root != NULL){
    storeSorted(root->left, arr, i);
    arr[i++] = root->key;
    storeSorted(root->right, arr, i);
  }
}

// Struct Node
struct Node{
  int key;
  struct Node *left, *right;
};

// A utility function to create a new BST Node
struct Node *newNode(int item){
  struct Node *temp = new Node;
  temp->key = item;
  temp->left = temp->right = NULL;
  return temp;
}

// insert function
Node* insert(Node* node, int key){
  /* If the tree is empty, return a new Node */
  if (node == NULL) return newNode(key);

  /* Otherwise, recur down the tree */
  if (key < node->key)
    node->left  = insert(node->left, key);
  else if (key > node->key)
    node->right = insert(node->right, key);

  /* return the (unchanged) Node pointer */
  return node;
}

// 3. Bubble and variants
// Bubble Sort -> run as bubbleSort(arr,size);
void bubbleSort(int *arr, int size){
  // Base case
  if (size == 1)
      return;
  // One pass of bubble sort. After this pass, the largest element is moved (or bubbled) to end.
  for (int i=0; i<size-1; i++)
    if (arr[i] > arr[i+1])
      swap(arr[i], arr[i+1]);
  // Largest element is fixed,
  // recur for remaining array
  bubbleSort(arr, size-1);
}

// CockTail Sort -> run as cockTailSort(arr,size);
void cockTailSort(int *arr, int size){
  bool swapped = true;
  int start = 0;
  int end = size - 1;
  while (swapped) {
    // reset the swapped flag on entering
    // the loop, because it might be true from
    // a previous iteration.
    swapped = false;
    // loop from left to right same as
    // the bubble sort
    for (int i = start; i < end; ++i) {
      if (arr[i] > arr[i + 1]) {
        swap(arr[i], arr[i + 1]);
        swapped = true;
      }
    }
    // if nothing moved, then array is sorted.
    if (!swapped)
      break;
    // otherwise, reset the swapped flag so that it
    // can be used in the next stage
    swapped = false;
    // move the end point back by one, because
    // item at the end is in its rightful spot
    --end;
    // from right to left, doing the
    // same comparison as in the previous stage
    for (int i = end - 1; i >= start; --i) {
      if (arr[i] > arr[i + 1]) {
        swap(arr[i], arr[i+1]);
        swapped = true;
      }
    }
    // increase the starting point, because
    // the last stage would have moved the next
    // smallest number to its rightful spot.
    ++start;
  }
}

// Shell Sort -> run as shellSort(arr,size);
void shellSort(int *arr, int n){
  // Start with a big gap, then reduce the gap
  for (int gap = n/2; gap > 0; gap /= 2){
    // Do a gapped insertion sort for this gap size.
    // The first gap elements a[0..gap-1] are already in gapped order
    // keep adding one more element until the entire array is
    // gap sorted 
    for (int i = gap; i < n; i += 1){
      // add a[i] to the elements that have been gap sorted
      // save a[i] in temp and make a hole at position i
      int temp = arr[i];
      // shift earlier gap-sorted elements up until the correct 
      // location for a[i] is found
      int j;            
      for (j = i; j >= gap && arr[j - gap] > temp; j -= gap)
        arr[j] = arr[j - gap];
      //  put temp (the original a[i]) in its correct location
      arr[j] = temp;
    }
  }
}

// Comb Sort -> run as combSort(arr,size);
void combSort(int *arr, int size){
  // Initialize gap
  int gap = size;
  // Initialize swapped as true to make sure that
  // loop runs
  bool swapped = true;
  // Keep running while gap is more than 1 and last
  // iteration caused a swap
  while (gap != 1 || swapped == true){
    // Find next gap
    gap = getNextGap(gap);
    // Initialize swapped as false so that we can
    // check if swap happened or not
    swapped = false;
    // Compare all elements with current gap
    for (int i=0; i<size-gap; i++){
      if (arr[i] > arr[i+gap]){
        swap(arr[i], arr[i+gap]);
        swapped = true;
      }
    }
  }
}

// To find gap between elements
int getNextGap(int gap){
  // Shrink gap by Shrink factor
  gap = (gap*10)/13;
  if (gap < 1)
    return 1;
  return gap;
}

// Cycle Sort -> run as cycleSort(arr,size);
void cycleSort(int *arr, int size){
  // count number of memory writes
  int writes = 0;
  // traverse array elements and put it to on
  // the right place
  for (int cycle_start = 0; cycle_start <= size - 2; cycle_start++) {
    // initialize item as starting point
    int item = arr[cycle_start];
    // Find position where we put the item. We basically
    // count all smaller elements on right side of item.
    int pos = cycle_start;
    for (int i = cycle_start + 1; i < size; i++)
      if (arr[i] < item)
        pos++;
    // If item is already in correct position
    if (pos == cycle_start)
      continue;
    // ignore all duplicate  elements
    while (item == arr[pos])
      pos += 1;
    // put the item to it's right position
    if (pos != cycle_start) {
      swap(item, arr[pos]);
      writes++;
    }

    // Rotate rest of the cycle
    while (pos != cycle_start) {
      pos = cycle_start;
      // Find position where we put the element
      for (int i = cycle_start + 1; i < size; i++)
        if (arr[i] < item)
          pos += 1;

      // ignore all duplicate  elements
      while (item == arr[pos])
        pos += 1;

      // put the item to it's right position
      if (item != arr[pos]) {
        swap(item, arr[pos]);
        writes++;
      }
    }
  }
  // Number of memory writes or swaps
  cout << "Number of memory writes or swaps: " << writes << endl ;
}

// 4. Distribution Sort
// Bucket Sort -> run as bucketSort(arr,size);
void bucketSort(int *arr, int count){
  int minValue = arr[0];
  int maxValue = arr[0];
  // get min and max value
  maxValue = findMaxElement(arr, count);
  minValue = findMinElement(arr, count);
  // Set leng for each bucket
  int bucketLength = maxValue - minValue + 1;
  vector<int>* bucket = new vector<int>[bucketLength];

  for (int i = 0; i < bucketLength; i++){
    bucket[i] = vector<int>();
  }
  // set arr[i] in respective bucket
  for (int i = 0; i < count; i++){
    bucket[arr[i] - minValue].push_back(arr[i]);
  }
  //
  int k = 0;
  for (int i = 0; i < bucketLength; i++){
    int bucketSize = bucket[i].size();

    if (bucketSize > 0){
      for (int j = 0; j < bucketSize; j++){
        arr[k] = bucket[i][j];
        k++;
      }
    }
  }
}

// Counting Sort -> run as countingSort(arr,size);
void countingSort(int *arr, int size, int RANGE=32, int radix=0, int expo=0){
  // RANGE depends of the max number in array
  // if radix = 1, means that main sort is Radix Sort
  // expo depends of de max number in array in Radix Sort

  // Create a count array to store count of inidividul
  // characters and initialize count array as 0
  int i, count[RANGE]={0};
  // The output character array that will have sorted arr
  int out[size];

  // Store count of each character
  for(i=0;i<size;i++)
    if(radix) count[(arr[i]/expo)%10]++;
    else count[arr[i]]++;
  // Change count[i] so that count[i] now contains actual
  // position of this character in output array
  for(i=1;i<RANGE;i++)
    count[i]+=count[i-1];
  // Build the output character array
  for(i=size-1;i>=0;i--){
    if(radix){
      out[count[(arr[i]/expo)%10] - 1] = arr[i];
      count[ (arr[i]/expo)%10 ]--;
    } else {
      out[count[arr[i]]-1]=arr[i];
      count[arr[i]]--;
    }
  }
  // Copy the output array to arr, so that arr now
  // contains sorted characters
  for(i=0;i<size;i++)
    arr[i]=out[i];
}

// Pigeonhole Sort -> run as pigeonholeSort(arr,size);
void pigeonHoleSort(int *arr, int n){
  // Find minimum and maximum values in arr
  int min = arr[0], max = arr[0];
  for (int i = 1; i < n; i++){
    if (arr[i] < min)
      min = arr[i];
    if (arr[i] > max)
      max = arr[i];
  }
  int range = max - min + 1; // Find range
  // Create an array of arrs. Size of array
  // range. Each arr represents a hole that
  // is going to contain matching elements.
  int holes[range]={0};
  // Traverse through input array and put every
  // element in its respective hole
  for (int i = 0; i < n; i++){
    holes[arr[i]-min]+=1;//.push_back(arr[i]);
  }
  // Traverse through all holes one by one. For
  // every hole, take its elements and put in
  // array.
  int index = 0;  // index in sorted array
  for (int count = 0; count < range; count++){
    while(holes[count]>0){
      holes[count] -=1;
      arr[index]=count+min;
      index++;
    }
  }
}

// Radix Sort
void radixSort(int *arr, int size){
  int expo, m;
  m = findMaxElement(arr, size);
  // Calling countingSort() for digit at (exp)th place in every input.
  for (expo = 1; m/expo > 0; expo *= 10)
    countingSort(arr, size, 32, 1, expo);
}

// Get maximum value from array.
int findMaxElement(int *arr, int size){
  int max = arr[0];
  for (int i = 1; i < size; i++)
    if (arr[i] > max)
      max = arr[i];
  return max;
}

// Get minimum value from array.
int findMinElement(int *arr, int size){
  int min = arr[0];
  for (int i = 1; i < size; i++)
    if (arr[i] < min)
      min = arr[i];
  return min;
}

// 5. Others
// Bitonic Sort -> run as bitonicSort(arr,0,N);
void bitonicSort(int *arr,int low, int high, int order=1){
  // order == 1 increase
  // order == 0 decrease
  if (high>1){
    int k = high/2;
    // sort in ascending order since order here is 1
    bitonicSort(arr, low, k, 1);
    // sort in descending order since order here is 0
    bitonicSort(arr, low+k, k, 0);
    // Will merge wole sequence in ascending order since order=1.
    bitonicMerge(arr,low, high, order);
  }
}

// compare and swap
void bitonicCompare(int *arr, int start, int distance, int order){
  for(int i = start; i < start+distance; i++)
    if (order==(arr[i]>arr[i+distance]))
      swap(arr[i],arr[i+distance]);
}

// The sequence to be sorted starts at index position low, size is the number of elements to be sorted
void bitonicMerge(int *arr, int low, int high, int order){
  if (high > 1){
    int k = high/2;
    bitonicCompare(arr, low, k, order);
    bitonicMerge(arr, low, k, order);
    bitonicMerge(arr, low+k, k, order);
  }
}

// Stooge Sort -> run as stoogeSort(arr,0,size-1);
void stoogeSort(int *arr, int low, int high){
  if (low >= high)
    return;
  else if (low < high){
    // If first element is smaller than last, swap them
    if (arr[low] > arr[high])
      swap(arr[low], arr[high]);
    // If there are more than 2 elements in the array
    if (high - low + 1 > 2) {
      int t = (high - low + 1) / 3;
      // Recursively sort first 2/3 elements
      stoogeSort(arr, low, high - t);
      // Recursively sort last 2/3 elements
      stoogeSort(arr, low + t, high);
      // Recursively sort first 2/3 elements again to confirm
      stoogeSort(arr, low, high - t);
    }
  }
}

// Pancake Sort -> run as pancakeSort(arr,size);
void pancakeSort(int *arr, int size){
  // Start from the complete array and one by one 
  // reduce current size by one
  for (int curr_size = size; curr_size > 1; --curr_size){
    // Find index of the maximum element in arr[0..curr_size-1]
    int mi = findMaxIndex(arr, curr_size);
    // Move the maximum element to end of 
    // current array if it's not already at the end
    if (mi != curr_size-1){
      // To move at the end, first move maximum number to beginning 
      flip(arr, mi);
      // Now move the maximum number to end by reversing current array
      flip(arr, curr_size-1);
    }
  }
}

// Reverses arr[0..i]
void flip(int *arr, int i){
  int start = 0;
  while (start < i){
    swap(arr[start], arr[i]);
    start++;
    i--;
  }
}
 
// Returns index of the maximum element in arr[0..size-1] 
int findMaxIndex(int *arr, int size){
  int mi, i;
  for (mi = 0, i = 0; i < size; ++i)
    if (arr[i] > arr[mi])
      mi = i;
  return mi;
}

// Bogo Sort -> run as bogoSort(arr,size);
void bogoSort(int *arr, int size){
  // if array is not sorted then shuffle the array again
  while(!isSorted(arr, size))
    shuffle(arr, size);
}

// verify is Sorted
bool isSorted(int *arr, int size){
  bool isOrdered = true;
  for(int i = 0; i < size-1; i++)
    if(arr[i] > arr[i + 1]){
        isOrdered = false;
        break;
    }
  return isOrdered;
}
 
// To generate permuatation of the array
void shuffle(int *arr, int size){
  for (int i=0; i < size; i++)
    swap(arr[i], arr[rand()%size]);
}

// swap positions
void swap(int * x, int * y){
  int temp = *x;
  *x = *y;
  *y = temp;
}
