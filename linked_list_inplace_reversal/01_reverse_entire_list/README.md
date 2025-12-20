Pattern: In-place Linked List Reversal

Core Idea:
Reverse the direction of pointers using 3 pointers

Steps:
1. Save next node
2. Reverse current pointer
3. Move forward

Template:
next = curr.Next
curr.Next = prev
prev = curr
curr = next

🧠 MEMORY TRICK (VERY IMPORTANT)
Visualize a TRAIN 🚆

Each coach points backward instead of forward
You reverse the connection one coach at a time

🔒 INTERVIEW LOCK QUESTIONS

If interviewer asks:

“Can you do it in O(1) space?”

You immediately say:

“Yes — in-place using 3 pointers.”