### 200. Number of Islands

Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

```
Example 1:

Input:
11110
11010
11000
00000

Output: 1
```

```
Example 2:

Input:
11000
11000
00100
00011

Output: 3
```

*Notes*
- a depth first search is the fasted way to solve this problem.
- the gride input is considered and undirected graph and should be treated as if each
number has a max of four vertices
- don't forget to check rows/columns x/y - it's easy to mix the placement of those up
