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