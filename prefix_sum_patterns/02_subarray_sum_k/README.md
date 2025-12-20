🎯 Goal:
You should instantly recognize:

“This is Prefix Sum + HashMap”

## Pattern: Prefix Sum + HashMap (Subarray Sum = K)

When to use:
- Count subarrays with given sum
- Negative numbers present
- Binary arrays with sum constraints

Key Idea:
If prefix[j] - prefix[i] = k,
then prefix[i] = prefix[j] - k

Why it works:
- Stores past prefix sums
- Converts O(n²) brute force to O(n)

Problems:
- Subarray Sum Equals K (LC 560)
- Binary Subarrays With Sum (LC 930)



🖼️ MENTAL IMAGE (VERY IMPORTANT)
🧾 Bank Balance Analogy

currentSum = current balance

Want a subarray sum = k

Ask: Have I seen balance (currentSum − k) before?

If YES → valid subarray exists.

🧱 STEP 1: Create Folder
mkdir 02_subarray_sum_k
cd 02_subarray_sum_k

✍️ STEP 2: Problem 1 — Subarray Sum Equals K (LC 560)
Why this problem is CRITICAL

Negative numbers ❌ sliding window

Counting subarrays

Direct FAANG favorite