# Double Linked List

A Doubly Linked List is a type of linked list where `each node contains a reference to both the next node and the previous node in the sequence`. This allows for more efficient operations compared to a singly linked list, where each node only has a reference to the next node.

## How It Works

1. *Nodes*: Each node in a doubly linked list contains data and `two pointers`, one `pointing to the next node and another pointing to the previous node`.
1. *Head and Tail*: The list has a `head` pointer pointing to the `first node` and a `tail` pointer pointing to the `last node`.
1. *Insertion*: Nodes can be inserted at the `beginning`, at the `end`, or at `any given position` within the list.
1. *Deletion*: Nodes can be deleted from the `beginning`, from the `end`, or from `any given position` within the list.

## Time Complexity

### Insertion

1. At the Beginning: `O(1)` - Updating the head pointer is a constant-time operation.
1. At the End: `O(1)` - Updating the tail pointer is a constant-time operation.
1. At a Given Position: `O(n)` - Requires traversal to the given position, which can involve `visiting all elements in the worst case`.

### Deletion

1. At the Beginning: `O(1)` - Updating the head pointer is a constant-time operation.
2. At the End: `O(1)` - Updating the tail pointer is a constant-time operation.
3. At a Given Position: `O(n)` - Requires traversal to the given position, which can involve `visiting all elements in the worst case`.

### Accessing

1. By Index: `O(n)` - Requires traversal to the given position, which can involve `visiting all elements in the worst case`.

## Interface Operations

1. Prepend: `O(1)` - Inserts an element at the `beginning` of the list.
2. InsertAt: `O(n)` - Inserts an element at a `given position`.
3. Append: `O(1)`  - Inserts an element at the `end` of the list.
4. Remove: `O(n)` - Removes a specific element from the list. -> MAYBE FIX THIS
5. RemoveAt: `O(n)` Removes an element at a `given position`.
6. Get: `O(n)` - Accesses an element at a `given index`.

## Code

!!! tip
    First focus on the new element or the remove element. After that, focus on the sibbiling nodes to point to the new direction.

!!! example
    ```bash
    $ go run src/linked_list/doubly_list.go
    ```

    ```go
    --8<-- "src/linked_list/doubly_list.go"
    ```
