package linkedlist

/*In a linked list of size n, where n is even, the ith node (0-indexed) of the linked list is known as the twin of the (n-1-i)th node, if 0 <= i <= (n / 2) - 1.

    For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only nodes with twins for n = 4.

The twin sum is defined as the sum of a node and its twin.

Given the head of a linked list with even length, return the maximum twin sum of the linked list.
*/
func pairSum(head *ListNode) int {
 valArr := []int{}
 len := 0
 cur := head
 for cur !=nil {
	valArr = append(valArr, cur.Val)
	len++
	cur = cur.Next
 }
 sum := -1
 for i:=0; i<len/2; i++ {
	curSum := valArr[i]	+ valArr[len-1-i]
	sum = max(sum, curSum)
 } 
 return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b 
}