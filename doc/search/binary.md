# Binary Search

Binary search is an efficient algorithm used for finding the position of a target value within a sorted array. `It works by repeatedly dividing the search interval in half`. If the target value is found, the algorithm `returns its index; otherwise, it returns -1`. This algorithm is particularly useful when dealing with large datasets because of its time complexity of O(log n), making it faster than linear search algorithms for large arrays.

!!! note "Time complexity"
    O(log n)

## Code

1. Initialization: The algorithm starts by initializing two variables, `lo (low)` and `hi (high)`, to the start and end indices of the data slice. These variables represent the current search space within the slice.
1. Loop Condition: `The algorithm enters a loop that continues as long as lo is less than hi`. This loop is the **`core of the binary search`**, where the search space is repeatedly halved.
1. Calculating the Middle Index: Inside the loop, the `middle index of the current search space is calculated`. This calculation ensures that the middle index is always **`rounded down`** to the nearest whole number, which is important for array indexing.
1. Comparing the Middle Value:
    - If the middle value is `equal to searchValue`, the algorithm `returns` the middle index, indicating that the search value has been found.
    - If the middle value is `greater than searchValue`, the search continues in the lower half of the current search space. `This is done by setting hi to the middle index`, effectively excluding the upper half from the search space.
    - If the middle value is `less than searchValue`, the search continues in the upper half of the current search space. `This is done by setting lo to the middle index plus one`, effectively excluding the lower half from the search space.
1. Loop Termination: `The loop continues until lo is no longer less than hi, which means the search space has been exhausted without finding the searchValue`. At this point, the algorithm returns -1, indicating that the search value is not present in the data slice.

!!! example
    ```bash
    $ go run src/search/binary.go
    Result for 69:  3
    Result for 1336:  -1
    Result for 69420:  10
    Result for 69421:  -1
    Result for 1:  0
    Result for 0:  -1
    ```

    ```go
    --8<-- "src/search/binary.go"
    ```
