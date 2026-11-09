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

### 209.长度最小的子数组
```
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：

输入：s = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
提示：

1 <= target <= 10^9
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^5
```

```go
func minSubArrayLen(target int, nums []int) int {

    // 滑动窗口
    result := len(nums)+1
    i := 0
    var max int
    for j := 0; j <= len(nums)-1; j++ {
        max = max + nums[j]
        for max >= target {
            if j-i+1 < result {
                result = j-i+1
            }
            max = max - nums[i]
            i++
        }
    }
    if result == len(nums)+1{
        return 0
    }
    return result
}
```

### 59.螺旋矩阵II
```
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
```

```go
func generateMatrix(n int) [][]int {
    
    startX, startY := 0, 0
    i, j := 0, 0
    offset := 1 
    num := 1
    numbers := make([][]int, n)
    for i := range numbers {
        numbers[i] = make([]int, n)
    }
    for l := 0; l < n/2; l++ {
        for j = startY; j < n-offset; j++ {
            numbers[startX][j] = num
            num++
        }
        for i = startX; i < n-offset; i++ {
            numbers[i][j] = num
            num++
        }
        for ; j > offset-1; j-- {
            numbers[i][j] = num
            num++
        }
        for ; i > offset-1; i-- {
            numbers[i][j] = num
            num++
        }
        offset++
        startX++
        startY++
    }
    if n % 2 != 0 {
        numbers[n/2][n/2] = n * n
    }
    return numbers
}
```

## 链表
### 203.移除链表元素
```
题意：删除链表中等于给定值 val 的所有节点。

示例 1： 输入：head = [1,2,6,3,4,5,6], val = 6 输出：[1,2,3,4,5]

示例 2： 输入：head = [], val = 1 输出：[]

示例 3： 输入：head = [7,7,7,7], val = 7 输出：[]
```

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    
    for head != nil && head.Val == val {
        head = head.Next
    }

    node := head
    for node != nil && node.Next != nil {
        if node.Next.Val == val {
            node.Next = node.Next.Next
        } else {
            node = node.Next
        }
    }

    return head
}
```

### 707.设计链表
```
题意：

在链表类中实现这些功能：
get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
```

```go
type SingleNode struct {
    Val  int
    Next *SingleNode
}

type MyLinkedList struct {
    dummyHead *SingleNode
    Size      int
}


func Constructor() MyLinkedList {
    node := &SingleNode{
        Val : -999, 
        Next: nil,
    }

    return MyLinkedList{
        dummyHead: node,
        Size     : 0,
    }
}


func (this *MyLinkedList) Get(index int) int {
    if this == nil || index < 0 || index >= this.Size {
        return -1
    }

    cur := this.dummyHead.Next
    for i := 0; i < index; i++ {
        cur = cur.Next
    }

    return cur.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {

    newNode := &SingleNode{Val: val}
    newNode.Next = this.dummyHead.Next
    this.dummyHead.Next = newNode
    this.Size++
}


func (this *MyLinkedList) AddAtTail(val int)  {
    newNode := &SingleNode{Val: val}
    cur := this.dummyHead
    for cur.Next != nil {
        cur = cur.Next
    }

    cur.Next = newNode
    this.Size++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 {
        index = 0
    } else if index > this.Size {
        return
    }
    newNode := &SingleNode{Val: val}
    cur := this.dummyHead
    for i := 0; i < index; i++ {
        cur = cur.Next
    }

    newNode.Next = cur.Next
    cur.Next = newNode
    this.Size++
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.Size {
        return
    }
    cur := this.dummyHead
    for i := 0; i < index; i++ {
        cur = cur.Next
    }
    if cur.Next != nil {
        cur.Next = cur.Next.Next
    }
    this.Size--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
```

```
题意：反转一个单链表。

示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL
```
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    // 使用双指针法
    cur := head
    var pre *ListNode
    var tmp *ListNode
    for cur != nil {
        tmp = cur.Next
        cur.Next = pre
        pre = cur
        cur = tmp
    }

    return pre
}
```

### 19.删除链表的倒数第N个节点
```
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
输入：head = [1,2,3,4,5], n = 2 输出：[1,2,3,5]

示例 2：

输入：head = [1], n = 1 输出：[]

示例 3：

输入：head = [1,2], n = 1 输出：[1]
```
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
    // 虚拟头节点
    dummyHead := ListNode{
        Val: 999,
        Next: head,
    }

    fast, slow := &dummyHead, &dummyHead
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }

    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }

    slow.Next = slow.Next.Next
    return dummyHead.Next
}
```