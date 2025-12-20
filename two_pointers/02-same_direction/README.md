🧠 SAME DIRECTION PATTERN (CORE IDEA)

fast → reads every element

slow → writes valid elements

Everything before slow is the answer

🔁 Template (MEMORIZE)
slow := 0
for fast := 0; fast < len(nums); fast++ {
    if condition {
        nums[slow] = nums[fast]
        slow++
    }
}



1️⃣ Remove duplicates from sorted array — Input: [1,1,2,2,3] → Output: [1,2,3]
2️⃣ Remove duplicates from sorted array II — Input: [1,1,1,2,2,3] → Output: [1,1,2,2,3]
3️⃣ Move zeroes — Input: [0,1,0,3,12] → Output: [1,3,12,0,0]
4️⃣ Remove element — Input: [3,2,2,3], val=3 → Output: [2,2]
5️⃣ Compress string — Input: ["a","a","b","b","c","c","c"] → Output: ["a","2","b","2","c","3"]
6️⃣ Duplicate zeros — Input: [1,0,2,3,0,4,5,0] → Output: [1,0,0,2,3,0,0,4]
7️⃣ Sort array by parity — Input: [3,1,2,4] → Output: [2,4,3,1]
8️⃣ Move negative numbers to beginning — Input: [-1,3,-2,4,5] → Output: [-1,-2,3,4,5]
9️⃣ Remove all occurrences of a character — Input: "banana", remove='a' → Output: "bnn"
🔟 Merge two sorted arrays in-place — Input: nums1=[1,2,3,0,0,0], nums2=[2,5,6] → Output: [1,2,2,3,5,6]