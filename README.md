# optimal-search-trees

Constructs an optimal binary search tree for elements with given access
probabilities.  The nodes will be labeled 1 through n for simplicity.  For the
same reason, we assume that we will only search for things contained in the
tree.  The implementation uses a naive dynamic programming algorithm that runs
in O(n^3).  For a more efficient solution, see Knuth's O(n^2) algorithm.

The programme is implemented in [Go](https://golang.org).  The code is rather
verbose with lots of debugging code, and mostly untested.
