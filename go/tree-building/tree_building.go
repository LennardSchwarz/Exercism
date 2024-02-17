package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// Sort records by ID
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	// Validate records and initialize nodes
	nodes := make(map[int]*Node)
	for i, record := range records {
		if record.ID != i {
			return nil, errors.New("records are not contiguous")
		}
		if record.ID == 0 && record.Parent != 0 {
			return nil, errors.New("root node has parent")
		}
		if record.ID != 0 && record.Parent >= record.ID {
			return nil, errors.New("node has invalid parent")
		}

		nodes[record.ID] = &Node{ID: record.ID}
	}

	// Build the tree
	for _, record := range records {
		if record.ID != 0 {
			parent := nodes[record.Parent]
			child := nodes[record.ID]
			parent.Children = append(parent.Children, child)
		}
	}

	return nodes[0], nil
}
