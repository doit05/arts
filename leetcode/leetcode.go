package main

import (
	"fmt"
)

func main() {
	node_1 := &Node{Val: 1}
	node_2 := &Node{Val: 2}
	node_3 := &Node{Val: 3}
	node_4 := &Node{Val: 4}
	node_5 := &Node{Val: 5}
	node_6 := &Node{Val: 6}
	node_7 := &Node{Val: 7}
	node_1.Left = node_2
	node_1.Right = node_3
	node_2.Left = node_4
	node_2.Right = node_5
	node_3.Left = node_6
	node_3.Right = node_7
	root := node_1
	nodeMap := make(map[int][]*Node, 1)
	travel(root, nodeMap, 0)
	for level, nodes := range nodeMap {
		fmt.Println(level, nodes)

	}
}

//Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func travel(root *Node, nodeMap map[int][]*Node, high int) *Node {
	if root == nil {
		return nil
	}

	level(root, nodeMap, high)
	travel(root.Left, nodeMap, high+1)
	travel(root.Right, nodeMap, high+1)
	return root
}

func level(node *Node, nodeMap map[int][]*Node, high int) {
	if node == nil {
		return
	}
	if _, ok := nodeMap[high]; !ok {
		nodeMap[high] = make([]*Node, 0, 1)
	}
	nodeMap[high] = append(nodeMap[high], node)
}

func connect(root *Node) *Node {
	list := levelOrder(root)
	rightList := getRightList(root)
	length := len(list)
	for i, j := 0, 1; j < length; {
		if _, ok := rightList[list[i]]; ok {
			list[i].Next = nil
		} else {
			list[i].Next = list[j]
		}
		i++
		j++
	}
	return root
}

func getRightList(root *Node) map[*Node]*Node {
	var list = make(map[*Node]*Node)
	node := root
	for node != nil {
		list[node] = node
		node = node.Right
	}
	return list
}

type Queue struct {
	front int
	rear  int
	list  []*Node
}

func NewQueue() *Queue {
	q := new(Queue)
	q.front = 0
	q.rear = 0
	q.list = make([]*Node, 0)
	return q
}

func (q *Queue) Push(i *Node) {
	q.list = append(q.list, i)
	q.rear++
}

func (q *Queue) Pop() *Node {
	if q.front != q.rear {
		res := q.list[q.front]
		q.front++
		return res
	}
	return nil
}

func levelOrder(root *Node) []*Node {
	list := make([]*Node, 0, 1)
	q := NewQueue()
	p := root
	q.Push(p)

	for q.front != q.rear {
		for n := q.rear - q.front; n > 0; n-- {
			k := q.Pop()
			if k != nil {
				list = append(list, k)
				q.Push(k.Left)
				q.Push(k.Right)
			}

		}
	}
	return list
}
