package deque

type Deque struct {
}

type impl struct {
	map_          []node
	start, finish DequeIter
}

func (i *impl) initializeMap(numElements int) {

}

func (i *impl) createNodes(start, finish int) {

}

func (i *impl) destroyNodes(start, finish int) {

}

func (i *impl) allocateNode() {
}

func (i *impl) deallocateNode() {
}

func (i *impl) allocateMap() {
}

func (i *impl) deallocateMap() {
}

type node []Value
