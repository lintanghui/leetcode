"""
Source: https://leetcode.com/problems/string-to-integer-atoi/
Author: Lin Tanghui
Date  : 2015/11/1
"""

"""
Implement atoi to convert a string to an integer.
Hint: Carefully consider all possible input cases. 
If you want a challenge, please do not see below and ask yourself what are the possible input cases.
"""
class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        str = str.strip()
        for i,v in enumerate(str):
            if v not in '+-' and not v.isdigit():
                str = str[:i]
                break
        try:
            ret = int(str)
        except:
            return 0
        return max(min(ret, 2**31-1), -2**31)