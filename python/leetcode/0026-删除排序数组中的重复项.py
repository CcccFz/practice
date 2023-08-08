# https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
# https://github.com/azl397985856/leetcode/blob/master/problems/26.remove-duplicates-from-sorted-array.md


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0

        slow = 0
        for fast in range(1, len(nums)):
            if nums[fast] != nums[slow]:
                slow += 1
                nums[slow] = nums[fast]

        return slow + 1
