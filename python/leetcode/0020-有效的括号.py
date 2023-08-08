# https://leetcode-cn.com/problems/valid-parentheses/
# https://github.com/azl397985856/leetcode/blob/master/problems/20.validParentheses.md
# 如果让你检查XML标签是否闭合如何检查， 更进一步如果要你实现一个简单的XML的解析器，应该怎么实现？

class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        dic = {'(': ')', '[': ']', '{': '}'}

        for x in s:
            if x in dic:
                stack.append(x)
            else:
                if len(stack):
                    if x == dic[stack.pop()]:
                        continue
                    else:
                        return False
                else:
                    return False

        return not len(stack)
