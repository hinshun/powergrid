package standard

import (
	"encoding/csv"
	"errors"
	"strconv"
	"strings"

	"github.com/hinshun/powergrid"
)

// standardBoard implements the powergrid.Board interface
type standardBoard struct {
	nodes []powergrid.Node
}

// NewBoard constructs a standard Board
func NewBoard(nodes []powergrid.Node) powergrid.Board {
	return &standardBoard{
		nodes: nodes,
	}
}

// Nodes returns the nodes in the board
func (b *standardBoard) Nodes() []powergrid.Node {
	return b.nodes
}

// a standardNode is a standard implementation of the powergrid.Node interface
type standardNode struct {
	name        string
	connections []powergrid.Connection
	cities      []powergrid.Player
}

// Name returns the name of the node
func (n *standardNode) Name() string {
	return n.name
}

// Connections returns the connections of the node
func (n *standardNode) Connections() []powergrid.Connection {
	return n.connections
}

// Connections returns the cities on the node
func (n *standardNode) Cities() []powergrid.Player {
	return n.cities
}

// NewBoardFromCSV constructs a standard Board using a string
// where each line is of the form:
// <node name>, <weight>, <node name>
// Each edge only needs to appear in the CSV once; this method will
// not deduplicate multiple occurrences of the same edge
// TODO(rabrams) refactor
func NewBoardFromCSV(s string) (powergrid.Board, error) {
	board := &standardBoard{}
	ioReader := strings.NewReader(s)
	csvReader := csv.NewReader(ioReader)
	record, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	nodesByName := make(map[string]*standardNode)
	for i := range record {
		line := record[i]
		if len(line) != 3 {
			return nil, errors.New("line must be a triplet")
		}
		left := strings.Trim(line[0], " ")
		right := strings.Trim(line[2], " ")
		weight, err := strconv.ParseUint(strings.Trim(line[1], " "), 10, 64)
		if err != nil {
			return nil, err
		}
		leftNode, ok := nodesByName[left]
		if !ok {
			leftNode = &standardNode{
				name: left,
			}
			nodesByName[left] = leftNode
			board.nodes = append(
				board.nodes,
				leftNode)
		}
		rightNode, ok := nodesByName[right]
		if !ok {
			rightNode = &standardNode{
				name: right,
			}
			nodesByName[right] = rightNode
			board.nodes = append(
				board.nodes,
				rightNode)
		}
		leftNode.connections = append(leftNode.connections,
			powergrid.Connection{
				Cost: powergrid.Elektro(weight),
				Node: rightNode,
			})
		rightNode.connections = append(
			rightNode.connections, powergrid.Connection{
				Cost: powergrid.Elektro(weight),
				Node: leftNode,
			})
	}
	return board, nil
}
