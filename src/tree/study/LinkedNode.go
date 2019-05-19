package study

import (
	"container/list"
	"fmt"
	"math"
)

//MyStack used for Pre/In/Post order traverse
type MyStack struct {
	list *list.List
}

//Push ...
func (stack *MyStack) Push(elem interface{}) {
	stack.list.PushBack(elem)
}

//Pop from back
func (stack *MyStack) Pop() interface{} {
	if elem := stack.list.Back(); elem != nil {
		stack.list.Remove(elem)
		return elem.Value
	}
	return nil
}

//Len ...
func (stack *MyStack) Len() int {
	return stack.list.Len()
}

//Empty ...
func (stack *MyStack) Empty() bool {
	return stack.list.Len() == 0
}

//MyQueue used for level traverse
type MyQueue struct {
	list *list.List
}

//Push ...
func (queue *MyQueue) Push(elem interface{}) {
	queue.list.PushBack(elem)
}

//Pop from front
func (queue *MyQueue) Pop() interface{} {
	if elem := queue.list.Front(); elem != nil {
		queue.list.Remove(elem)
		return elem.Value
	}
	return nil
}

//Node ...
type Node struct {
	data        interface{}
	left, right *Node
}

//NewNode ...
func NewNode(data interface{}) *Node {
	return &Node{data: data}
}

//String ...
func (node *Node) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", node.data, node.left, node.right)
}

//SetLeftNode ...
func (node *Node) SetLeftNode(l *Node) {
	if node == nil {
		return
	}
	node.left = l
}

//SetRightNode ...
func (node *Node) SetRightNode(r *Node) {
	if node == nil {
		return
	}
	node.right = r
}

//GetLeftNode ...
func (node *Node) GetLeftNode() *Node {
	if node == nil {
		return nil
	}
	return node.left
}

//GetRightNode ...
func (node *Node) GetRightNode() *Node {
	if node == nil {
		return nil
	}
	return node.right
}

//PreOrderTraverse 前序遍历
func (node *Node) PreOrderTraverse() {
	if node == nil {
		return
	}
	fmt.Print(node.data, " ")
	node.left.PreOrderTraverse()
	node.right.PreOrderTraverse()
}

//PreOrderTraverseByStack 前序遍历
//前序遍历的思路是通过栈，将右子树先行压栈，然后左子树压栈
func (node *Node) PreOrderTraverseByStack() {
	if node == nil {
		return
	}

	stack := MyStack{list: list.New()}
	stack.Push(node)

	for !stack.Empty() {
		nd := stack.Pop().(*Node)
		fmt.Print(nd.data, " ")

		if nd.right != nil {
			stack.Push(nd.right)
		}

		if nd.left != nil {
			stack.Push(nd.left)
		}
	}

}

//InOrderTraverse 中序遍历
func (node *Node) InOrderTraverse() {
	if node == nil {
		return
	}
	node.left.InOrderTraverse()
	fmt.Print(node.data, " ")
	node.right.InOrderTraverse()
}

//InOrderTraverseByStack 中序遍历
// 先将左子树压栈，直到最左边得左子树为空，然后左子树出栈，打印，返回上一层父节点，然后右子树
func (node *Node) InOrderTraverseByStack() {
	if node == nil {
		return
	}

	stack := MyStack{list: list.New()}
	current := node
	for !stack.Empty() || current != nil {
		if current != nil {
			stack.Push(current)
			current = current.left
		} else {
			current = stack.Pop().(*Node)
			fmt.Print(current.data, " ")
			current = current.right
		}
	}
}

//PostOrderTraverse 后序遍历
func (node *Node) PostOrderTraverse() {
	if node == nil {
		return
	}
	node.left.PostOrderTraverse()
	node.right.PostOrderTraverse()
	fmt.Print(node.data, " ")
}

//PostOrderTraverseByStack 后序遍历
//遍历顺序：左-右-根
//stack1 压出栈顺序： 根压->根出->根左子入->根右子入->根右子出->根左子出
//stack2 压栈顺序： 根->右->左， 出栈顺序： 左-右-根
func (node *Node) PostOrderTraverseByStack() {
	if node == nil {
		return
	}

	stack1, stack2 := MyStack{list: list.New()}, MyStack{list: list.New()}
	stack1.Push(node)

	for !stack1.Empty() {
		nd := stack1.Pop().(*Node)
		stack2.Push(nd)

		if nd.left != nil {
			stack1.Push(nd.left)
		}

		if nd.right != nil {
			stack1.Push(nd.right)
		}
	}

	for !stack2.Empty() {
		n := stack2.Pop().(*Node)
		fmt.Print(n.data, " ")
	}
}

//LevelTraverse 层序遍历
//借助队列实现，入对顺序：根->左子->右子  出队顺序： 根->左子->右子
func (node *Node) LevelTraverse() {
	if node == nil {
		return
	}
	var nlast *Node

	last := node
	level := 1

	queue := MyQueue{list: list.New()}
	queue.Push(node)

	fmt.Println(fmt.Sprintf("-----this is %d level-----", level))
	for queue.list.Len() > 0 {
		nd := queue.Pop().(*Node)

		if nd.left != nil {
			queue.Push(nd.left)
			nlast = nd.left
		}

		if nd.right != nil {
			queue.Push(nd.right)
			nlast = nd.right
		}
		fmt.Print(nd.data, " ")

		if last == nd {
			last = nlast
			level++
			fmt.Println()
			fmt.Println(fmt.Sprintf("-----this is %d level-----", level))
		}
	}
}

//Depth get tree depth
func (node *Node) Depth() int {
	if node == nil {
		return 0
	}
	left := node.left.Depth() + 1
	right := node.right.Depth() + 1

	return int(math.Max(float64(left), float64(right)))
}

//DepthByLevelTraverse ...
func (node *Node) DepthByLevelTraverse() int {
	if node == nil {
		return 0
	} else if node.left == nil && node.right == nil {
		return 1
	} else {
		depth := node.DepthByLevelT()
		return depth
	}
}

//DepthByLevelT ...
func (node *Node) DepthByLevelT() int {
	queue := list.New()
	queue.PushBack(node)

	maxDepth := 0
	for {
		len := queue.Len()
		if len == 0 {
			break
		}
		for i := 0; i < len; i++ { //遍历每层结点，然后maxDepth+1
			front := queue.Front()
			nd := (front.Value).(*Node)
			queue.Remove(front)

			if nd.left != nil {
				queue.PushBack(nd.left)
			}

			if nd.right != nil {
				queue.PushBack(nd.right)
			}
		}
		maxDepth++
	}
	return maxDepth
}
