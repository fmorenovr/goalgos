package efficientsort

import(
  "github.com/jenazads/goutils";
  "github.com/jenazads/gods/trees/bstree";
)

// Tree Sort
func TreeSort(arr []interface{}, comp goutils.TypeComparator, op goutils.TypeOperator, toInt func(interface{})(interface{}), low, high int) () {
  if(low<high){
    tree := bstree.NewBSTree(comp, op)
    tree.Insert(toInt(arr[low]), arr[low])
    for i:=low+1; i<high; i++ {
      tree.Insert(toInt(arr[i]), arr[i])
    }
    keys:=tree.Values()
    j:=0
    for i:=low; i<high; i++ {
      arr[i] = keys[j]
      j++
    }
  }
}
