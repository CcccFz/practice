# https://leetcode-cn.com/problems/trapping-rain-water/
# https://github.com/azl397985856/leetcode/blob/master/problems/42.trapping-rain-water.md


class Solution:
    def trap(self, height: List[int]) -> int:
        volume = 0
        max_height = 0
        left_max_height = []
        right_max_height = []
        cnt = len(height)

        for i in range(cnt):
            max_height = max(max_height, height[i])
            left_max_height.append(max_height)

        max_height = 0

        for i in range(cnt - 1, -1, -1):
            max_height = max(max_height, height[i])
            right_max_height.insert(0, max_height)

        for i in range(cnt):
            volume += min(left_max_height[i], right_max_height[i]) - height[i]

        return volume
