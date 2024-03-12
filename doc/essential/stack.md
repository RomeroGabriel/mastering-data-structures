# Stack

A Stack is a `linear` data structure that follows the **`Last In First Out (LIFO)`** principle, meaning the last element added to the stack will be the first one to be removed. This characteristic makes stacks particularly useful in scenarios where the `order of operations is reversed`, such as in function call stacks, undo operations in text editors, or backtracking algorithms.

## Interface

1. Push: Elements are `added` (inserted) at the top of the stack.
1. Pop: Elements are `removed` from the top of the stack.
1. Top/Peek: Accesses the `top element` of the stack without removing it.

## Time Complexity

1. Push: `O(1)` - Adding an element to the top of the stack is a constant-time operation.
1. Pop: `O(1)` - Removing an element from the top of the stack is also a constant-time operation.
1. Top/Peek: `O(1)` - Accessing the top element of the stack is a constant-time operation.

!!! example "Code Implementation"
    ```bash
    $ go run src/essential/stack.go
    Pop 9:  9
    Legth:  2
    Pop 11:  11
    Pop 7:  7
    Peek 5:  5
    Pop 5:  5
    Pop nothing:  <nil>
    Peek 70:  70
    Length:  1
    ```

    ```go
    --8<-- "src/essential/stack.go"
    ```
