## Pattern: Prefix Sum + Modulo

When to use:
- Subarray sum divisible by k
- Modulo-based constraints

Key Idea:
If prefix[i] % k == prefix[j] % k,
then sum(i+1..j) is divisible by k

Why it works:
- Same remainder implies difference is multiple of k

Important Trick:
(remainder + k) % k to handle negatives

Problems:
- Subarray Sums Divisible by K (LC 974)
- Continuous Subarray Sum (LC 523)


🔑 CORE IDEA (SAY THIS ALOUD)

If two prefix sums have the same remainder when divided by k, their difference is divisible by k.