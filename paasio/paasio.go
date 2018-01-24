package paasio

import (
	"io"
	"sync"
)

type counter struct {
	totalBytes int64
	nops       int
	mutex      *sync.Mutex
}

type paasioReader struct {
	reader io.Reader
	counter
}

type paasioWriter struct {
	writer io.Writer
	counter
}

type paasioReaderWriter struct {
	WriteCounter
	ReadCounter
}

func (c *counter) getCount() (int64, int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.totalBytes, c.nops
}

func (c *counter) addCount(byteCount int64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.totalBytes += byteCount
	c.nops++
}

func (r *paasioReader) Read(p []byte) (int, error) {
	count, err := r.reader.Read(p)
	r.counter.addCount(int64(count))
	return count, err
}

func (r *paasioReader) ReadCount() (n int64, nops int) {
	n, nops = r.counter.getCount()
	return
}

func (w *paasioWriter) Write(p []byte) (int, error) {
	count, err := w.writer.Write(p)
	w.counter.addCount(int64(count))
	return count, err
}

func (w *paasioWriter) WriteCount() (n int64, nops int) {
	n, nops = w.counter.getCount()
	return
}

func newCounter() counter {
	return counter{mutex: &sync.Mutex{}}
}

// NewWriteCounter creates simple paasio writer.
func NewWriteCounter(writer io.Writer) WriteCounter {
	return &paasioWriter{writer: writer, counter: newCounter()}
}

// NewReadCounter creates simple paasio reader.
func NewReadCounter(r io.Reader) ReadCounter {
	return &paasioReader{reader: r, counter: newCounter()}
}

// NewReadWriteCounter creates simple reader writer.
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &paasioReaderWriter{
		NewWriteCounter(rw),
		NewReadCounter(rw),
	}
}
