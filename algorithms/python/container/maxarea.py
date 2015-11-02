"""
Source: https://leetcode.com/problems/container-with-most-water/
Author: Lin Tanghui
Date  : 2015/11/1
"""

"""
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). 
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). 
Find two lines, which together with x-axis forms a container,
such that the container contains the most water.
"""
class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        maxVal= 0
        # Use two pointer: head pointer and tail pointer
        i = 0;j = len(height)-1
        while i != j:
            currVal = j-i
            if height[i] < height[j]:
                currVal, i, j = currVal * height[i], i+1, j 
            else :
                currVal, i, j = currVal * height[j], i, j-1 
            maxVal = currVal if currVal > maxVal else maxVal
        return maxVal;
        