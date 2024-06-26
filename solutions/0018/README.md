## 18. 4Sum

Given an array `nums` of `n` integers, return _an array of all the **unique** quadruplets_ `[nums[a], nums[b], nums[c], nums[d]]` such that:

- $0 \leqslant$ `a`, `b`, `c`, `d` $<$ `n`
- `a`, `b`, `c`, and `d` are **distinct**.
- `nums[a] + nums[b] + nums[c] + nums[d] == target`

You may return the answer in **any order**.

<br>

### Example 1

```
Input: nums = [1, 0, -1, 0, -2, 2], target = 0
Output: [[-2, -1, 1, 2], [-2, 0, 0, 2], [-1, 0, 0, 1]]
```

### Example 2

```
Input: nums = [2, 2, 2, 2, 2], target = 8
Output: [[2, 2, 2, 2]]
```

<br>

### Constraints

- $1 \leqslant$ `len(nums)` $\leqslant 200$
- $-10^9 \leqslant$ `nums[i]` $\leqslant 10^9$
- $-10^9 \leqslant$ `target` $\leqslant 10^9$
