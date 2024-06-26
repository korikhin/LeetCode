## 125. Valid Palindrome

A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string `s`, _return `true` if it is a **palindrome**, or `false` otherwise_.

<br>

### Example 1

```
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
```

### Example 2

```
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
```

### Example 3

```
Input: s = " "
Output: true
Explanation:
  The s is an empty string "" after removing non-alphanumeric characters.
  Since an empty string reads the same forward and backward, it is a palindrome.
```

<br>

### Constraints

- $1 \leqslant$ `len(s)` $\leqslant 2 \cdot 10^5$
- `s` consists only of printable ASCII characters.
