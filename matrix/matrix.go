package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix is simple matrix representation.
type Matrix struct {
	col, row int
	values   [][]int
}

// New creates Matrix type.
func New(s string) (*Matrix, error) {
	var lines []string
	lines = strings.Split(s, "\n")
	if len(lines) == 0 {
		return &Matrix{}, errors.New("Empty input")
	}
	cols := len(strings.Fields(lines[0]))
	vals := make([][]int, len(lines))
	for i := range vals {
		vals[i] = make([]int, cols)
	}
	m := &Matrix{row: len(lines), col: cols, values: vals}
	var curRow int
	for _, line := range lines {
		if line == "" {
			return &Matrix{}, errors.New("Empty row received")
		}
		vals := strings.Fields(line)
		if len(vals) != cols {
			return &Matrix{}, errors.New("Invalid Input")
		}
		var col int
		for _, val := range vals {
			curVal, err := strconv.Atoi(val)
			if err != nil {
				return &Matrix{}, errors.New("Invalid Values")
			}
			m.values[curRow][col] = curVal
			col++
		}
		curRow++
	}
	return m, nil
}

// Set assigns element to row and column.
func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= m.row || col < 0 || col >= m.col {
		return false
	}
	m.values[row][col] = val
	return true
}

// Rows returns all rows of matrix.
func (m *Matrix) Rows() [][]int {
	rows := make([][]int, m.row)
	for i := range rows {
		rows[i] = make([]int, m.col)
		for j := 0; j < m.col; j++ {
			rows[i][j] = m.values[i][j]
		}
	}
	return rows
}

// Cols returns all columns of matrix.
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, m.col)
	for c := 0; c < m.col; c++ {
		cols[c] = make([]int, m.row)
		for r := 0; r < m.row; r++ {
			cols[c][r] = m.values[r][c]
		}
	}
	return cols
}
