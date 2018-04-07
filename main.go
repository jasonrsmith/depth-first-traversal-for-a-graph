package main

type Graph struct {
	NodeCount int
	EdgePairs [][]int
}

func NewGraph(nodeCount int) Graph {
	edgePairs := make([][]int, nodeCount)
	for i := range edgePairs {
		edgePairs[i] = make([]int, 0)
	}
	return Graph{
		nodeCount,
		edgePairs,
	}
}

func (g *Graph) AddPair(a, b int) {
	if a >= g.NodeCount || b >= g.NodeCount {
		panic("Value in pair exceeds node count")
	}
	g.EdgePairs[a] = append(g.EdgePairs[a], b)
}

func (g *Graph) FindPathDepthFirst(n int) []int {
	path := make([]int, 0)
	if n >= g.NodeCount {
		panic("Node to search is out of bounds")
	}
	path = g.findPathDepthFirstHelp(n, path)
	return path
}

func (g *Graph) findPathDepthFirstHelp(n int, path []int) []int {
	path = append(path, n)
	for i := range g.EdgePairs[n] {
		path = g.findPathDepthFirstHelp(g.EdgePairs[n][i], path)
	}
	return path
}

func main() {
}

/*
Write a function to print the depth first traversal for a undirected graph from a given source s.

Input:
The task is to complete the function DFS which takes 3 arguments an integer denoting the starting node (s) of the dfs travel , a  graph (g)  and an array of visited nodes (vis)  which are initially all set to false .
There are multiple test cases. For each test case, this method will be called individually.

Output:
The function should print the depth first traversal for the graph from the given source.

Note : The expected output button always produces DFS starting from node 1.

Constraints:
1 <=T<= 100
1 <=Number of  edges<= 100

Example:
Input:
1
4
1 2 1 3 1 4 3 5

Output:
1 2 3 5 4    //dfs from node 1


There is  one test cases.  First line of each test case represent an integer N denoting no of edges and then in the next line N pairs of values a and b are fed which denotes there is an edge from a to b .
*/
