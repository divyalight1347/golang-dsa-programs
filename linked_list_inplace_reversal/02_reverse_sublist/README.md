# Reverse Sublist (Linked List In-place Reversal)

LeetCode: 92 — Reverse Linked List II  
Difficulty: Medium  
Pattern: Linked List In-place Reversal

---

## 🧠 Problem Summary

Given a linked list, reverse the nodes from position `left` to `right` (1-indexed) **in-place**.

Example:
Input:  1 → 2 → 3 → 4 → 5, left = 2, right = 4  
Output: 1 → 4 → 3 → 2 → 5

---

## 🔑 Core Idea (Memorize This)

> Skip nodes until `left - 1`,  
> reverse exactly `right - left + 1` nodes,  
> then reconnect the list.

---

## 🧠 High-Level Steps

1. Use a **dummy node** to handle edge cases
2. Move `prev` to node before `left`
3. Reverse `right - left + 1` nodes using 3-pointer technique
4. Reconnect the reversed sublist back to the main list

---

## 🔁 Reversal Template (Golden Rule)

```text
next = curr.Next
curr.Next = prev
prev = curr
curr = next


🧠 Memory Trick

✂️ Cut → Reverse → Stitch

Imagine cutting the sublist out, flipping it, and stitching it back.

⏱️ Complexity

Time: O(n)

Space: O(1) (in-place)