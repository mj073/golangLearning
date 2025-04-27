package binarytree

import (
	"reflect"
	"testing"
)

var (
//	                 a
//	                / \
//	               b   c
//	              / \   \
//	             d   e   f
	tree = Tree{
		RootNode: &Node {
			Data: "a",
			Left: &Node {
				Data: "b",
				Left: &Node{
					Data: "d",
				},
				Right: &Node{
					Data: "e",
				},
			},
			Right: &Node{
				Data: "c",
				Right: &Node{
					Data: "f",
				},
			},
		},
	}
)
func TestTree_Dft(t *testing.T) {
	if r := tree.Dft(); !reflect.DeepEqual(r, []interface{}{"a","b","d","e","c","f"}) {
		t.Fatalf("Failed")
	}
}

func TestTree_DftRecursive(t *testing.T) {
	if r := tree.DftRecursive(tree.RootNode); !reflect.DeepEqual(r, []interface{}{"a","b","d","e","c","f"}) {
		t.Fatalf("Failed, %v", r)
	}
	//tree.DftRecursive(tree.RootNode)
}

func TestTree_Bft(t *testing.T) {
	if r := tree.Bft(); !reflect.DeepEqual(r, []interface{}{"a","b","c","d","e","f"}) {
		t.Fatalf("Failed, %v", r)
	}
}

func TestTree_Bft2(t *testing.T) {
	if r := tree.Bft2(); !reflect.DeepEqual(r, []interface{}{"a","b","c","d","e","f"}) {
		t.Fatalf("Failed, %v", r)
	}
	/*a := []int{0}
	_ = a[0]
	a = a[1:] // Does not Panic
	t.Log(len(a))*/
}