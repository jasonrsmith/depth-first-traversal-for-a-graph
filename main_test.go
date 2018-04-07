package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	testNodeCount := 4
	g := NewGraph(testNodeCount)
	assert.Equal(t, testNodeCount, len(g.EdgePairs))
	assert.Equal(t, 0, len(g.EdgePairs[0]))
}

func TestAddPair(t *testing.T) {
	testNodeCount := 4
	g := NewGraph(testNodeCount)
	g.AddPair(0, 1)
	g.AddPair(2, 3)
	assert.Equal(t, 1, g.EdgePairs[0][0])
	assert.Equal(t, 3, g.EdgePairs[2][0])
}

func TestAddPairPanicsWhenAddingPairExceedsNodeCount(t *testing.T) {
	testNodeCount := 4
	g := NewGraph(testNodeCount)
	g.AddPair(0, 1)
	assert.Panics(t, func() { g.AddPair(4, 1) })
}

func TestDFSSimple(t *testing.T) {
	testNodeCount := 2
	expectedPath := []int{0, 1}
	g := NewGraph(testNodeCount)
	g.AddPair(0, 1)
	path := g.FindPathDepthFirst(0)
	assert.Equal(t, expectedPath, path)
}

func TestDFSThreeNodes(t *testing.T) {
	testNodeCount := 3
	expectedPath := []int{0, 1, 2}
	g := NewGraph(testNodeCount)
	g.AddPair(0, 1)
	g.AddPair(1, 2)
	path := g.FindPathDepthFirst(0)
	assert.Equal(t, expectedPath, path)
}

func TestDFSBacktrack(t *testing.T) {
	testNodeCount := 5
	expectedPath := []int{0, 1, 2, 4, 3}
	g := NewGraph(testNodeCount)
	g.AddPair(0, 1)
	g.AddPair(0, 2)
	g.AddPair(0, 3)
	g.AddPair(2, 4)
	path := g.FindPathDepthFirst(0)
	assert.Equal(t, expectedPath, path)
}
