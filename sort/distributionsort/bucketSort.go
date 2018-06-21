package distributionsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/golgos/sort/simplesort";
)

// Bucket Sort
func BucketSort(arr []interface{}, comp goutils.TypeComparator, op goutils.TypeOperator, low, high int) () {
  bucketSize:=4
  maxValueIndex, minValueIndex := FindMaxMinElementIndex(arr, comp, low, high)
  minValue, maxValue := arr[minValueIndex], arr[maxValueIndex]
  numOfBuckets := goutils.ToInt(op(op(op(maxValue, minValue, "-"),bucketSize, "/"), 1, "+"));
  buckets := make([][]interface{}, numOfBuckets)

  for i := low; i < high; i++ {
    index_curr := goutils.ToInt(op(op(arr[i], minValue, "-"), bucketSize, "/"))
    buckets[index_curr] = append(buckets[index_curr],arr[i]);
  }
  k := low;
  for i := 0; i < numOfBuckets; i++ {
    bucketSize := len(buckets[i])
    simplesort.InsertionSort(buckets[i], comp, 0, bucketSize)
    for j := 0; j < bucketSize; j++ {
      arr[k] = buckets[i][j];
      k++;
    }
  }
}
