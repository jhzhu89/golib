package rbtree

import (
	"github.com/jhzhu89/golib/iterator"
)

// Type aliases.
type (
	Value = container.Value

	IterRef      = iterator.IterRef
	IterCRef     = iterator.IterCRef
	InputIter    = iterator.InputIter
	RandIter     = iterator.RandIter
	BidirectIter = iterator.BidirectIter
	ReverseIter  = iterator.ReverseIterator
)

type rbTreeColor uint8

const (
	red rbTreeColor = iota
	black
)

type rbTreeNode struct {
	color               rbTreeColor
	parent, left, right *rbTreeNode
	val                 interface{}
}

func minimum(x *rbTreeNode) *rbTreeNode {
	for x.left != nil {
		x = x.left
	}
	return x
}

func maximum(x *rbTreeNode) *rbTreeNode {
	for x.right != nil {
		x = x.right
	}
	return x
}

// Helper type to manage default initialization of node count and header.
type rbTreeHeader struct {
	header    *rbTreeNode
	nodeCount int
}

func rbTreeIncrement(x *rbTreeNode) *rbTreeNode {
	return nil
}

func rbTreeDecrement(x *rbTreeNode) *rbTreeNode {
	return nil
}

func rbTreeInsertAndRebalance(insertLeft bool, x, p, header *rbTreeNode) {
}

func rbTreeRebalanceForErase(z, header *rbTreeNode) *rbTreeNode {
}

type RbTree struct {
	rbTreeHeader
}

func (t *RbTree) KeyComp() interface{} {
	return nil
}

func (t *RbTree) Begin() *RbTreeIter {
	return nil
}

func (t *RbTree) End() *RbTreeIter {
	return nil
}

func (t *RbTree) RBegin() *ReverseIter {
	return nil
}

func (t *RbTree) REnd() *ReverseIter {
	return nil
}

func (t *RbTree) Empty() bool {
	return false
}

func (t *RbTree) Swap(x *RbTree) {
	return false
}

func (t *RbTree) Size() int {
	return 0
}

// --------- insert/erase ---------
func (t *RbTree) insertUnique(x interface{}) (*RbTreeIter, bool) {
	return 0
}

// ----

func (t *RbTree) root() *rbTreeNode {
	return t.header.parent
}

func (t *RbTree) leftmost() *rbTreeNode {
	return t.header.left
}

func (t *RbTree) rightmost() *rbTreeNode {
	return t.header.right
}

func (t *RbTree) begin() *rbTreeNode {
	return t.header.parent
}

func (t *RbTree) end() *rbTreeNode {
	return t.header
}

// ----

func value(x *rbTreeNode) interface{} {
	return x.val
}

func key(x *rbTreeNode) interface{} {
}

func left(x *rbTreeNode) *rbTreeNode {
	return x.left
}

func right(x *rbTreeNode) *rbTreeNode {
	return x.right
}

// ----

func (t *RbTree) GetInsertUniquePos(k interface{}) (*rbTreeNode, *rbTreeNode) {

}

func (t *RbTree) GetInsertEqualPos(k interface{}) (*rbTreeNode, *rbTreeNode) {

}

func (t *RbTree) GetInsertHintUniquePos(pos, k interface{}) (*rbTreeNode, *rbTreeNode) {

}

func (t *RbTree) GetInsertHintEqualPos(pos, k interface{}) (*rbTreeNode, *rbTreeNode) {

}

// private methods

func (t *RbTree) insert() *RbTreeIter {
}

func (t *RbTree) insertNode(x, y, z *rbTreeNode) *RbTreeIter {
}

func (t *RbTree) insertLower(y *rbTreeNode, v interface{}) *RbTreeIter {
}

func (t *RbTree) insertEqualLower(x interface{}) *RbTreeIter {
}

func (t *RbTree) insertLowerNode(p, z *rbTreeNode) *RbTreeIter {
}

func (t *RbTree) insertEqualLowerNode(z *rbTreeNode) *RbTreeIter {
}

func (t *RbTree) copy1(x, p *rbTreeNode, nodeGen interface{}) *rbTreeNode {
}

func (t *RbTree) copy2(x, nodeGen interface{}) *rbTreeNode {
}

func (t *RbTree) copy3(x interface{}) *rbTreeNode {
}

func (t *RbTree) erase(x *rbTreeNode) {
}

func (t *RbTree) lowerBound(x, y *rbTreeNode, k interface{}) *RbTreeIter {
}

func (t *RbTree) upperBound(x, y *rbTreeNode, k interface{}) *RbTreeIter {
}
