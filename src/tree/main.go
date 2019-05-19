package main

import (
	"fmt"
	"tree/study"
)

func main() {

	//ArrayNode Operate
	fmt.Println("\nTreeNode base Array")
	stack := study.NewArrayStack()
	isempty := stack.IsEmpty()

	fmt.Println("stack is empty:", isempty)
	stack.Push("0")
	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	stack.Push("4")
	stack.Push("5")
	stack.Push("6")
	stack.Print()

	fmt.Println("PreOrderTraverseByArray")
	stack.PreOrderTraverseByArray(0)
	fmt.Println("\nInOrderTraverseByArray")
	stack.InOrderTraverseByArray(0)
	fmt.Println("\nPostOrderTraverseByArray")
	stack.PostOrderTraverseByArray(0)

	//TreeNode Traverse
	fmt.Println("\n------------TreeNode base Link---------------")

	lNode := study.NewNode("1")
	rNode := study.NewNode("2")

	llNode := study.NewNode("3")
	lrNode := study.NewNode("4")

	rlNode := study.NewNode("5")
	rrNode := study.NewNode("6")

	treeNode := study.NewNode("root")
	treeNode.SetLeftNode(lNode)
	treeNode.SetRightNode(rNode)

	treeNode.GetLeftNode().SetLeftNode(llNode)
	treeNode.GetLeftNode().SetRightNode(lrNode)

	treeNode.GetRightNode().SetLeftNode(rlNode)
	treeNode.GetRightNode().SetRightNode(rrNode)

	fmt.Println("PreOrderTraverse")
	treeNode.PreOrderTraverse()
	fmt.Println("\nPreOrderTraverseByStack")
	treeNode.PreOrderTraverseByStack()

	fmt.Println("\nInOrderTraverse")
	treeNode.InOrderTraverse()
	fmt.Println("\nInOrderTraverseByStack")
	treeNode.InOrderTraverseByStack()

	fmt.Println("\nPostOrderTraverse")
	treeNode.PostOrderTraverse()
	fmt.Println("\nPostOrderTraverseByStack")
	treeNode.PostOrderTraverseByStack()

	fmt.Println("\nLevelTraverse")
	treeNode.LevelTraverse()

	//BST
	fmt.Println("\n------------BST base Link---------------")

	bst := study.NewBSTree(4, study.CompareFunc)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)
	bst.Insert(4)
	bst.Insert(5)
	bst.Insert(6)

	bst.MaxNode()
	bst.MinNode()

	fmt.Println("\nLevelTraverse")
	bst.Node.LevelTraverse()

	fmt.Println("\nfind 3")
	node := bst.Find(3)
	fmt.Println(node)

	fmt.Println("Delete 1")
	bst.Delete(1)
	node = bst.Find(1)
	fmt.Println(node)
	fmt.Println("\nLevelTraverse")
	bst.Node.LevelTraverse()

	fmt.Println("Delete 4")
	bst.Delete(4)
	node = bst.Find(4)
	fmt.Println(node)
	fmt.Println("LevelTraverse")
	bst.Node.LevelTraverse()

	depth := bst.Node.Depth()
	fmt.Println("Tree depth is ", depth)

	depthBT := bst.Node.DepthByLevelTraverse()
	fmt.Println("Tree depth by level traverse is ", depthBT)
}
