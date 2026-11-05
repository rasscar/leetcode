/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
    if head == nil || head.Next == nil {
        return []int{-1, -1}
    }

    prev := head
    curr := head.Next
    pos := 2

    first := -1
    last := -1
    minDist := int(^uint(0) >> 1) 

    for curr != nil && curr.Next != nil {
        next := curr.Next

        isLocalMax := prev.Val < curr.Val && curr.Val > next.Val
        isLocalMin := prev.Val > curr.Val && curr.Val < next.Val

        if isLocalMax || isLocalMin {
            if first == -1 {
                first = pos
            } else {
                if pos-last < minDist {
                    minDist = pos - last
                }
            }
            last = pos
        }

        prev = curr
        curr = next
        pos++
    }

    if first == -1 || first == last {
        return []int{-1, -1}
    }

    return []int{minDist, last - first}
}
