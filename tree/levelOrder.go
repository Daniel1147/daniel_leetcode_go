package main

import "fmt"
import . "github.com/Daniel1147/tree"

var ansList [][]int

func levelOrder(root *TreeNode) [][]int {
    ansList = make([][]int, 0)
    addNode(root, 0)

    return ansList
}

func addNode(node *TreeNode, level int) {
    if node == nil {
        return
    }

    if len(ansList) <= level {
        ansList = append(ansList, []int{})
    }

    addNode(node.Left, level + 1)
    addNode(node.Right, level + 1)

    ansList[level] = append(ansList[level], node.Val)

    return
}

func main() {
    treeNode := Sample()
    treeNode.Show()
    fmt.Println(levelOrder(treeNode))

	return
}

