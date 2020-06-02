package basic

import (
	"fmt"
	"testing"
)

/*
			0
	1    ---|---	2
3  -|-	4	  5    -|-	6

*/

func Test_Bitree(t *testing.T) {
	bt := &BiTree{}
	bt.Init(CfInt)
	root := NewBiTreeNode(0)
	n1 := NewBiTreeNode(1)
	n2 := NewBiTreeNode(2)
	n3 := NewBiTreeNode(3)
	n4 := NewBiTreeNode(4)
	n5 := NewBiTreeNode(5)
	n6 := NewBiTreeNode(6)
	var err error
	err = bt.InsertAsLeftChild(nil, root)
	checkErr(err, t)
	err = bt.InsertAsLeftChild(root, n1)
	checkErr(err, t)
	err = bt.InsertAsRightChild(root, n2)
	checkErr(err, t)
	err = bt.InsertAsLeftChild(n1, n3)
	checkErr(err, t)
	err = bt.InsertAsRightChild(n1, n4)
	checkErr(err, t)
	err = bt.InsertAsLeftChild(n2, n5)
	checkErr(err, t)
	err = bt.InsertAsRightChild(n2, n6)
	checkErr(err, t)

	t.Log("root", bt.Root().Data())                  //0
	t.Log("left", bt.Root().Left().Left().Data())    //3
	t.Log("right", bt.Root().Right().Right().Data()) //6
	t.Log("size", bt.Size())                         //7

	t.Log("删除左子树")
	bt.RemeveLeftChild(root)

	t.Log("root", bt.Root().Data())                  //0
	t.Log("left", bt.Root().Left())                  //nil
	t.Log("right", bt.Root().Right().Right().Data()) //6
	t.Log("size", bt.Size())                         //4

	//重新加回来
	err = bt.InsertAsLeftChild(root, n1)
	checkErr(err, t)
	err = bt.InsertAsLeftChild(n1, n3)
	checkErr(err, t)
	err = bt.InsertAsRightChild(n1, n4)
	checkErr(err, t)

	t.Log("root", bt.Root().Data())                  //0
	t.Log("left", bt.Root().Left().Left().Data())    //3
	t.Log("right", bt.Root().Right().Right().Data()) //6
	t.Log("size", bt.Size())                         //7

	fmt.Println("前序")
	//前序
	l1 := new(List)
	bt.Root().Preorder(l1)
	l1.Traverse(PrintNode)

	fmt.Println("前序非递归")
	//前序
	l11 := new(List)
	bt.Root().PreorderNo(l11)
	l11.Traverse(PrintNode)

	fmt.Println("中序")
	//中序
	l2 := new(List)
	bt.Root().Inorder(l2)
	l2.Traverse(PrintNode)

	fmt.Println("中序非递归")
	//中序
	l22 := new(List)
	bt.Root().InorderNo(l22)
	l22.Traverse(PrintNode)

	fmt.Println("后序")
	//后序
	l3 := new(List)
	bt.Root().Postorder(l3)
	l3.Traverse(PrintNode)

	fmt.Println("后序非递归")
	//后序
	l33 := new(List)
	bt.Root().PostorderNo(l33)
	l33.Traverse(PrintNode)

	fmt.Println("层次")
	//层次
	l4 := new(List)
	bt.Root().LevelorderNo(l4)
	l4.Traverse(PrintNode)

}

func checkErr(err error, t *testing.T) {
	if err != nil {
		panic(err)
	}
}
