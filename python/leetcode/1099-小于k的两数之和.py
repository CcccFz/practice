# https://leetcode-cn.com/problems/two-sum-less-than-k/


class Solution:
    def twoSum(self, nums: List[int], k: int) -> int:
        if not nums:
            return -1

        nums = sorted(nums)

        i, j, max_value = 0, len(nums) - 1, float('-inf')
        while i < j:
            if nums[i] + nums[j] >= k:
                j -= 1
            else:
                max_value = max(max_value, nums[i] + nums[j])
                i += 1

        return -1 if max_value == float('-inf') else max_value
