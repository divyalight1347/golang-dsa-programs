sliding-window/
│
├── 01_fixed_window/
│   ├── max_sum_subarray_k.go
│   └── README.md
│
├── 02_variable_window/
│   ├── longest_substring_no_repeat.go
│   ├── longest_substring_k_distinct.go
│   └── README.md
│
├── 03_window_with_frequency/
│   ├── minimum_window_substring.go
│   └── README.md
│
├── 04_window_with_condition/
│   ├── fruit_into_baskets.go
│   └── README.md
│
├── 05_two_pointer_window/
│   ├── container_with_most_water.go
│   └── README.md
│
└── README.md   ← MASTER CHEAT SHEET


# Sliding Window — FAANG Master Pattern

Core Idea:
Maintain a window using two pointers and update results without reprocessing elements.

Types:
1. Fixed Size Window
2. Variable Size Window
3. Window with Frequency Map
4. Window with Condition

Golden Template:
Expand right pointer
Shrink left pointer when condition breaks
Update result

Time Complexity:
O(n)

Common Pitfalls:
- Forgetting to shrink window
- Updating result at wrong time
- Overcomplicating conditions


🧠 TWO CORE TYPES (MEMORIZE THIS FIRST)
1️⃣ Fixed Size Window

Window size = k

Mostly sums, averages

2️⃣ Variable Size Window

Window grows/shrinks based on condition

Most FAANG problems


🧠 GOLDEN SLIDING WINDOW TEMPLATE
Variable Window (MOST IMPORTANT)
left := 0
for right := 0; right < len(s); right++ {
	// include right

	for condition is broken {
		// remove left
		left++
	}

	// update answer
}


🔥 This template solves 70% of problems.

🧠 INTERVIEW TRIGGERS (VERY IMPORTANT)

Use sliding window when:

Contiguous subarray/substring

Optimize length / sum

Two pointers move forward

No need to revisit elements