package distributionsort

import(
  "math";
  "github.com/jenazads/goutils";
  "github.com/jenazads/goalgos/sort/simplesort";
  "github.com/jenazads/goalgos/sort/sortfunctions";
)

// Bucket Sort
func BucketSort(arr []interface{}, comp goutils.TypeComparator, op goutils.TypeOperator, low, high int) () {
  if(low<high){
    bucketSize:=sortfunctions.HighestPowerofTwoLessThan(int(math.Sqrt(float64(high-low))))
    minValue, maxValue := sortfunctions.FindMaxMinElementIndex(arr, comp, low, high)
    numOfBuckets := goutils.ToInt(op(op(op(maxValue, minValue, "-"),bucketSize, "/"), 1, "+"));
    buckets := make([][]interface{}, numOfBuckets)

    for i:=low; i<high; i++ {
      curr_index := goutils.ToInt(op(op(arr[i], minValue, "-"), bucketSize, "/"))
      buckets[curr_index] = append(buckets[curr_index],arr[i]);
    }
    index := low;
    for i := 0; i < numOfBuckets; i++ {
      bucketLength := len(buckets[i])
      if bucketLength > 0 {
        simplesort.InsertionSort(buckets[i], comp, 0, bucketLength)
      }
      for j := 0; j < bucketLength; j++ {
        arr[index] = buckets[i][j];
        index++;
      }
    }
  }
}
