# Queue

A Queue data structure operates based on the `First In First Out (FIFO)` principle, meaning that the first element added to the queue will be the first one to be removed. This characteristic makes queues particularly useful in scenarios `where order matters`, such as scheduling processes, handling requests in a server, or managing tasks in a print queue.

## Interface

1. Enqueue: Elements are `added` (inserted) at the end of the queue.
1. Dequeue: Elements are `removed` from the front of the queue.
1. Peek: Accesses the `first element` in the queue without removing it.
1. Rear: Accesses the `last element` in the queue.

Queues can be implemented using various data structures, including `arrays` or `linked lists`. The choice of implementation can affect the time complexity of operations like enqueue and dequeue.

## Time Complexity

1. Enqueue: `O(1)` - Adding an element to the end of the queue is a constant-time operation.
1. Dequeue: `O(1)` - Removing an element from the front of the queue is also a constant-time operation.
1. Peek/Rear: `O(1)` - Accessing the first or last element of the queue is a constant-time operation.

## Code

!!! example
    ```bash
    $ go run src/essential/queue.go
    Deque 5 (FIFO):  5
    Lenght:  2
    Deque 7 (FIFO):  7
    Deque 9 (FIFO):  9
    Peek 11:  11
    Deque 11 (FIFO):  11
    Deque nill:  <nil>
    Lenght:  0
    ```

    ```go
    --8<-- "src/essential/queue.go"
    ```
