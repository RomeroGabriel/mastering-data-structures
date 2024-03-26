# Tree Search

## Breadth-First Search

Breadth-First Search (BFS) is a graph traversal algorithm that `explores nodes level by level, starting from the root node and moving outward to the next levels`. It's particularly useful for `finding the shortest path in unweighted graphs or trees`.

BFS uses a [queue](../essential/queue.md#queue) to store nodes that need to be explored. The queue is used to keep track of nodes to be explored in `First In First Out (FIFO)`. This ensures that nodes are visited in the order they are discovered, which is crucial for maintaining the `level-by-level traversal` pattern. The `root node is enqueued first`.

??? example "Code Implementation"

    ```bash
    $ go run src/tree/bfs.go
    Is 2 present in treeNode?  false
    Is 1 present in treeNode?  true
    Is 3 present in treeNode?  true
    Is 4 present in treeNode?  true
    Is 8 present in treeNode?  true
    Is 9 present in treeNode?  false
    ```

    ```go
    --8<-- "src/tree/bfs.go"
    ```

## Depth-First Search
