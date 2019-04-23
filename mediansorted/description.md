### 4. Median of Two Sorted Arrays

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

```
Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
```

```
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
```

*Notes*
- if we're going to get o(log (m+n)) time complexity you'll at least need to do a merge sort
- **but** if you use a count sort you're looking at an overall complexity of o(n+k) - superior to mergesort
- remember that the index array's length should be based on the highest _value_ not the combined lenghts of the arrays
