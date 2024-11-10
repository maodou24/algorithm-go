package list

import (
	"github.com/wgzhulin/algo-go/internel/list"
)

func reverse(head *list.Node) *list.Node {
	prev, next := head, head.Next

	for next != nil {
		temp := next.Next

		next.Next = prev
		prev = next
		next = temp
	}

	head.Next = nil
	return prev
}
