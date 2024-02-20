# Big O Time Complexity

Big O notation is a mathematical notation used in computer science to `describe the performance or complexity of an algorithm`. It specifically focuses on the **`worst-case scenario`**, which is the `maximum amount of time taken by an algorithm to complete its execution`.

The objective of Big O notation is to provide a high-level understanding of the efficiency of an algorithm, allowing developers to make informed decisions about which algorithms to use in different situations. **`It helps in comparing the efficiency of algorithms with varying input sizes and is particularly useful in the analysis of algorithms for large data sets`**.

**`Big O notation works by analyzing the number of operations an algorithm performs relative to the size of its input`**. It does not measure the actual time taken by an algorithm but rather the number of steps it takes to complete. For example, if an algorithm has a time complexity of O(n), it means that the number of operations increases linearly with the size of the input.

!!! info "Important Concepts"
    1. Growth is with respect to the input
    1. Constants are dropped
    1. Worst case is usually the way we measure

## Constants are Dropped

When expressing time complexity in Big O notation, `constants are dropped because they do not significantly affect the growth rate of the algorithm as the input size increases`. For example, if an algorithm has a time complexity of O(2n), it is simplified to O(n) because the constant factor of 2 does not change the linear growth rate. **`The focus in Big O notation is on the rate of growth, not the absolute number of operations`**.

## Big O Classification

There are several types of Big O notations, each representing a different growth rate:

1. O(1): `Constant time complexity`. The algorithm takes the same amount of time to complete regardless of the input size.
1. O(log n): `Logarithmic time complexity`. The algorithm's time to complete increases logarithmically with the input size, which is common in divide-and-conquer algorithms.
1. O(n): `Linear time complexity`. The algorithm's time to complete increases linearly with the input size.
1. O(n log n): `Linearithmic time complexity`. The algorithm's time to complete is between linear and quadratic, often seen in efficient sorting algorithms like merge sort and quicksort.
1. O(n^2): `Quadratic time complexity`. The algorithm's time to complete increases quadratically with the input size.
1. O(2^n): `Exponential time complexity`. This complexity indicates that the algorithm's time to complete doubles with each additional element in the input.
1. O(n!): `Factorial time complexity`. This complexity is the worst-case scenario, indicating that the number of operations grows factorially with the input size.

![[Big o Cheatsheet - Data structures and Algorithms with thier complexities](https://www.hackerearth.com/practice/notes/big-o-cheatsheet-series-data-structures-and-algorithms-with-thier-complexities-1/)](https://he-s3.s3.amazonaws.com/media/uploads/ece920b.png)
