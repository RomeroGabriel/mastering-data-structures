# Bubble Sort

Bubble Sort is a simple sorting algorithm that `works by repeatedly swapping adjacent elements if they are in the wrong order`. The time complexity of Bubble Sort is `O(n^2)`, where `n` is the number of elements in the array. This is because, in the worst case, the algorithm makes `n(n-1)/2` comparisons and swaps, which is the sum of the first n natural numbers.

!!! note "Time complexity"
    O(n^2)

!!! warning
    Despite its simplicity, Bubble Sort is not efficient for large datasets due to its high time complexity.

## Code

1. Comparison and Swapping: The algorithm starts by comparing the first element with the second element. If the first element is greater than the second, they are swapped. `This process is repeated for each pair of adjacent elements in the array`.
1. Passes: This process is repeated for each element in the array, with the **`largest element bubbling up to its correct position at the end of the array`**. This is done for the number of passes equal to the number of elements in the array.
1. Optimization: `Since the largest elements are already at the end of the array after the first pass, there's no need to compare them again`. This optimization involves reducing the number of elements being compared in each pass.
1. Termination: The algorithm terminates when no more swaps are needed, indicating that the array is sorted.

??? example "Illustration"
    <figure markdown="span">
        ![Step 1 -Bubble Sort – Data Structure and Algorithm Tutorials from geeksforgeeks](https://media.geeksforgeeks.org/wp-content/uploads/20230526103842/1.webp){ width="600" }
        <figcaption>[*Step 1 -Bubble Sort – Data Structure and Algorithm Tutorials from geeksforgeeks*](https://media.geeksforgeeks.org/wp-content/uploads/20230526103842/1.webp)</figcaption>
    </figure>
    <figure markdown="span">
        ![Step 2 -Bubble Sort – Data Structure and Algorithm Tutorials from geeksforgeeks](https://media.geeksforgeeks.org/wp-content/uploads/20230526103914/2.webp){ width="600" }
        <figcaption>[*Step 2 -Bubble Sort – Data Structure and Algorithm Tutorials from geeksforgeeks*](https://media.geeksforgeeks.org/wp-content/uploads/20230526103842/1.webp)</figcaption>
    </figure>
    <figure markdown="span">
        ![Step 3 -Bubble Sort – Data Structure and Algorithm Tutorials from geeksforgeeks](https://media.geeksforgeeks.org/wp-content/uploads/20230526103949/3.webp){ width="600" }
        <figcaption>[*Step 3 -Bubble Sort – Data Structure and Algorithm Tutorials from geeksforgeeks*](https://media.geeksforgeeks.org/wp-content/uploads/20230526103842/1.webp)</figcaption>
    </figure>

!!! example
    ```bash
    $ go run src/sort/bubble_sort.go
    Inital data:  [9 3 7 4 69 420 42]
    Final data:  [3 4 7 9 42 69 420]
    ```

    ```go
    --8<-- "src/sort/bubble_sort.go"
    ```
