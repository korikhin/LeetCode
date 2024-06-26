## 289

The board is made up of an$m \times n$ grid of cells, where each cell has an initial state: **live** (represented by a `1`) or **dead** (represented by a `0`). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

- Any live cell with fewer than two live neighbors dies as if caused by under-population.
- Any live cell with two or three live neighbors lives on to the next generation.
- Any live cell with more than three live neighbors dies, as if by over-population.
- Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously. Given the current state of the $m \times n$ grid `board`, return the next state.

<br>

### Example 1

```
Input: board = [[0, 1, 0], [0, 0, 1], [1, 1, 1], [0, 0, 0]]
Output: [[0, 0, 0], [1, 0, 1], [0, 1, 1], [0, 1, 0]]
```

### Example 2

```
Input: board = [[1, 1], [1, 0]]
Output: [[1, 1], [1, 1]]
```

<br>

### Constraints

- `m == len(board)`
- `n == len(board[i])`
- $1 \leqslant$ `m`, `n` $\leqslant 25$
- `board[i][j]` is `0` or `1`

<br>

**Follow up:**

- Could you solve it in-place? Remember that the board needs to be updated simultaneously: You cannot update some cells first and then use their updated values to update other cells.
- In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches upon the border of the array (i.e., live cells reach the border). How would you address these problems?
