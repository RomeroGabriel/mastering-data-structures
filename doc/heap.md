# Heap

A Heap is a specialized tree-based data structure that satisfies the `heap property`. It is a `complete binary tree, meaning all levels are fully filled except possibly for the last level, which is filled from left to right`. The heap property ensures that for every node, the value of its children is less than or equal to its own value. This property makes the `root node contain the maximum or minimum value`, depending on the type of heap, and the values decrease or increase as you move down the tree.

## Types of Heaps

1. `Max-Heap`: In a max-heap, the value of each node is greater than or equal to the values of its children. `The root node contains the maximum value`.
1. `Min-Heap`: In a min-heap, the value of each node is less than or equal to the values of its children. `The root node contains the minimum value`.

## Heap Operations

1. Insertion: This involves placing the new element at the end of the heap and then **`bubbling up`** to its correct position to maintain the heap property.
1. Deletion: `Removing the root element` (which is the maximum or minimum value depending on the type of heap). This involves replacing the root with the last element in the heap, then "bubbling down" to restore the heap property.
1. Heapify: `Adjusting the heap to maintain the heap property after an element has been removed or added`. This is done by comparing the removed or added element with its children and swapping them if necessary.

## Applications

1. Priority Queues: Heaps are often used to implement priority queues, where the highest (or lowest) priority element is always at the root of the heap.
1. Dijkstra's Algorithm: Heaps are used in Dijkstra's algorithm for finding the shortest path in a graph.
1. Heap Sort: Heap sort is a sorting algorithm that uses a heap to sort an array.

## Time Complexity

1. Insertion: `O(log n)` - Inserting a new element requires bubbling it up to its correct position, which takes logarithmic time in the size of the heap.
1. Deletion: `O(log n)` - Deleting the root element requires bubbling down the last element to restore the heap property, which also takes logarithmic time.
1. Heapify: `O(log n)` - Adjusting the heap to maintain the heap property after an element has been removed or added takes logarithmic time.

## Code

??? example "Code Implementation"

    ```bash
    $ go run src/heap.go
    Heap length:  0
    heap.length == 8:  8
    heap delete return 1 (smallest):  1
    heap delete return 3: 3
    heap delete return 4: 4
    heap delete return 5: 5
    heap.length == 4:  4
    heap delete return 7: 7
    heap delete return 8: 8
    heap delete return 69: 69
    heap delete return 420: 420
    heap.length == 0:  0
    ```

    ```go
    --8<-- "src/heap.go"
    ```
