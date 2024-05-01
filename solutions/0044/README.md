## 44. Wildcard Matching

Given an input string `s` and a pattern `p`, implement wildcard pattern matching with support for `'?'` and `'*'` where:

* `'?'` Matches any single character.
* `'*'` Matches any sequence of characters (including the empty sequence).

The matching should cover the entire input string (not partial).

<br>

### Example 1

```
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

### Example 2

```
Input: s = "aa", p = "*"
Output: true
Explanation: '*' matches any sequence.
```

### Example 3

```
Input: s = "cb", p = "?a"
Output: false
Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
```

<br>

### Constraints

* $0 \leqslant$ `s.length, p.length` $\leqslant 2000$
* `s` contains only lowercase English letters.
* `p` contains only lowercase English letters, `'?'` or `'*'`.