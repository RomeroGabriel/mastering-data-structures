# Tree Traversals

Tree traversal is a method used to `visit each node in a tree data structure exactly once`. There are several ways to traverse a tree, each with its own order of visiting nodes. The three primary tree traversal methods are `pre-order`, `in-order`, and `post-order`.

!!! note "Pre-Order"
    In pre-order traversal, the algorithm `visits the root node first`, then recursively `traverses the left subtree`, and finally the `right subtree`.

!!! note "In-Order"
    In in-order traversal, the algorithm `visits the nodes in the left subtree first`, followed by the `root node`, and finally the nodes in the `right subtree`.

!!! note "Post-Order"
    In post-order traversal, the algorithm visits the `left subtree first`, then the `right subtree`, and finally the `root node`.

!!! example
    ```bash
    $ go run src/tree/traversals.go
    PreOrder:   [5 3 1 4 7 6 8]
    InOrder:    [3 1 4 5 7 6 8]
    PostOrder:  [3 1 4 7 6 8 5]
    ```

    ```go
    --8<-- "src/tree/traversals.go"
    ```
