🧠 PARTITIONING / DUTCH NATIONAL FLAG (CORE IDEA)

We divide the array into 3 regions:

[ 0s | 1s | unknown | 2s ]
   ^     ^          ^
  low   mid        high

Pointer meaning

low → boundary for 0s

mid → current element

high → boundary for 2s

🔁 UNIVERSAL TEMPLATE (MEMORIZE THIS)
low, mid, high := 0, 0, len(nums)-1

for mid <= high {
    if nums[mid] == 0 {
        swap(nums[low], nums[mid])
        low++
        mid++
    } else if nums[mid] == 1 {
        mid++
    } else {
        swap(nums[mid], nums[high])
        high--
    }
}