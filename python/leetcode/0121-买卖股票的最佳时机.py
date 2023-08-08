# https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
# https://github.com/azl397985856/leetcode/blob/master/problems/121.best-time-to-buy-and-sell-stock.md


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if not prices:
            return 0

        min_price = float('inf')
        max_profit = 0

        for price in prices:
            if price < min_price:
                min_price = price
            if max_profit < price - min_price:
                max_profit = price - min_price

        return max_profit
