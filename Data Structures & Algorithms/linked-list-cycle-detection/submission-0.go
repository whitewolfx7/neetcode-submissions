/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {

    var slow *ListNode = head
	var fast *ListNode = head

	if head == nil || head.Next == nil{
		return false
	}

	for fast != nil && fast.Next != nil{

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast{
			return true
		}
	}

	return false
}
