## 25. Reverse Nodes in k-Group

Given the `head` of a linked list, reverse the nodes of the list `k` at a time, and return _the modified list_.

Positive integer `k` is less than or equal to the length of the linked list. If the number of nodes is not a multiple of `k` then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

<br>

### Example 1

```
Input: head = [1, 2, 3,  4,  5], k = 2
Output: [2, 1, 4, 3, 5]
```

### Example 2

```
Input: head = [1, 2, 3, 4, 5], k = 3
Output: [3, 2, 1, 4, 5]
```

<br>

### Constraints

- The number of nodes in the list is `n`.
- $1 \leqslant$ `k` $\leqslant$ `n` $\leqslant 5000$
- $0 \leqslant$ `Node.Val` $\leqslant 100$
