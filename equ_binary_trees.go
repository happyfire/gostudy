// 这是go tour中等价二叉树的练习，需要"golang.org/x/tour/tree"才能编译
// Walk方法中序遍历（左根右）二叉树，填入到一个channel中
// Same方法比较两个channel

package main

//
//import "golang.org/x/tour/tree"
//import "fmt"
//
//// Walk walks the tree t sending all values
//// from the tree to the channel ch.
//func Walk(t *tree.Tree, ch chan int) {
//  if t == nil {
//		return
//	}
//
//	if t.Left != nil {
//		Walk(t.Left, ch)
//	}
//
//	ch <- t.Value
//
//	if t.Right != nil {
//		Walk(t.Right, ch)
//	}
//}
//
//// Same determines whether the trees
//// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	go Walk(t1, ch1)
//	go Walk(t2, ch2)
//	for i := 0; i < 10; i++ {
//		v1 := <-ch1
//		v2 := <-ch2
//		//fmt.Println(v1,v2)
//		if v1 != v2 {
//			return false
//		}
//	}
//	return true
//}
//
//func main() {
//	t1 := tree.New(1)
//	t2 := tree.New(1)
//
//	fmt.Println(t1)
//	fmt.Println(t2)
//
//	fmt.Println(Same(t1, t2))
//
//}
