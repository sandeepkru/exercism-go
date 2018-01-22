package tree

import (
	"errors"
	"sort"
)

// Record is simple representation of stored recrod.
type Record struct {
	ID, Parent int
}

// Node define in-memory view of node structure.
type Node struct {
	ID       int
	Children []*Node
}

// Sorting Interface.
type onID []Record

func (records onID) Len() int {
	return len(records)
}

func (records onID) Swap(i, j int) {
	records[i], records[j] = records[j], records[i]
}

func (records onID) Less(i, j int) bool {
	return records[i].ID < records[j].ID
}

// Build builds inmemory tree from persisted recrods.
// Learnt approach from https://github.com/ThomasZumsteg/exercism-go/blob/master/tree-building/tree.go
func Build(records []Record) (*Node, error) {
	if len(records) <= 0 {
		return nil, nil
	}

	sort.Sort(onID(records))
	nodes := make([]*Node, len(records))
	for i, record := range records {
		nodes[i] = &Node{ID: record.ID}

		if i == 0 && (record.ID != 0 || record.Parent != 0) {
			return nil, errors.New("Invalid root record")
		} else if i == 0 { // root node validation successful, continue
			continue
		} else if i != record.ID || record.ID <= record.Parent { // to handle duplicates and incorrect relations
			return nil, errors.New("Invalid record")
		}

		if parent := nodes[record.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[i])
		}
	}

	return nodes[0], nil
}
