//Structures in GO program
package main

import "fmt"

type Node struct {
	data int
	name string
}

func main() {
	var node1 Node
	var node2 Node
	node1.data = 2
	node1.name = "Abc"
	node2.data = 3
	node2.name = "Def"

	fmt.Printf("\nNode 1 data %d & name is %s \n", node1.data, node1.name)
	fmt.Printf("Node 2 data %d & name is %s \n", node2.data, node2.name)
	node3 := struct_func(node1)
	fmt.Printf("\n Changed Node 1 data %d & name is %s \n", node3.data, node3.name)

	//Pointer to a structure
	var struct_ptr *Node
	struct_ptr = &node2
	fmt.Printf("\nThe value of the structure ponter is data is %d and name is %s \n ", struct_ptr.data, struct_ptr.name)

}

func struct_func(node Node) Node {
	node.data = node.data + 5
	node.name = "Dirty bit"
	return node
}
