package standard

// TODO(rabrams) import and use ginkgo
import (
	"reflect"
	"testing"

	"github.com/hinshun/powergrid"
)

const sampleCSVNoRecurrence = `Phoenix, 14, San Diego
Denver, 16, Kansas City`

const sampleCSVLeftRecurrence = `Denver, 0, Cheyenne
Denver, 16, Kansas City`

const sampleCSVRightRecurrence = `Phoenix, 14, San Diego
Los Angeles, 3, San Diego`

func TestCreateBoardFromCSVNoRecurrence(t *testing.T) {
	board, err := NewBoardFromCSV(sampleCSVNoRecurrence)
	if err != nil {
		t.Fail()
	}
	nodesByName := make(map[string]powergrid.Node)
	nodes := board.Nodes()
	for i := range nodes {
		node := nodes[i]
		nodesByName[node.Name()] = node
	}
	// maps place name to its expected connection
	testCases := map[string]powergrid.Connection{
		"Phoenix": powergrid.Connection{
			Cost: 14,
			Node: nodesByName["San Diego"],
		},
		"San Diego": powergrid.Connection{
			Cost: 14,
			Node: nodesByName["Phoenix"],
		},
		"Denver": powergrid.Connection{
			Cost: 16,
			Node: nodesByName["Kansas City"],
		},
		"Kansas City": powergrid.Connection{
			Cost: 16,
			Node: nodesByName["Denver"],
		},
	}
	for startName, expectedConnection := range testCases {
		expectedConnections := []powergrid.Connection{
			expectedConnection,
		}
		node := nodesByName[startName]
		if !reflect.DeepEqual(expectedConnections, node.Connections()) {
			t.Fail()
		}
	}
}

func TestCreateBoardFromCSVLeftRecurrence(t *testing.T) {
	board, err := NewBoardFromCSV(sampleCSVLeftRecurrence)
	if err != nil {
		t.Fail()
	}
	nodesByName := make(map[string]powergrid.Node)
	nodes := board.Nodes()
	for i := range nodes {
		node := nodes[i]
		nodesByName[node.Name()] = node
	}
	denverNode := nodesByName["Denver"]
	if len(denverNode.Connections()) != 2 {
		t.Fail()
	}
	connections := denverNode.Connections()
	for i := range connections {
		connection := connections[i]
		if connection.Node.Name() == "Cheyenne" {
			if connection.Cost != 0 {
				t.Fail()
			}
		} else if connection.Node.Name() == "Kansas City" {
			if connection.Cost != 16 {
				t.Fail()
			}
		} else {
			t.Fail()
		}
	}
}

func TestCreateBoardFromCSVRightRecurrence(t *testing.T) {
	board, err := NewBoardFromCSV(sampleCSVRightRecurrence)
	if err != nil {
		t.Fail()
	}
	nodesByName := make(map[string]powergrid.Node)
	nodes := board.Nodes()
	for i := range nodes {
		node := nodes[i]
		nodesByName[node.Name()] = node
	}
	sdNode := nodesByName["San Diego"]
	if len(sdNode.Connections()) != 2 {
		t.Fail()
	}
	connections := sdNode.Connections()
	for i := range connections {
		connection := connections[i]
		if connection.Node.Name() == "Los Angeles" {
			if connection.Cost != 3 {
				t.Fail()
			}
		} else if connection.Node.Name() == "Phoenix" {
			if connection.Cost != 14 {
				t.Fail()
			}
		} else {
			t.Fail()
		}
	}
}
