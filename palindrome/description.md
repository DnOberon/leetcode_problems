### 5. Longest Palindromic Substring

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

```
Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```

```
Example 2:

Input: "cbbd"
Output: "bb"
```

*Notes*
- if you want O(n) time complexity, you must use Manacher's Algorithm
- find the lenght of each palindrome of each index, then use that information to build the longest

