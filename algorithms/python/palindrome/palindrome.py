"""
Source: https://leetcode.com/problems/palindrome-number/
Author: Lin Tanghui
Date  : 2015/11/1
"""

"""
Determine whether an integer is a palindrome. 
Do this without extra space.
"""
class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        x = str(x)
        if len(x)==1:
            return True
        lenth = len(x)-1
        i =0
        while (x[i]==x[lenth] and i<len(x)/2):
            i=i+1
            lenth=lenth-1
        return False if x[i]!=x[lenth] else True