## Pattern: Basic Prefix Sum

When to use:
- Multiple range sum queries
- Immutable array

Key Formula:
prefix[i+1] = prefix[i] + nums[i]
sum(l..r) = prefix[r+1] - prefix[l]

Why it works:
- Avoids recomputing sums
- Converts O(n) range sum to O(1)

Problems:
- Running Sum (LC 1480)
- Range Sum Query – Immutable (LC 303)
