# Quick Sort

QuickSort is a `divide-and-conquer` sorting algorithm that works by selecting a *`pivot element`* from the array and *`partitioning the other elements into two sub-arrays`*, according to whether they are `less than or greater than the pivot`. The sub-arrays are then recursively sorted.

## How Works

1. *Choosing a Pivot*: The pivot can be any element from the array. Common choices include the `first element`, the `last element`, or a `random element`. The **`choice of pivot can significantly affect the algorithm's performance`**.
1. *Partitioning*: The array is partitioned around the pivot, such that all elements `less than the pivot come before it, and all elements greater than the pivot come after it`. This is done in-place, meaning no additional space is required.
1. *Recursive Sorting*: The partitioning step divides the array into two sub-arrays. QuickSort is then recursively applied to each sub-array.

??? example "Illustration"
    <figure markdown="span">
        ![Example of Quick Sort from geeksforgeeks](https://media.geeksforgeeks.org/wp-content/uploads/20231219164812/Quick-Sort-Algorithm.png){ width="600" }
        <figcaption>[*Example of Quick Sort from geeksforgeeks*](https://www.geeksforgeeks.org/quick-sort-in-c/)</figcaption>
    </figure>

## Time Complexity

1. **Best Case: O(N log N)**: This occurs when the pivot chosen at each step **`divides the array into roughly equal half`**. In this case, the algorithm will make `balanced partitions`, leading to efficient sorting.
1. **Worst Case: O(N^2)**: The worst-case scenario for QuickSort occurs when the pivot at each step consistently results in `highly unbalanced partitions`, such as when the array is already sorted and the pivot is always chosen as the smallest or largest element. This can lead to the algorithm performing poorly.

## Code

!!! example "Implementing Quick Sort"
    ```bash
    $ go run src/sort/quick_sort.go
    ```

    ```go
    --8<-- "src/sort/quick_sort.go"
    ```
