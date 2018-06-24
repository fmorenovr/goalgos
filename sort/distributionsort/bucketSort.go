package distributionsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgs/sort/efficientsort";
  "github.com/jenazads/goalgs/sort/sortfunctions";
)

// Bucket Sort
func BucketSort(arr []interface{}, comp goutils.TypeComparator, op goutils.TypeOperator, low, high int) () {
  bucketSize:=4
  maxValueIndex, minValueIndex := sortfunctions.FindMaxMinElementIndex(arr, comp, low, high)
  minValue, maxValue := arr[minValueIndex], arr[maxValueIndex]
  numOfBuckets := goutils.ToInt(op(op(op(maxValue, minValue, "-"),bucketSize, "/"), 1, "+"));
  buckets := make([][]interface{}, numOfBuckets)

  for i := low; i < high; i++ {
    curr_index := goutils.ToInt(op(op(arr[i], minValue, "-"), bucketSize, "/"))
    buckets[curr_index] = append(buckets[curr_index],arr[i]);
  }
  index := low;
  for i := 0; i < numOfBuckets; i++ {
    bucketSize := len(buckets[i])
    if bucketSize > 0 {
      efficientsort.QuickSort(buckets[i], comp, 0, bucketSize-1)
    }
    for j := 0; j < bucketSize; j++ {
      arr[index] = buckets[i][j];
      index++;
    }
  }
}
