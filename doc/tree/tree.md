# Tree

## Terminoly

1. *Root*: The topmost parent node, also known as the `first node`.
2. *Height*: The longest path from the root to the most child node.
3. *Binary* tree: A tree structure where each node `has at most two children`, but can have zero children as well.
4. *General* tree: A tree with 0 or more children.
5. *Binary Search Tree*: A type of binary tree where `nodes are ordered in a specific way`, with at most two children
6. *Leaves*: nodes without children.
7. *Balanced Tree*: A tree is perfectly balanced when any node's left and right children have the `same height`.
8. *Branching Factor* - The amount of children a tree has.

## Binary Tree Comparison

??? example "Code Implementation"

    ```bash
    $ go run src/tree/compare_binary.go
    A and B is equal?  true
    A and B is equal?  false
    ```

    ```go
    --8<-- "src/tree/compare_binary.go"
    ```
