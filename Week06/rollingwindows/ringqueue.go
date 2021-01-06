package rollingwindows

type RingQueue struct {
	data []interface{}
	head int
	tail int
	maxSize int
}

func NewRingQueue(maxSize int) *RingQueue {
	return &RingQueue{
		data:		make([]interface{}, maxSize),
		head:		0,
		tail:		0,
		maxSize:	maxSize,
	} 
}

func (r *RingQueue) Push(data interface{}) {
	r.data[r.tail] = data
	r.tail = (r.tail + 1) % r.maxSize

	if r.isFull() {
		r.head = r.tail + 1		
	}
}

func (r *RingQueue) Pop() interface{} {
	if r.tail == r.head {
		return nil
	}
	data := r.data[r.head]
	r.data[r.head] = nil
	r.head = (r.head + 1) % r.maxSize
	return data
}

func (r *RingQueue) isFull() bool {
	return (r.tail + 1) % r.maxSize == r.head
}

func (r *RingQueue) curSize() int {
	return (r.tail - r.head + r.maxSize) % r.maxSize
}


func (r *RingQueue) GetLast() interface{} {
	if r.head == r.tail {
		return nil
	} else {
		return r.data[(r.tail-1+r.maxSize)%r.maxSize]
	}
}

func (r *RingQueue) Clear() {
	cur := r.curSize()
	for i := 0; i < cur; i++ {
		r.Pop()
	}
}