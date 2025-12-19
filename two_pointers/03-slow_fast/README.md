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


8️⃣ Reorder List — Input: 1→2→3→4 → Output: 1→4→2→3
9️⃣ Detect Cycle in Circular Array — Input: [2,-1,1,2,2] → Output: true
🔟 Circular Array Loop — Input: [1,1,2] → Output: true

🧠 WHEN TO USE SLOW–FAST

If the problem mentions:

Cycle / loop

Middle element

Duplicate without modifying array

Linked list traversal