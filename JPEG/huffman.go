package JPEG

import (
	"sort"
)

type node struct {
	symbol int
	val int
	code string
	left *node
	right *node
}

type nodeArray []node

var codeTable map[int]string

func huffman(m map[int]int)(map[int]string){

	codeTable = make(map[int]string)
	var array nodeArray
	for k, v := range m {
		array = append(array, node{k, v,"", nil, nil})
	}
	sort.Sort(array)
	for ;len(array) != 1 ;  {
		left := &array[0]
		right := &array[1]
		array = append(array, node{0, left.val + right.val, "", left, right})
		array = append(array[2:])
		sort.Sort(array)
	}
	root := &array[0]
	encode(root, "")
	return codeTable
}

func encode(root *node, code string)  {
	root.code = code
	if root.left == nil && root.right == nil {
		codeTable[root.symbol] = root.code
		return
	}
	if root.left != nil {
		encode(root.left, code+"0")
	}
	if root.right != nil {
		encode(root.right, code+"1")
	}
}

func (a nodeArray) Less(i, j int) bool { return a[i].val < a[j].val }

func (a nodeArray) Len() int {
	return len(a)
}

func (a nodeArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}