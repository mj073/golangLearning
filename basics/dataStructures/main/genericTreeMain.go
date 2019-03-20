package main

import "basics/dataStructures/genericTree"

func main(){
	tree := &genericTree.Tree{}
	tree.Insert(0,genericTree.Person{1,"Mahesh",26,genericTree.Male,false})
	tree.DisplayTree()
}
