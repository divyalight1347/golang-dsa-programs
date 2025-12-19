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