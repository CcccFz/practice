# https://leetcode-cn.com/problems/maximum-subarray/
# https://github.com/azl397985856/leetcode/blob/master/problems/53.maximum-sum-subarray-cn.md
# 如果数组是二维数组，求最大子数组的和？
# 如果要求最大子序列的乘积？


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0

        global_max = current_max = nums[0]
        for i in range(1, len(nums)):
            current_max = max(current_max + nums[i], nums[i])
            global_max = max(global_max, current_max)

        return global_max
