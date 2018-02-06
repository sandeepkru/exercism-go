package circular

import "errors"

// Buffer is simple datastructure implements circular buffer
type Buffer struct {
	size  int
	tail  int
	head  int
	buf   []byte
	count int
}

// NewBuffer creates a new buffer with specified size.
func NewBuffer(size int) *Buffer {
	return &Buffer{size: size, tail: 0, head: 0, buf: make([]byte, size)}
}

// ReadByte returns byte from head if it contains valid data otherwise returns error.
func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 { // no elements exist.
		return ' ', errors.New("Buffer is empty")
	}
	val := b.buf[b.head]
	b.count--
	b.head++
	b.head = b.head % b.size
	return val, nil
}

// WriteByte writes error if there is space left in given buffer.
func (b *Buffer) WriteByte(c byte) error {
	if b.count == b.size {
		return errors.New("Buffer is full")
	}

	b.buf[b.tail] = c
	b.tail++
	b.count++
	b.tail = b.tail % b.size

	return nil
}

// Overwrite overwrites buffer at head.
func (b *Buffer) Overwrite(c byte) {
	if b.count < b.size { // Overwrite on to next if we have more space in buffer
		b.count++
	}
	b.buf[b.tail] = c
	// readjust head as we overwrote bytes.
	b.head = b.tail + 1
	b.head = b.head % b.size
	// move on tail
	b.tail++
	b.tail = b.tail % b.size
}

// Reset resets buffer
func (b *Buffer) Reset() { // put buffer in an empty state {
	b.tail = 0
	b.head = 0
	b.count = 0
}

// go test -run ^NOTHING -bench 'BenchmarkOverwrite|BenchmarkWriteRead'
// goos: darwin
// goarch: amd64
// BenchmarkOverwrite-8   	100000000	        20.6 ns/op	4862152591.62 MB/s
// BenchmarkWriteRead-8   	100000000	        20.2 ns/op	4952793059.67 MB/s
// PASS
// ok  	_/Users/sandeep/practice/go-exercism/go/circular-buffer	4.129s
