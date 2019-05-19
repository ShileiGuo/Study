package study

import "fmt"

//BSTree named binary search tree
type BSTree struct {
	*Node
	comepareFunc func(v, nodeV interface{}) int
}

//NewBSTree ...
func NewBSTree(rootV interface{}, comepareFunc func(v, nodeV interface{}) int) *BSTree {
	return &BSTree{Node: NewNode(rootV), comepareFunc: comepareFunc}
}

var CompareFunc = func(v, nodeV interface{}) int {
	vInt := v.(int)
	nodeVInt := nodeV.(int)

	if vInt > nodeVInt {
		return 1
	} else if vInt < nodeVInt {
		return -1
	}
	return 0
}

//Insert ...
func (bst *BSTree) Insert(v interface{}) bool {
	p := bst.Node

	for p != nil {
		comepareResult := bst.comepareFunc(v, p.data)
		if comepareResult == 0 {
			return false
		} else if comepareResult > 0 {
			if nil == p.right {
				p.right = NewNode(v)
				return true
			}
			p = p.right
		} else if comepareResult < 0 {
			if nil == p.left {
				p.left = NewNode(v)
				return true
			}
			p = p.left
		}
	}
	return false
}

//Find ...
func (bst *BSTree) Find(v interface{}) *Node {
	p := bst.Node

	for p != nil {
		comepareResult := bst.comepareFunc(v, p.data)
		if comepareResult == 0 {
			return p
		} else if comepareResult > 0 {
			p = p.right
		} else if comepareResult < 0 {
			p = p.left
		}
	}
	return nil
}

//Delete ...
func (bst *BSTree) Delete(v interface{}) bool {
	p := bst.Node
	var pp *Node //p的父节点
	for p != nil && bst.comepareFunc(v, p.data) != 0 {
		pp = p
		if bst.comepareFunc(v, p.data) > 0 {
			p = p.right
		} else {
			p = p.left
		}
	}
	if p == nil { //没有发现待删除的node
		fmt.Println("Not found")
		return false
	}

	//待删除的节点，有左右子树，此时应该把右子树中最小的节点替换p，把最小节点从右子树中删除，
	//最小节点可能没有左孩子，但是有右孩子，所以要结合最下面看
	if p.left != nil && p.right != nil {
		minP := p.right
		minPP := p //minPP表示minP的父节点
		for minP.left != nil {
			minPP = minP
			minP = minP.left
		}

		p.data = minP.data //将 minP 的数据替换到 p 中
		p = minP           //删除p节点变为删除minP节点操作，这一点不好理解，要结合后面看
		pp = minPP
	}

	//删除节点是叶子节点或者仅有一个子节点
	var child *Node //p的孩子
	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}

	if pp == nil { //删除的是根节点，
		bst.Node = child
	} else if pp.left == p {
		pp.left = child
	} else if pp.right == p {
		pp.right = child
	}
	return true
}

//MinNode ...
func (bst *BSTree) MinNode() *Node {
	if bst == nil {
		return nil
	}
	node := bst.Node
	for node.left != nil {
		node = node.left
	}
	fmt.Println("min node is ", node.data)
	return node
}

//MaxNode ...
func (bst *BSTree) MaxNode() *Node {
	if bst == nil {
		return nil
	}
	node := bst.Node
	for node.right != nil {
		node = node.right
	}
	fmt.Println("max node is ", node.data)
	return node
}
