
package main

import "fmt"

type Node struct {
	Left *Node
	Right *Node
	Data interface{}
}
var tree Node
var level = 0
var avg []float32

func main(){
	data := []int{3,9,20,15,7}

	for _, i := range data {
		//tree.setBinarySearchTree(i)
		tree.setBinaryTree(i)
	}
	tree.inorder()
	fmt.Println()
	tree.preorder()
	fmt.Println()
	tree.postorder()
	fmt.Println()
	tree.bfs()
	fmt.Println(avg)
}

func (n *Node) setBinaryTree(x int){
	if n.Data == nil {
		n.Data = x
	}else if (n.Left != nil){
		n.Left.setBinaryTree(x)
	}else if (n.Left == nil){
		n.Left = &Node{}
		n.Left.Data = x
	}else if (n.Right != nil){
		n.Right.setBinaryTree(x)
	}else if (n.Right == nil){
		n.Right = &Node{}
		n.Right.Data = x
	}
}
func (n *Node) setBinarySearchTree(x int){
	if n.Data == nil {
		n.Data = x
	}else if (n.Left != nil) && (x < n.Data.(int)){
		n.Left.setBinarySearchTree(x)
	}else if (n.Left == nil) && (x < n.Data.(int)){
		n.Left = &Node{}
		n.Left.Data = x
	}else if (n.Right != nil) &&(x > n.Data.(int)){
		n.Right.setBinarySearchTree(x)
	}else if (n.Right == nil) && (x > n.Data.(int)){
		n.Right = &Node{}
		n.Right.Data = x
	}

}
// LEFT -> ROOT -> RIGHT
func (n *Node) inorder(){
	if n == nil{
		fmt.Println("No elements to display..")
		return
	}
	if n.Left != nil{
		n.Left.inorder()
	}
	fmt.Printf("%d->",n.Data)
	if n.Right != nil{
		n.Right.inorder()
	}
}
// ROOT -> LEFT -> RIGHT
func (n *Node) preorder(){
	if n == nil{
		fmt.Println("No elements to display..")
		return
	}
	fmt.Printf("%d->",n.Data)
	if n.Left != nil{
		n.Left.preorder()
	}
	if n.Right != nil{
		n.Right.preorder()
	}
}
// LEFT -> RIGHT -> ROOT
func (n *Node) postorder(){
	if n == nil{
		fmt.Println("No elements to display..")
		return
	}
	if n.Left != nil{
		n.Left.postorder()
	}
	if n.Right != nil{
		n.Right.postorder()
	}
	fmt.Printf("%d->",n.Data)
}

func (n *Node) bfs(){
	var temp Node
	temp.Left = n
	sum := 0
	nodes := 0
	if n == nil{
		fmt.Println("No elements to display..")
		return
	}
	if level == 0{
		fmt.Printf("%d->",n.Data)
		avg = append(avg,float32(n.Data.(int)))
		level++
	}
	if n.Left != nil{
		fmt.Printf("%d->",n.Left.Data)
		sum = sum + n.Left.Data.(int)
		nodes++
	}
	if n.Right != nil{
		fmt.Printf("%d->",n.Right.Data)
		sum = sum + n.Right.Data.(int)
		nodes++
	}
	avg = append(avg,float32(float32(sum) / float32(nodes)))
	level++
	if n.Left != nil{
		n.Left.bfs()
	}
	if n.Right != nil{
		n.Right.bfs()
	}
}