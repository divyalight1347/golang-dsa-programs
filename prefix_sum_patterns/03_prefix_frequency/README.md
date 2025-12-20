## Pattern: Prefix Sum Frequency

When to use:
- Count subarrays with exact conditions
- Frequency matters more than position
- Array can be transformed (odd/even, 0/1)

Key Idea:
- Convert the problem into subarray sum = k
- Store frequency of prefix sums

Why it works:
- Each matching prefix frequency contributes to multiple subarrays

Problems:
- Count Number of Nice Subarrays (LC 1248)
