package ast

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/PrawnSkunk/genetic-al-go-rithm/lib/snoise2"
)

// Node evaluated an x and y
type Node interface {
	Eval(x, y float32) float32
	String() string
	AddRandom(node Node)
	NodeCounts() (nodeCount, nilCount int)
}

// LeafNode has no children
type LeafNode struct{}

// AddRandom fulfills LeafNode interface
func (leaf *LeafNode) AddRandom(node Node) {
	// Do not attempt to add a node to a leaf node
}

// NodeCounts fulfills LeafNode interface
func (leaf *LeafNode) NodeCounts() (nodeCount, nilCount int) {
	return 1, 0 // Leaf node has no nil pointers
}

// SingleNode has one child
type SingleNode struct {
	Child Node
}

// AddRandom fulfills SingleNode interface
func (single *SingleNode) AddRandom(node Node) {
	if single.Child == nil {
		single.Child = node // No child has been assigned
	} else {
		single.Child.AddRandom(node) // Child has already been assigned, pass the node along
	}
}

// NodeCounts fulfills SingleNode interface
func (single *SingleNode) NodeCounts() (nodeCount, nilCount int) {
	if single.Child == nil {
		nodeCount = 1
		nilCount = 1
	} else {
		childNodeCount, childNilCount := single.Child.NodeCounts()
		nodeCount = 1 + childNodeCount // Return itself + count
		nilCount = childNilCount
	}
	return nodeCount, nilCount
}

// DoubleNode has two children
type DoubleNode struct {
	LeftChild  Node
	RightChild Node
}

// AddRandom fulfills DoubleNode interface
func (double *DoubleNode) AddRandom(node Node) {
	// We have a choice of adding it to the right or left side
	r := rand.Intn(2) // Number between 0 and 1
	if r == 0 {
		// Add left side
		if double.LeftChild == nil {
			double.LeftChild = node
		} else {
			double.LeftChild.AddRandom(node)
		}
	} else {
		// Add right side
		if double.RightChild == nil {
			double.RightChild = node
		} else {
			double.RightChild.AddRandom(node)
		}
	}
}

// NodeCounts fulfills DoubleNode interface
func (double *DoubleNode) NodeCounts() (nodeCount, nilCount int) {
	var leftCount, leftNilCount, rightCount, rightNilCount int
	if double.LeftChild == nil {
		// No left child
		leftNilCount = 1
		leftCount = 0
	} else {
		// There is a left child
		leftCount, leftNilCount = double.LeftChild.NodeCounts()
	}

	if double.RightChild == nil {
		// No right child
		rightNilCount = 1
		rightCount = 0
	} else {
		// There is a right child
		rightCount, rightNilCount = double.RightChild.NodeCounts()
	}
	return 1 + leftCount + rightCount, leftNilCount + rightNilCount // +1 is ourselves
}

// TripleNode has two children
type TripleNode struct {
	LeftChild   Node
	MiddleChild Node
	RightChild  Node
}

// AddRandom fulfills TripleNode interface
func (triple *TripleNode) AddRandom(node Node) {
	// We have a choice of adding it to three children
	r := rand.Intn(3)
	if r == 0 {
		// Add left node
		if triple.LeftChild == nil {
			triple.LeftChild = node
		} else {
			triple.LeftChild.AddRandom(node)
		}
	} else if r == 1 {
		// Add middle node
		if triple.MiddleChild == nil {
			triple.MiddleChild = node
		} else {
			triple.MiddleChild.AddRandom(node)
		}
	} else {
		// Do not add a node to a constant node
		if triple.RightChild == nil {
			triple.RightChild = node
		} else {
			triple.RightChild.AddRandom(node)
		}
	}
}

// NodeCounts fulfills TripleNode interface
func (triple *TripleNode) NodeCounts() (nodeCount, nilCount int) {
	var leftCount, leftNilCount, middleCount, middleNilCount, rightCount, rightNilCount int
	// Handle left branch
	if triple.LeftChild == nil {
		// No left child
		leftNilCount = 1
		leftCount = 0
	} else {
		// There is a left child
		leftCount, leftNilCount = triple.LeftChild.NodeCounts()
	}

	// Handle middle branch
	if triple.MiddleChild == nil {
		// No middle child
		middleNilCount = 1
		middleCount = 0
	} else {
		// There is a middle child
		middleCount, middleNilCount = triple.MiddleChild.NodeCounts()
	}

	// Handle right branch
	if triple.RightChild == nil {
		// No right child
		rightNilCount = 1
		rightCount = 0
	} else {
		// There is a right child
		rightCount, rightNilCount = triple.RightChild.NodeCounts()
	}

	return 1 + leftCount + middleCount + rightCount, leftNilCount + middleNilCount + rightNilCount // + 1 is ourself
}

// OpConstant represents a constant value
type OpConstant struct {
	LeafNode
	value float32
}

// Eval fulfills OpConstant interface
func (op *OpConstant) Eval(x, y float32) float32 {
	return op.value
}

func (op *OpConstant) String() string {
	return strconv.FormatFloat(float64(op.value), 'f', 9, 32) // Converts floating point to a string
}

// OpX represents some variable x
type OpX struct {
	LeafNode
}

// Eval fulfills OpX interface
func (op *OpX) Eval(x, y float32) float32 {
	return x
}

func (op *OpX) String() string {
	return "X"
}

// OpY represents some variable y
type OpY struct {
	LeafNode
}

// Eval fulfills OpY interface
func (op *OpY) Eval(x, y float32) float32 {
	return y
}

func (op *OpY) String() string {
	return "Y"
}

// OpPlus adds two nodes
type OpPlus struct {
	DoubleNode
}

// Eval fulfills OpPlus interface
func (op *OpPlus) Eval(x, y float32) float32 {
	return op.LeftChild.Eval(x, y) + op.RightChild.Eval(x, y)
}

func (op *OpPlus) String() string {
	return "( + " + op.LeftChild.String() + " " + op.RightChild.String() + " )"
}

// OpMinus subtracts two nodes
type OpMinus struct {
	DoubleNode
}

// Eval fulfills OpMinus interface
func (op *OpMinus) Eval(x, y float32) float32 {
	return op.LeftChild.Eval(x, y) - op.RightChild.Eval(x, y)
}

func (op *OpMinus) String() string {
	return "( - " + op.LeftChild.String() + " " + op.RightChild.String() + " )"
}

// OpMult multiplies two nodes
type OpMult struct {
	DoubleNode
}

// Eval fulfills OpMult interface
func (op *OpMult) Eval(x, y float32) float32 {
	return op.LeftChild.Eval(x, y) * op.RightChild.Eval(x, y)
}

func (op *OpMult) String() string {
	return "( * " + op.LeftChild.String() + " " + op.RightChild.String() + " )"
}

// OpDiv divides two nodes
type OpDiv struct {
	DoubleNode
}

// Eval fulfills OpDiv interface
func (op *OpDiv) Eval(x, y float32) float32 {
	return op.LeftChild.Eval(x, y) / op.RightChild.Eval(x, y)
}

func (op *OpDiv) String() string {
	return "( / " + op.LeftChild.String() + " " + op.RightChild.String() + " )"
}

// OpSine calculates the sine of a given node
type OpSine struct {
	SingleNode
}

// Eval fulfills OpSine interface
func (op *OpSine) Eval(x, y float32) float32 {
	return float32(math.Sin(float64(op.Child.Eval(x, y))))
}

func (op *OpSine) String() string {
	return "( Sine " + op.Child.String() + " )"
}

// OpCos calculates the cosine of a given node
type OpCos struct {
	SingleNode
}

// Eval fulfills OpSine interface
func (op *OpCos) Eval(x, y float32) float32 {
	return float32(math.Cos(float64(op.Child.Eval(x, y))))
}

func (op *OpCos) String() string {
	return "( Cos " + op.Child.String() + " )"
}

// OpAtan calculates the atangent of a given node
type OpAtan struct {
	SingleNode
}

// Eval fulfills OpAtan interface
func (op *OpAtan) Eval(x, y float32) float32 {
	return float32(math.Atan(float64(op.Child.Eval(x, y))))
}

func (op *OpAtan) String() string {
	return "( Atan " + op.Child.String() + " )"
}

// OpAtan2 calculates the atangent of two nodes
type OpAtan2 struct {
	DoubleNode
}

// Eval fulfills OpAtan2 interface
func (op *OpAtan2) Eval(x, y float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

func (op *OpAtan2) String() string {
	return "( Atan2 " + op.LeftChild.String() + " " + op.RightChild.String() + " )"
}

// OpNoise generates noise given two nodes
type OpNoise struct {
	DoubleNode
}

// Eval fulfills OpNoise interface
func (op *OpNoise) Eval(x, y float32) float32 {
	// (80*) makes it between 0-2
	// (- 2.0) maks it between -1 and 1
	return 80*snoise2.Snoise2(op.LeftChild.Eval(x, y), op.RightChild.Eval(x, y)) - 2.0
}

func (op *OpNoise) String() string {
	return "( SimplexNoise " + op.LeftChild.String() + " " + op.RightChild.String() + " )"
}

// OpLerp interpolates between two nodes (and a percentage)
type OpLerp struct {
	TripleNode
}

// Eval fulfills OpLerp interface
func (op *OpLerp) Eval(x, y float32) float32 {
	b1 := op.LeftChild.Eval(x, y)
	b2 := op.MiddleChild.Eval(x, y)
	pct := float32(math.Abs(float64(op.RightChild.Eval(x, y))))
	return b1 + pct*(b1-b2)
}

func (op *OpLerp) String() string {
	return "( Lerp " + op.LeftChild.String() + " " + op.MiddleChild.String() + " " + op.RightChild.String() + " )"
}

// GetRandomNode selects a random node with 1 or more children
func GetRandomNode() Node {
	// Non-leaf nodes only
	r := rand.Intn(10)
	switch r {
	case 0:
		return &OpPlus{}
	case 1:
		return &OpMinus{}
	case 2:
		return &OpMult{}
	case 3:
		return &OpDiv{}
	case 4:
		return &OpAtan2{}
	case 5:
		return &OpAtan{}
	case 6:
		return &OpCos{}
	case 7:
		return &OpSine{}
	case 8:
		return &OpNoise{}
	case 9:
		return &OpLerp{}
	}
	panic("Get Random Node Failed")
}

// GetRandomLeaf selects a random leaf node
func GetRandomLeaf() Node {
	r := rand.Intn(3)
	switch r {
	case 0:
		return &OpX{}
	case 1:
		return &OpY{}
	case 2:
		return &OpConstant{LeafNode{}, rand.Float32()*2 - 1} // Between 0-1
	}
	panic("Get Random Leaf Failed")
}
