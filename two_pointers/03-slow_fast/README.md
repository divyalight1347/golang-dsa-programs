🧠 SLOW & FAST PATTERN (CORE IDEA)

slow → moves 1 step

fast → moves 2 steps

If they meet, a cycle exists

🧠 Memory trick

🐢 (slow) and 🐇 (fast) running on a circular track

🔁 UNIVERSAL TEMPLATE
slow, fast := start, start

for fast condition {
    slow = move once
    fast = move twice

    if slow == fast {
        cycle found
    }
}


🥇 TIER-1 (ABSOLUTE MUST-DO)

1️⃣ Linked List Cycle — Input: 3→2→0→-4 (cycle at 2) → Output: true
2️⃣ Linked List Cycle II (Find cycle start) — Input: 1→2→3→4→2 → Output: Node(2)
3️⃣ Middle of the Linked List — Input: 1→2→3→4→5 → Output: 3
4️⃣ Happy Number — Input: 19 → Output: true

➡️ Asked by Google, Amazon, Meta very frequently

🥈 TIER-2 (VERY COMMON)

5️⃣ Remove Nth Node From End of List — Input: 1→2→3→4→5, n=2 → Output: 1→2→3→5
6️⃣ Palindrome Linked List — Input: 1→2→2→1 → Output: true
7️⃣ Find Duplicate Number — Input: [1,3,4,2,2] → Output: 2

➡️ Asked by Amazon, Google, Apple

🥉 TIER-3 (OCCASIONAL BUT IMPORTANT)

8️⃣ Reorder List — Input: 1→2→3→4 → Output: 1→4→2→3
9️⃣ Detect Cycle in Circular Array — Input: [2,-1,1,2,2] → Output: true
🔟 Circular Array Loop — Input: [1,1,2] → Output: true

🧠 WHEN TO USE SLOW–FAST

If the problem mentions:

Cycle / loop

Middle element

Duplicate without modifying array

Linked list traversal