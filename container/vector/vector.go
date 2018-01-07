package vector

type Vector struct {
}

// Iterators

func (d *Vector) Begin() *VectorIter {
	return nil //clone(d.start)
}

func (d *Vector) End() *VectorIter {
	return nil //clone(d.finish)
}

func (d *Vector) RBegin() *ReverseIter {
	return nil //iterator.NewReverseIterator(d.finish)
}

func (d *Vector) REnd() *ReverseIter {
	return nil //iterator.NewReverseIterator(d.start)
}

// Element access

func (d *Vector) At(n int) Value {
	return nil
}

func (d *Vector) Front() Value {
	return d.At(0)
}

func (d *Vector) Back() Value {
	return nil
}

// Capacity

// Empty returns true if the Deuqe is empty.
func (d *Vector) Empty() bool {
	return true
}

func (d *Vector) Size() int {
	return 0
}

func (d *Vector) MaxSize() int {
	return 0
}

func (d *Vector) Reserve(newCap int) {
	return 0
}

func (d *Vector) Capacity() int {
	return 0
}

func (d *Vector) ShrinkToFit() bool {
	return false
}

// Modifiers
func (d *Vector) Clear() {
	return false
}

func (d *Vector) Insert(pos *VectorIter, val Value) *VectorIter {
	return nil
}

func (i insertFunc) Insert(it Iter, val Value) Iter {
	return nil
}

func (d *Vector) InsertRange(pos *VectorIter, first, last InputIter) *VectorIter {
	return nil
}

func (d *Vector) FillInsert(pos *VectorIter, n int, val Value) *VectorIter {
	return nil
}

func (d *Vector) Erase(pos *VectorIter) *VectorIter {
	return nil
}

func (d *Vector) EraseRange(first, last *VectorIter) *VectorIter {
	return nil
}

func (d *Vector) Swap(x *Vector) {

}

func (d *Vector) PushBack(val Value) {

}

func (d *Vector) PopBack(val Value) {

}

func (d *Vector) Resize(newSize int) {

}

func (d *Vector) ResizeAssign(newSize int, val Value) {

}
