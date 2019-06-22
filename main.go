package main

import (
	"fmt"
	"math"
)

type vector []float64
type field [][]float64

func make2dSliceOfInt(n int) [][]int {
	P := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		P[i] = make([]int, n+1)
	}
	return P
}

func make2dSliceOfDouble(n int) field {
	P := make(field, n+1)
	for i := 1; i <= n; i++ {
		P[i] = make([]float64, n+1)
	}
	return P
}

func print2dSliceOfInt(slice [][]int, name string) {
	n := len(slice) - 1

	fmt.Printf(" %s | ", name)
	for i := 1; i <= n; i++ {
		fmt.Printf("%2d ", i)
	}
	fmt.Println()
	for i := 1; i <= n; i++ {
		fmt.Print("---")
	}
	fmt.Println("----")

	for i := 1; i <= n; i++ {
		fmt.Printf("%2d | ", i)
		for j := 1; j <= n; j++ {
			fmt.Printf("%2d ", slice[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func print2dSliceOfDouble(slice field, name string) {
	n := len(slice) - 1

	fmt.Printf(" %s | ", name)
	for i := 1; i <= n; i++ {
		fmt.Printf("%4d ", i)
	}
	fmt.Println()
	for i := 1; i <= n; i++ {
		fmt.Print("-----")
	}
	fmt.Println("-----")

	for i := 1; i <= n; i++ {
		fmt.Printf("%2d | ", i)
		for j := 1; j <= n; j++ {
			fmt.Printf("%1.2f ", slice[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

}

// obviously correct
func initialize(γ []float64) (field, field, [][]int) {
	n := len(γ)
	P := make2dSliceOfDouble(n)
	Γ := make2dSliceOfDouble(n)
	R := make2dSliceOfInt(n)

	for i := 1; i <= n; i++ {
		//P[i][i-1] = 0
		P[i][i] = γ[i-1]
		Γ[i][i] = γ[i-1]
		R[i][i] = i
	}

	return P, Γ, R
}

func computeCore(γ vector, P, Γ field, R [][]int) {
	n := len(γ)
	for d := 1; d <= n-1; d++ {
		for i := 1; i <= n-d; i++ {
			j := i + d
			root := i
			t := math.Inf(1)

			for k := i; k <= j; k++ {
				x := P[i][k-1]
				if k+1 <= n {
					x += P[k+1][j]
				}
				if x <= t {
					t = x
					root = k
				}
			}

			Γ[i][j] = Γ[i][j-1] + γ[j-1]
			P[i][j] = t + Γ[i][j]
			R[i][j] = root
		}
	}
}

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func extractTree(R [][]int, i, j int) *Tree {
	if i > j {
		return nil
	}

	root := R[i][j]

	return &Tree{
		root,
		extractTree(R, i, root-1),
		extractTree(R, root+1, j),
	}
}

func computeOptimalSearchTree(γ []float64) (tree *Tree, cost float64) {
	P, Γ, R := initialize(γ)

	computeCore(γ, P, Γ, R)

	//println("P[i][j]: weighted inner path length of an optimal search tree for the node set {i,…,j}")
	//print2dSliceOfDouble(P, "P")

	fmt.Println("R[i][j]: root of an optimal search tree for {i,…,j}")
	print2dSliceOfInt(R, "R")

	return extractTree(R, 1, len(γ)), P[1][len(γ)]
}

func formatTree(tree *Tree) string {
	if tree == nil {
		return "()"
	} else if tree.left == nil && tree.right == nil {
		return fmt.Sprint(tree.value)
	}

	return fmt.Sprintf("(%v %v %v)", tree.value, formatTree(tree.left), formatTree(tree.right))
}

func main() {
	//γ := []float64{0.2, 0.3, 0.4, 0.04, 0.06}
	γ := []float64{0.18, 0.22, 0.15, 0.1, 0.06, 0.04, 0.25}
	tree, weight := computeOptimalSearchTree(γ)
	fmt.Println(formatTree(tree), weight)
}
