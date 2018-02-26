package react

// Got implementation from https://github.com/prolixe/exercism/blob/master/go/react/react.go

type cell struct {
	id         int
	value      int
	react      *reactor
	parents    []*cell
	children   []*cell
	function   func(...int) int
	callbacks  map[int]func(int)
	callbackID int
	dirty      bool
}

type reactor struct {
	cells []*cell
}

type canceler struct {
	cancel func()
}

// New returns new instance of reactor.
func New() Reactor {
	return &reactor{}
}

func (c *canceler) Cancel() {
	c.cancel()
}

func (r *reactor) CreateInput(i int) InputCell {
	newInput := &cell{
		id:    len(r.cells),
		value: i,
		react: r}
	r.cells = append(r.cells, newInput)
	return newInput
}

func (r *reactor) CreateCompute1(parent Cell, f func(int) int) ComputeCell {
	p := parent.(*cell)
	newComputeCell := &cell{
		id:        len(r.cells),
		value:     f(p.value),
		react:     r,
		parents:   []*cell{p},
		function:  func(args ...int) int { return f(args[0]) },
		callbacks: make(map[int]func(int)),
	}
	p.children = append(p.children, newComputeCell)
	r.cells = append(r.cells, newComputeCell)
	return newComputeCell
}

func (r *reactor) CreateCompute2(parent1, parent2 Cell, f func(int, int) int) ComputeCell {
	p1 := parent1.(*cell)
	p2 := parent2.(*cell)

	newComputeCell := &cell{
		id:        len(r.cells),
		value:     f(p1.value, p2.value),
		react:     r,
		parents:   []*cell{p1, p2},
		function:  func(args ...int) int { return f(args[0], args[1]) },
		callbacks: make(map[int]func(int)),
	}
	p1.children = append(p1.children, newComputeCell)
	p2.children = append(p2.children, newComputeCell)

	r.cells = append(r.cells, newComputeCell)
	return newComputeCell
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) update() {
	if c.dirty == false {
		panic("Called update on non dirty cell")
	}

	c.dirty = false
	pValues := make([]int, 0)
	for _, p := range c.parents {
		pValues = append(pValues, p.value)
	}
	oldValue := c.value

	c.value = c.function(pValues...)
	if c.value != oldValue {
		for _, child := range c.children {
			child.dirty = true
		}

		for _, callback := range c.callbacks {
			callback(c.value)
		}
	}
}

func (c *cell) SetValue(v int) {
	c.value = v
	c.react.update(c.id)
}

func (c *cell) AddCallback(f func(int)) Canceler {
	id := c.callbackID
	c.callbackID++
	c.callbacks[id] = f
	return &canceler{
		cancel: func() {
			delete(c.callbacks, id)
		},
	}
}

func (r *reactor) update(id int) {
	c := r.cells[id]
	var hasDirtyCell bool
	for _, child := range c.children {
		child.dirty = true
		hasDirtyCell = true
	}

	for hasDirtyCell {
		hasDirtyCell = false
		for _, c := range r.cells[id:] {
			if c.dirty {
				c.update()
				hasDirtyCell = true
			}
		}
	}
}
