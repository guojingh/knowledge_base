# leetcode 

## 数组

### 704.二分法查找
```
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果 target 存在返回下标，否则返回 -1。

你必须编写一个具有 O(log n) 时间复杂度的算法。


示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
 

提示：

你可以假设 nums 中的所有元素是不重复的。
n 将在 [1, 10000]之间。
nums 的每个元素都将在 [-9999, 9999]之间。
```

```go
func search(nums []int, target int) int {

    // 左闭右闭
    left, right := 0, len(nums)-1
    for {
        middle := (left + right) / 2
        if left <= right {
            if nums[middle] > target {
                right = middle - 1
            } else if nums[middle] < target {
                left = middle + 1 
            } else {
                return middle
            }
        } else {
            break
        }
    }
    return -1
}
```

### 27.移除元素
```
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1: 给定 nums = [3,2,2,3], val = 3, 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。 你不需要考虑数组中超出新长度后面的元素。

示例 2: 给定 nums = [0,1,2,2,3,0,4,2], val = 2, 函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
```

```go
func removeElement(nums []int, val int) int {
    // 快慢指针进行解决，时间复杂度只有O(n)
    fast, slow := 0, 0
    for ; fast < len(nums); fast++ {
        if nums[fast] != val {
            nums[slow] = nums[fast]
            slow++
        }
    }  
    return slow
} 
```

### 977.有序数组的平方
```
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]
示例 2：

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
```

```go
func sortedSquares(nums []int) []int {

    // 使用双指针，时间复杂度只有O(n)
    k := len(nums) - 1
    res := make([]int, len(nums))
    for i,j := 0,len(nums)-1; i <= j; {
        if nums[i] * nums[i] < nums[j] * nums[j] {
            res[k] = nums[j] * nums[j]
            j--
        } else {
            res[k] = nums[i] * nums[i]
            i++
        }
        k--
    }

    return res
}
```